package neo4jdb

import(
	"fmt"
)

func InsertToDB(eventType string, msg string) (err error) {
	fmt.Println(eventType)
	fmt.Println(msg)
	return
}
