package buyingEventConnection

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
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.buyingEventModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		level_index smallint,
		total_buying_count int,
		total_buying_day int,         
		total_buying_hour int,        
		first_buying_year_of_day smallint,   
		first_buying_year smallint,        
		first_buying_hour smallint,        
		first_buying_minute smallint,      
		first_buying_product_type blob, 
		second_buying_year_of_day smallint,  
		second_buying_hour smallint,       
		second_buying_minute smallint, 
		second_buying_product_type blob,
		third_buying_year_of_day smallint,       
		third_buying_hour smallint,            
		third_buying_minute smallint,          
		third_buying_product_type blob,     
		fourth_buying_year_of_day smallint,      
		fourth_buying_hour smallint,           
		fourth_buying_minute smallint,         
		fourth_buying_product_type blob,    
		fifth_buying_year_of_day smallint,       
		fifth_buying_hour smallint,            
		fifth_buying_minute smallint,          
		fifth_buying_product_type blob,     
		penultimate_buying_year_of_day smallint, 
		penultimate_buying_hour smallint,       
		penultimate_buying_minute smallint,    
		penultimate_buying_product_type blob,
		last_buying_year_of_day smallint,        
		last_buying_year smallint,              
		last_buying_hour smallint,             
		last_buying_minute smallint,           
		last_buying_product_type blob,
		first_day_buying_count smallint, 
		second_day_buying_count smallint,
		third_day_buying_count smallint, 
		fourth_day_buying_count smallint,
		fifth_day_buying_count smallint, 
		sixth_day_buying_count smallint, 
		seventh_day_buying_count smallint,
		sunday_buying_count smallint,   
		monday_buying_count smallint,   
		tuesday_buying_count smallint,  
		wednesday_buying_count smallint,
		thursday_buying_count smallint, 
		friday_buying_count smallint,   
		saturday_buying_count smallint, 
		am_buying_count smallint,       
		pm_buying_count smallint,      
		buying_0_to_5_hour_count smallint, 
		buying_6_to_11_hour_count smallint,
		buying_12_to_17_hour_count smallint,
		buying_18_to_23_hour_count smallint,
		buying_day_average_buying_count float,
		level_based_average_buying_count float)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create buyingEventModels table err: ": err.Error(),
		})
	}

	return &session, err
}
