package cassandra

import (
	"os"
	"time"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

func ConnectDatabase() *gocql.Session {

	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 5
	cluster.Timeout = time.Second * 5
	cluster.NumConns = 10
	cluster.ReconnectInterval = time.Second * 1
	cluster.SocketKeepalive = 0
	cluster.DisableInitialHostLookup = true
	cluster.IgnorePeerAddr = true
	cluster.Events.DisableNodeStatusEvents = true
	cluster.Events.DisableTopologyEvents = true
	cluster.Events.DisableSchemaEvents = true
	cluster.WriteCoalesceWaitTime = 0
	cluster.ReconnectionPolicy = &gocql.ConstantReconnectionPolicy{MaxRetries: 5000, Interval: 5 * time.Second}
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASS"),
	}
	session, err := cluster.CreateSession()
	if err != nil {
		clogger.Error(&logger.Messages{
			"connection err: ": err.Error(),
		})
		return nil
	}

	if err = session.Query(`CREATE KEYSPACE IF NOT EXISTS MLDatabase WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3}`).Exec(); err != nil {
		clogger.Error(&logger.Messages{
			"create keyspace err: ": err.Error(),
		})
	}


	for _, q := range GetTableQueries() {
		err = session.Query(q).Exec()
		clogger.Error(&logger.Messages{
			"create table err: ": err.Error(),
		})
	}

	return session
}
