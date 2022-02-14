package advEventConnection

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
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.advEventModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		level_index smallint,
		total_adv_day int,
		total_adv_count int,
		total_adv_hour int,
		total_adv_minute int,
		level_based_average_adv_count float,
		average_adv_daily_click_count float,
		first_adv_year_of_day smallint,
		first_adv_year smallint,                         
		first_week_day smallint,                           
		first_adv_click_hour smallint,                     
		first_adv_click_minute smallint,                   
		first_adv_type blob,                           
		second_adv_year_of_day smallint,                   
		second_adv_hour smallint,                          
		second_adv_minute smallint,                         
		second_adv_type blob,                          
		third_adv_year_of_day smallint,                    
		third_adv_hour smallint,                           
		third_adv_minute smallint,                         
		third_adv_type blob,                           
		fourth_adv_year_of_day smallint,                   
		fourth_adv_hour smallint,                          
		fourth_adv_minute smallint,                        
		fourth_adv_type blob,                          
		fifth_adv_year_of_day smallint,                    
		fifth_adv_hour smallint,                           
		fifth_adv_minute smallint,                         
		fifth_adv_type blob,                           
		penultimate_adv_year_of_day smallint,              
		penultimate_adv_hour smallint,                     
		penultimate_adv_minute smallint,                   
		penultimate_adv_type blob,                     
		last_adv_year_of_day smallint,                     
		last_adv_year smallint,                            
		last_adv_click_hour smallint,                      
		last_adv_click_minute smallint,                    
		last_adv_type blob,                            
		first_half_hour_adv_click_count smallint,           
		first_hour_adv_click_count smallint,               
		first_two_hour_adv_click_count smallint,           
		firs_three_hour_adv_click_count smallint,          
		first_six_hour_adv_click_count smallint,           
		first_twelve_hour_adv_click_count smallint,        
		first_day_adv_click_count smallint,                
		second_day_adv_click_count smallint,               
		third_day_adv_click_count smallint,                
		fourth_day_adv_click_count smallint,               
		fifth_day_adv_click_count smallint,                
		sixth_day_adv_click_count smallint,                
		seventh_day_adv_click_count smallint,               
		penultimate_day_adv_click_count smallint,          
		last_day_adv_click_count smallint,                  
		last_minus_first_day_adv_click_count smallint,      
		last_minus_penultimate_day_adv_click_count smallint,
		last_day_adv_click_count_minus_average_daily_adv_click_count smallint,
		sunday_adv_click_count smallint,       
		monday_adv_click_count smallint,      
		tuesday_adv_click_count smallint,     
		wednesday_adv_click_count smallint,    
		thursday_adv_click_count smallint,    
		friday_adv_click_count smallint,       
		saturday_adv_click_count smallint,    
		am_adv_click_count smallint,          
		pm_adv_click_count smallint,          
		adv_click_0_to_5_hour_count smallint, 
		adv_click_6_to_11_hour_count smallint,
		adv_click_12_to_17_hour_count smallint,
		adv_click_18_to_23_hour_count smallint)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create advEventModels table err: ": err.Error(),
		})
	}

	return &session, err
}
