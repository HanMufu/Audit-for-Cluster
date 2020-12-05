package neo4jdb

import (
	"audit-cluster/settings"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"go.uber.org/zap"
)

var driver neo4j.Driver
var session neo4j.Session

func Init(config *settings.Neo4jConfig) (err error) {
	driver, err := neo4j.NewDriver(config.Host, neo4j.BasicAuth(config.User, config.Password, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		return err
	}

	session, err = driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}

	zap.L().Info("Init neo4j success")

	return
}

func Close() {
	_ = driver.Close()
	_ = session.Close()
}
