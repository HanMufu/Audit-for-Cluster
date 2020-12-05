package neo4jdb

import(
	"fmt"
	"strings"
	// "github.com/neo4j/neo4j-go-driver/neo4j"
	"go.uber.org/zap"
)

var MachineID string

/*
SYSCALL
audit(1607204155.092:2130707): arch=c000003e syscall=228 success=yes exit=0 a0=7 a1=7ffeaa82f500 a2=55a7c0776cf0 a3=0 items=0 ppid=1816 pid=1822 auid=1000 uid=1000 gid=1001 euid=1000 suid=1000 fsuid=1000 egid=1001 sgid=1001 fsgid=1001 tty=(none) ses=6 comm="sshd" exe="/usr/sbin/sshd" subj==unconfined key=(null)
0 audit(1607204155.092:2130707):
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
25 subj==unconfined
26 key=(null)
*/
func InsertToDB(eventType string, msg string, machineID string) (err error) {
	if(eventType != "SYSCALL") {
		return
	}
	fmt.Println(eventType)
	fmt.Println(msg)
	message := strings.Split(msg, " ")
	m := make(map[string]string)
	for _, subMsg := range message {
		tmp := strings.Split(subMsg, "=")
		if(len(tmp) == 2) {
			m[tmp[0]] = tmp[1]
			fmt.Println(tmp[0], tmp[1])
			fmt.Println("map:", m)
		}
	}
	// create a new database for this machine if not existed

	// store this process into DB (adding vertex)

	// if _, err := session.
	// 	WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
	// 		return u.persistUser(tx, user)
	// 	}); err != nil {
	// 	return err
	// }

	// connect this process to parent process(ppid) (adding edge)
	// connect this process to user who currently login(auid) (adding edge)
	// connect this process to current user group(gid) (adding edge)
	return
}

// func persistUser(tx neo4j.Transaction, user *User) (interface{}, error) {
// 	query := "CREATE (:User {email: $email, username: $username, password: $password})"
// 	hashedPassword, err := hash(user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	parameters := map[string]interface{}{
// 		"email":    user.Email,
// 		"username": user.Username,
// 		"password": hashedPassword,
// 	}
// 	_, err = tx.Run(query, parameters)
// 	return nil, err
// }

