package screenClickConnection

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
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.screenClickModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		level_index smallint,                                     
		first_click_session_year_of_day smallint,                      
		first_click_session_year smallint,                          
		first_click_session_hour smallint,                           
		first_click_session_minute smallint,                        
		first_touch_count smallint,                                
		second_click_session_hour smallint,                         
		second_click_session_minute smallint,                       
		second_touch_count smallint,                               
		third_click_session_hour smallint,                          
		third_click_session_minute smallint,                        
		third_touch_count smallint,                                
		fourth_click_session_hour smallint,                         
		fourth_click_session_minute smallint,                       
		fourth_touch_count smallint,                               
		fifth_click_session_hour smallint,                          
		fifth_click_session_minute smallint,                        
		fifth_touch_count smallint,                                 
		penultimate_click_session_hour smallint,                    
		penultimate_click_session_minute smallint,                  
		penultimate_touch_count smallint,                          
		last_click_session_year_of_day smallint,                      
		last_click_session_year smallint,                           
		last_click_session_hour smallint,                           
		last_click_session_minute smallint,                         
		last_touch_count smallint,                                 
		first_start_x_cor float,                                 
		first_start_y_cor float,                                 
		first_finish_x_cor float,                                
		first_finish_y_cor float,                                
		second_start_x_cor float,                                
		second_start_y_cor float,                                
		second_finish_x_cor float,                               
		second_finish_y_cor float,                               
		third_start_x_cor float,                                 
		third_start_y_cor float,                                 
		third_finish_x_cor float,                                
		third_finish_y_cor float,                                
		fourth_start_x_cor float,                                
		fourth_start_y_cor float,                                
		fourth_finish_x_cor float,                               
		fourth_finish_y_cor float,                               
		fifth_start_x_cor float,                                 
		fifth_start_y_cor float,                                 
		fifth_finish_x_cor float,                                
		fifth_finish_y_cor float,                               
		penultimate_start_x_Cor float,                           
		penultimate_start_y_cor float,                           
		penultimate_finish_x_cor float,                          
		penultimate_finish_y_cor float,                          
		last_start_x_cor float,                                  
		last_start_y_cor float,                                  
		last_finish_x_cor float,                                 
		last_finish_y_cor float,                                  
		first_half_hour_touch_count int,                        
		first_hour_touch_count int,                            
		first_two_hour_touch_count int,                         
		first_three_hour_touch_count int,                       
		first_six_hour_touch_count int,                         
		first_twelve_hour_touch_count int,                      
		first_minus_last_touch_count smallint,                       
		first_finger_id blob,                                  
		penultimate_finger_id blob,                            
		last_finger_id blob,                                   
		first_day_click_count int,                            
		second_day_click_count int,                             
		third_day_click_count int,                             
		fourth_day_click_count int,                             
		fifth_day_click_count int,                             
		sixth_day_click_count int,                             
		seventh_day_click_count int,                           
		total_click_day int,                                  
		total_click_count int,                                
		total_click_session_count int,                         
		total_click_hour int,                                 
		total_click_minute int,                               
		total_start_x_cor float,                                 
		total_start_y_cor float,                                 
		total_finish_x_cor float,                                
		total_finish_y_cor float,                                
		session_based_avegare_start_x_cor float,                   
		session_based_avegare_start_y_cor float,                  
		session_based_avegare_finish_x_cor float,                  
		session_based_avegare_finish_y_cor float,                  
		session_based_avegare_click_count float,                  
		daily_avegare_click_count float,                         
		last_touch_count_minus_session_based_avegare_click_count float)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create screenClickModels table err: ": err.Error(),
		})
	}

	return &session, err
}
