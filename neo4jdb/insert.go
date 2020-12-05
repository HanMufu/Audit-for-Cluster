package neo4jdb

import(
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func InsertToDB(eventType string, msg string) (err error) {
	fmt.Println(eventType)
	// fmt.Println(msg)
	return
}

func TestConnection() (string, error) {
	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}
