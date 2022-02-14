package screenSwipeConnection

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
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS MLDatabase.screenSwipeModels (
		project_id text PRIMARY KEY,
		client_id text
		customer_id text,
		level_index smallint,                                     
		total_swipe_session_count int,       
		total_swipe_hour int,               
		first_swipe_year_of_day smallint,          
		first_swipe_year smallint,               
		first_swipe_hour smallint,               
		first_swipe_week_day smallint,             
		first_swipe_minute smallint,             
		fist_swipe_direction blob,           
		first_swipe_start_x_cor float,           
		first_swipe_start_y_cor float,          
		first_swipe_finish_x_cor float,         
		first_swipe_finish_y_cor float,         
		second_swipe_direction blob,         
		second_swipe_start_x_cor float,          
		second_swipe_start_y_cor float,         
		second_swipe_finish_x_cor float,        
		second_swipe_finish_y_cor float,        
		third_swipe_direction blob,          
		third_swipe_start_x_cor float,          
		third_swipe_start_y_cor float,          
		third_swipe_finish_x_cor float,         
		third_swipe_finish_y_cor float,         
		fourth_swipe_direction blob,         
		fourth_swipe_start_x_cor float,         
		fourth_swipe_start_y_cor float,         
		fourth_swipe_finish_x_cor float,        
		fourth_swipe_finish_y_cor float,        
		fifth_swipe_direction blob,          
		fifth_swipe_start_x_cor float,          
		fifth_swipe_start_y_cor float,          
		fifth_swipe_finish_x_cor float,         
		fifth_swipe_finish_y_cor float,         
		penultimate_swipe_direction blob,    
		penultimate_swipe_start_x_cor float,    
		Penultimate_swipe_start_y_cor float,    
		penultimate_swipe_finish_x_cor float,    
		penultimate_swipe_finish_y_cor float,   
		penultimate_swipe_year_of_day smallint,    
		penultimate_swipe_year smallint,         
		penultimate_swipe_hour smallint,         
		penultimate_swipe_week_day smallint,      
		penultimate_swipe_minute smallint,       
		last_swipe_direction blob,           
		last_swipe_start_x_cor float,           
		last_swipe_start_y_cor float,            
		last_swipe_finish_x_cor float,          
		last_swipe_finish_y_cor float,           
		last_swipe_year_of_day smallint,           
		last_swipe_year smallint,                
		last_swipe_hour smallint,                
		last_swipe_week_day smallint,             
		last_swipe_minute smallint,              
		first_day_total_swipe_up_count int,    
		first_day_total_swipe_down_count int,  
		first_day_total_swipe_right_count int, 
		first_day_total_swipe_left_count int,   
		first_day_swipe_total_start_x_cor float,  
		first_day_swipe_total_start_y_cor float,  
		first_day_swipe_total_finish_x_cor float, 
		first_day_swipe_total_finish_y_cor float, 
		second_day_total_swipe_up_count int,   
		second_day_total_swipe_down_count int, 
		second_day_total_swipe_right_count int,
		second_day_total_swipe_left_count int, 
		second_day_swipe_total_start_x_cor float,  
		second_day_swipe_total_start_y_cor float, 
		second_day_swipe_total_finish_x_cor float,
		second_day_swipe_total_finish_y_cor float,
		third_day_total_swipe_up_count int,    
		third_day_total_swipe_down_count int,  
		third_day_total_swipe_right_count int, 
		third_day_total_swipe_left_count int,   
		third_day_Swipetotal_start_x_cor float,  
		third_day_Swipetotal_start_y_cor float,  
		third_day_Swipetotal_finish_x_cor float,  
		third_day_Swipetotal_finish_y_cor float, 
		fourth_day_total_swipe_up_count int,   
		fourth_day_total_swipe_down_count int, 
		fourth_day_total_swipe_right_count int,
		fourth_day_total_swipe_left_count int, 
		fourth_day_swipe_total_start_x_cor float, 
		fourth_day_swipe_total_start_y_cor float, 
		fourth_day_swipe_total_finish_x_cor float,
		fourth_day_swipe_total_finish_y_cor float,
		fifth_day_total_swipe_up_count int,    
		fifth_day_total_swipe_down_count int,  
		fifth_day_total_swipe_right_count int, 
		fifth_day_total_swipe_left_count int,  
		fifth_day_swipe_total_start_x_cor float,  
		fifth_day_swipe_total_start_y_cor float,  
		fifth_day_swipe_total_finish_x_cor float, 
		fifth_day_swipe_total_finish_y_cor float, 
		sixth_day_total_swipe_up_count int,    
		sixth_day_total_swipe_down_count int,  
		sixth_day_total_swipe_right_count int, 
		sixth_day_total_swipe_left_count int,  
		sixth_day_swipe_total_start_x_cor float,  
		sixth_day_swipe_total_start_y_cor float,  
		sixth_day_swipe_total_finish_x_cor float, 
		sixth_day_swipe_total_finish_y_cor float, 
		seventh_day_total_swipe_up_count int,  
		seventh_day_total_swipe_down_count int,
		seventh_day_total_swipe_right_count int,
		seventh_day_total_swipe_left_count int,
		seventh_day_swipe_total_start_x_cor float,
		seventh_day_swipe_total_start_y_cor float,
		seventh_day_swipe_total_finish_x_cor float,
		seventh_day_swipe_total_finish_y_cor float,
		total_swipe_up_count int,            
		total_swipe_down_count int,          
		total_swipe_right_count int,         
		total_swipe_left_count int,          
		total_swipe_start_x_cor float,          
		total_swipe_start_y_cor float,          
		total_swipe_finish_x_cor float,         
		total_swipe_finish_y_cor float)`)
	if err != nil {
		clogger.Error(&logger.Messages{
			"create screenSwipeModels table err: ": err.Error(),
		})
	}

	return &session, err
}
