package logic

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/elastic/go-libaudit/v2"
	"github.com/elastic/go-libaudit/v2/auparse"
	"github.com/pkg/errors"

	"audit-cluster/neo4jdb"
)

var (
	fs          = flag.NewFlagSet("audit", flag.ExitOnError)
	diag        = fs.String("diag", "", "dump raw information from kernel to file")
	rate        = fs.Uint("rate", 0, "rate limit in kernel (default 0, no rate limit)")
	backlog     = fs.Uint("backlog", 8192, "backlog limit")
	immutable   = fs.Bool("immutable", false, "make kernel audit settings immutable (requires reboot to undo)")
	receiveOnly = fs.Bool("ro", false, "receive only using multicast, requires kernel 3.16+")
)

func Read(machineID string) error {
	if os.Geteuid() != 0 {
		return errors.New("you must be root to receive audit data")
	}

	// Write netlink response to a file for further analysis or for writing
	// tests cases.
	var diagWriter io.Writer
	if *diag != "" {
		f, err := os.OpenFile(*diag, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
		diagWriter = f
	}

	log.Println("starting netlink client")

	var err error
	var client *libaudit.AuditClient
	if *receiveOnly {
		client, err = libaudit.NewMulticastAuditClient(diagWriter)
		if err != nil {
			return errors.Wrap(err, "failed to create receive-only audit client")
		}
		defer client.Close()
	} else {
		client, err = libaudit.NewAuditClient(diagWriter)
		if err != nil {
			return errors.Wrap(err, "failed to create audit client")
		}
		defer client.Close()

		status, err := client.GetStatus()
		if err != nil {
			return errors.Wrap(err, "failed to get audit status")
		}
		log.Printf("received audit status=%+v", status)

		if status.Enabled == 0 {
			log.Println("enabling auditing in the kernel")
			if err = client.SetEnabled(true, libaudit.WaitForReply); err != nil {
				return errors.Wrap(err, "failed to set enabled=true")
			}
		}

		if status.RateLimit != uint32(*rate) {
			log.Printf("setting rate limit in kernel to %v", *rate)
			if err = client.SetRateLimit(uint32(*rate), libaudit.NoWait); err != nil {
				return errors.Wrap(err, "failed to set rate limit to unlimited")
			}
		}

		if status.BacklogLimit != uint32(*backlog) {
			log.Printf("setting backlog limit in kernel to %v", *backlog)
			if err = client.SetBacklogLimit(uint32(*backlog), libaudit.NoWait); err != nil {
				return errors.Wrap(err, "failed to set backlog limit")
			}
		}

		if status.Enabled != 2 {
			log.Printf("setting kernel settings as immutable")
			if err = client.SetImmutable(libaudit.NoWait); err != nil {
				return errors.Wrap(err, "failed to set kernel as immutable")
			}
		}

		log.Printf("sending message to kernel registering our PID (%v) as the audit daemon", os.Getpid())
		if err = client.SetPID(libaudit.NoWait); err != nil {
			return errors.Wrap(err, "failed to set audit PID")
		}
	}

	return receive(client, machineID)
}

func receive(r *libaudit.AuditClient, machineID string) error {
	for {
		rawEvent, err := r.Receive(false)
		if err != nil {
			return errors.Wrap(err, "receive failed")
		}

		// Messages from 1300-2999 are valid audit messages.
		if rawEvent.Type < auparse.AUDIT_USER_AUTH ||
			rawEvent.Type > auparse.AUDIT_LAST_USER_MSG2 {
			continue
		}

		// fmt.Printf("type=%v msg=%v\n", rawEvent.Type, string(rawEvent.Data))
		// store into neo4j
		neo4jdb.InsertToDB(fmt.Sprintf("%v", rawEvent.Type), string(rawEvent.Data), machineID)
	}
}
