package levelBaseSessionConnection

import (
	"os"
	"time"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func ConnectDatabase() (*gocqlx.Session, error) {
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
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		clogger.Error(&logger.Messages{
			"connection err: ": err.Error(),
		})
		return nil, err
	}
	if err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS MLDatabase WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3}`); err != nil {
		clogger.Error(&logger.Messages{
			"create keyspace err: ": err.Error(),
		})
	}
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.levelBaseSessionModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		total_level_base_session_minute int,             
		total_level_base_session_count int,               
		first_level_session_level_index smallint,               
		first_level_session_duration smallint,                
		first_level_session_year_of_day smallint,               
		first_level_session_year smallint,                     
		first_level_session_week_day smallint,                 
		first_level_session_hour smallint,                     
		first_level_session_minute smallint,                  
		second_level_session_level_index smallint,             
		second_level_session_duration smallint,              
		third_level_session_level_index smallint,              
		third_level_session_duration smallint,                
		four_level_session_level_index smallint,               
		four_level_session_duration smallint,                 
		five_level_session_level_index smallint,               
		five_level_session_duration smallint,                 
		six_level_session_level_index smallint,                
		six_level_session_duration smallint,                  
		seven_level_session_level_index smallint,              
		seven_level_session_duration smallint,               
		first_quarter_hour_total_level_base_session_count smallint,
		first_half_hour_total_level_base_session_count smallint,   
		first_hour_total_level_base_session_count smallint,       
		first_two_hour_total_level_base_session_count smallint,   
		first_three_hour_total_level_base_session_count smallint, 
		first_six_hour_total_level_base_session_count smallint,   
		first_twelve_hour_total_level_base_session_count smallint,
		first_day_total_level_base_session_count smallint,       
		penultimate_level_session_level_index smallint,         
		penultimate_level_session_level_duration smallint,     
		last_level_session_level_index smallint,               
		last_level_session_level_duration smallint,            
		last_level_session_year_of_day smallint,                
		last_level_session_year smallint,                     
		last_level_session_week_day smallint,                  
		last_level_session_hour smallint,                     
		last_level_session_minute smallint)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create levelBaseSessionModels table err: ": err.Error(),
		})
	}

	return &session, err
}
