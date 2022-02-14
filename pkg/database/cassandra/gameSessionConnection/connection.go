package gameSessionConnection

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
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.gameSessionModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		first_session_year_of_day smallint,                                      
		first_session_year smallint,                                          
		first_session_week_day smallint,                                       
		first_session_hour smallint,                                          
		first_session_duration smallint,                                       
		first_session_minute smallint,                                        
		second_session_hour smallint,                                         
		second_session_duration smallint,                                     
		second_session_minute smallint,                                       
		third_session_hour smallint,                                           
		third_session_duration smallint,                                      
		third_session_minute smallint,                                         
		fourth_session_hour smallint,                                         
		fourth_session_duration smallint,                                     
		fourth_session_minute smallint,                                        
		fifth_session_hour smallint,                                           
		fifth_session_duration smallint,                                      
		fifth_session_minute smallint,                                         
		penultimate_session_hour smallint,                                     
		penultimate_session_duration smallint,  
		penultimate_session_minute smallint,                                 
		last_session_year_of_day smallint,                                      
		last_session_year smallint,                                           
		last_session_hour smallint,                                           
		last_session_duration smallint,                                       
		last_session_minute smallint,                                          
		last_duration_minus_penultimate_duration smallint,                      
		first_half_hour_total_session_count smallint,                             
		first_half_hour_total_session_duration smallint,                         
		first_hour_total_session_count smallint,                                
		first_hour_total_session_duration smallint,                             
		first_two_hour_total_session_count smallint,                             
		first_two_hour_total_session_duration smallint,                          
		first_three_hour_total_session_count smallint,                           
		first_three_hour_total_session_duration smallint,                         
		first_six_hour_total_session_count smallint,                              
		first_six_hour_total_session_duration smallint,                          
		first_twelve_hour_total_session_count smallint,                          
		first_twelve_hour_total_session_duration smallint,                       
		total_session_day int,                                            
		total_session_hour int,                                          
		total_session_minute int,                                        
		total_session_duration int,                                      
		total_session_count int,                                         
		first_day_total_session_count smallint,                                 
		first_day_total_session_duration smallint,                              
		second_day_total_session_count smallint,                                
		second_day_total_session_duration smallint,                             
		third_day_total_session_count smallint,                                 
		third_day_total_session_duration smallint,                             
		fourth_day_total_session_count smallint,                                
		fourth_day_total_session_duration smallint,                             
		fifth_day_total_session_count smallint,                                 
		fifth_day_total_session_duration smallint,                              
		sixth_day_total_session_count smallint,                                 
		sixth_day_total_session_duration smallint,                              
		seventh_day_total_session_dount smallint,                               
		seventh_day_total_session_duration smallint,                            
		min_session_duration smallint,                                        
		max_session_duration smallint,                                        
		daily_avegare_session_count float,                                  
		daily_average_session_duration float,                               
		session_based_avegare_session_duration float,                        
		daily_avegare_session_count_minus_first_day_session_count float,         
		daily_avegare_session_duration_minus_first_day_session_duration float,   
		session_based_avegare_session_duration_minus_first_session_duration float,
		session_based_avegare_session_duration_minus_last_session_duration float,
		sunday_session_count smallint,                                        
		monday_session_count smallint,                                        
		tuesday_session_count smallint,                                       
		wednesday_session_count smallint,                                     
		thursday_session_count smallint,                                      
		friday_session_count smallint,                                        
		saturday_session_count smallint,                                      
		am_session_count smallint,                                            
		pm_session_count smallint,                                            
		session_0_to_5_hour_count smallint,                                      
		session_6_to_11_hour_count smallint,                                      
		session_12_to_17_hour_count smallint,                                     
		session_18_to_23_hour_count smallint)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create gameSessionModels table err: ": err.Error(),
		})
	}

	return &session, err
}
