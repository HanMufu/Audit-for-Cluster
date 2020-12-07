package neo4jdb

import (
	"fmt"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.uber.org/zap"
	// "github.com/neo4j/neo4j-go-driver/neo4j"
	// "go.uber.org/zap"
)

/*
SYSCALL
audit(1607204155.092:2130707): arch=c000003e syscall=228 success=yes exit=0 a0=7 a1=7ffeaa82f500 a2=55a7c0776cf0 a3=0 items=0 ppid=1816 pid=1822 auid=1000 uid=1000 gid=1001 euid=1000 suid=1000 fsuid=1000 egid=1001 sgid=1001 fsgid=1001 tty=(none) ses=6 comm="sshd" exe="/usr/sbin/sshd" subj==unconfined key=(null)
0 audit(1607204155.092:2130707): (ignore this line)
1 arch=c000003e
2 syscall=228
3 success=yes
4 exit=0
5 a0=7
6 a1=7ffeaa82f500
7 a2=55a7c0776cf0
8 a3=0
9 items=0
10 ppid=1816
11 pid=1822
12 auid=1000
13 uid=1000
14 gid=1001
15 euid=1000
16 suid=1000
17 fsuid=1000
18 egid=1001
19 sgid=1001
20 fsgid=1001
21 tty=(none)
22 ses=6
23 comm="sshd"
24 exe="/usr/sbin/sshd"
25 subj==unconfined (ignore this line)
26 key=(null)
*/
func InsertToDB(eventType string, msg string, machineID string) (err error) {
	if eventType != "SYSCALL" {
		return
	}
	fmt.Println(eventType)
	fmt.Println(msg)
	message := strings.Split(msg, " ")
	m := make(map[string]string)
	for id, subMsg := range message {
		if id == 0 {
			m["audit"] = subMsg
		}
		tmp := strings.Split(subMsg, "=")
		if len(tmp) == 2 {
			m[tmp[0]] = tmp[1]
			// fmt.Println(id, subMsg)
			// fmt.Println(tmp[0], tmp[1])
			// fmt.Println("map:", m)
		}
	}
	// Add current SYSCALL into DB
	registerSYSCALL(m, machineID)
	// Add current process into DB if not existed
	registerProcess(m, machineID, "pid")
	// Add parent process into DB if not existed
	registerProcess(m, machineID, "ppid")
	// Add an edge point to parent process
	registerParentProcessEdge(m, machineID)
	// _, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
	// 	return findProcess(tx, m, machineID)
	// })

	// connect this process to parent process(ppid) (adding edge)
	// connect this process to user who currently login(auid) (adding edge)
	// connect this process to current user group(gid) (adding edge)
	return
}

func registerSYSCALL(m map[string]string, machineID string) (err error) {
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer func() {
		err = session.Close()
	}()
	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return presistSYSCALL(tx, m, machineID)
		}); err != nil {
		zap.L().Debug("Add SYSCALL failed. ", zap.Error(err))
		return err
	}
	return nil
}

func presistSYSCALL(tx neo4j.Transaction, m map[string]string, machineID string) (interface{}, error) {
	zap.L().Info("presistSYSCALL...")
	query := `CREATE (:SYSCALL { 
		arch: $arch, 
		pid: $pid,
		syscall: $syscall,
		succes: $succes,
		exit: $exit,
		a0: $a0,
		a1: $a1,
		a2: $a2,
		a3: $a3,
		items: $items,
		ppid: $ppid,
		auid: $auid,
		uid: $uid,
		gid: $gid,
		euid: $euid,
		suid: $suid,
		fsuid: $fsuid,
		egid: $egid,
		sgid: $sgid,
		fsgid: $fsgid,
		tty: $tty,
		ses: $ses,
		comm: $comm,
		exe: $exe,
		key: $key, 
		machineID: $machineID, 
		auditTimeStamp: $audit
	});`
	parameters := map[string]interface{}{
		"arch":      m["acrh"],
		"pid":       m["pid"],
		"syscall":   m["syscall"],
		"succes":    m["succes"],
		"exit":      m["exit"],
		"a0":        m["a0"],
		"a1":        m["a1"],
		"a2":        m["a2"],
		"a3":        m["a3"],
		"items":     m["items"],
		"ppid":      m["ppid"],
		"auid":      m["auid"],
		"uid":       m["uid"],
		"gid":       m["gid"],
		"euid":      m["euid"],
		"suid":      m["suid"],
		"fsuid":     m["fsuid"],
		"egid":      m["egid"],
		"sgid":      m["sgid"],
		"fsgid":     m["fsgid"],
		"tty":       m["tty"],
		"ses":       m["ses"],
		"comm":      m["comm"],
		"exe":       m["exe"],
		"key":       m["key"],
		"machineID": machineID,
		"audit":     m["audit"],
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}

func registerProcess(m map[string]string, machineID string, whichProcess string) (err error) {
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer func() {
		err = session.Close()
	}()
	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return presistProcess(tx, m, machineID, whichProcess)
		}); err != nil {
		zap.L().Debug("Add Process failed. ", zap.Error(err))
		return err
	}
	return nil
}

func presistProcess(tx neo4j.Transaction, m map[string]string, machineID string, whichProcess string) (interface{}, error) {
	zap.L().Info("presistProcess...")
	query := `MERGE (:PROCESS { 
		pid: $pid,
		machineID: $machineID
	});`
	parameters := map[string]interface{}{
		"pid":       m[whichProcess],
		"machineID": machineID,
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}

func registerParentProcessEdge(m map[string]string, machineID string) (err error) {
	session := driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer func() {
		err = session.Close()
	}()
	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return presistParentProcessEdge(tx, m, machineID)
		}); err != nil {
		zap.L().Debug("Add ParentProcessEdge failed. ", zap.Error(err))
		return err
	}
	return nil
}

func presistParentProcessEdge(tx neo4j.Transaction, m map[string]string, machineID string) (interface{}, error) {
	zap.L().Info("presistParentProcessEdge...")
	query := `MATCH (a:PROCESS),(b:PROCESS) where a.pid=$pid and b.pid=$ppid and a.machineID=$machineID and a.machineID=b.machineID MERGE (a)-[:FORK_FROM]->(b);`
	parameters := map[string]interface{}{
		"pid":       m["pid"],
		"ppid":      m["ppid"],
		"machineID": machineID,
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}
