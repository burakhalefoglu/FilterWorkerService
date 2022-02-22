package cassandra

var tableQueries = [9]string{
	`CREATE TABLE IF NOT EXISTS ClientDatabase.adv_event_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		sixth_adv_year_of_day smallint, 
		sixth_adv_hour smallint,      
		sixth_adv_minute smallint,    
		sixth_adv_type blob,      
		seventh_adv_year_of_day smallint,
		seventh_adv_hour smallint,    
		seventh_adv_minute smallint,  
		seventh_adv_type blob,                               
		penultimate_adv_year_of_day smallint,              
		penultimate_adv_hour smallint,                     
		penultimate_adv_minute smallint,                   
		penultimate_adv_type blob,                     
		last_adv_year_of_day smallint,                     
		last_adv_year smallint,                            
		last_adv_click_hour smallint,                      
		last_adv_click_minute smallint,                    
		last_adv_type blob,
		first_five_minutes_adv_click_count smallint,
		first_ten_minutes_adv_click_count smallint,
		first_quarter_hour_adv_click_count smallint,                           
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
		adv_click_18_to_23_hour_count smallint, 
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.buying_event_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		sixth_buying_year_of_day smallint,   
		sixth_buying_hour smallint,       
		sixth_buying_minute smallint,      
		sixth_buying_product_type blob, 
		seventh_buying_year_of_day smallint, 
		seventh_buying_hour smallint,      
		seventh_buying_minute smallint,    
		seventh_buying_product_type blob,    
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
		level_based_average_buying_count float,
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.game_session_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		sixth_session_hour smallint,     
		sixth_session_duration smallint, 
		sixth_session_minute smallint,   
		seventh_session_hour smallint,   
		seventh_session_duration smallint,
		seventh_session_minute smallint,                                         
		penultimate_session_hour smallint,                                     
		penultimate_session_duration smallint,  
		penultimate_session_minute smallint,                                 
		last_session_year_of_day smallint,                                      
		last_session_year smallint,                                           
		last_session_hour smallint,                                           
		last_session_duration smallint,                                       
		last_session_minute smallint,                                          
		last_duration_minus_penultimate_duration smallint,
		first_five_minutes_total_session_count smallint,  
		first_five_minutes_total_session_duration smallint,
		first_ten_minutes_total_session_count smallint,   
		first_ten_minutes_total_session_duration smallint,
		first_quarter_hour_total_session_count smallint,  
		first_quarter_hour_total_session_duration smallint,                     
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
		session_18_to_23_hour_count smallint, 
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.level_base_session_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		first_five_minutes_total_level_base_session_count smallint,
		first_ten_minutes_total_level_base_session_count smallint,             
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
		last_level_session_minute smallint,
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.hardware_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
		device_type smallint,        
		graphics_device_type smallint,
		graphics_memory_size smallint,
		operating_system smallint,  
		processor_count smallint,   
		processor_type smallint,    
		system_memory_size smallint,
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.screen_click_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		sixth_click_session_hour smallint,   
		sixth_click_session_minute smallint, 
		sixth_touch_count smallint,         
		seventh_click_session_hour smallint, 
		seventh_click_session_minute smallint,
		seventh_touch_count smallint,                          
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
		sixth_start_x_cor float,  
		sixth_start_y_cor float,   
		sixth_finish_x_cor float, 
		sixth_finish_y_cor float, 
		seventh_start_x_cor float,
		seventh_start_y_cor float,
		seventh_finish_x_cor float,
		seventh_finish_y_cor float,                              
		penultimate_start_x_cor float,                           
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
		last_touch_count_minus_session_based_avegare_click_count float, 
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.screen_swipe_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
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
		sixth_swipe_direction blob,  
		sixth_swipe_start_x_cor float,  
		sixth_swipe_start_y_cor float,  
		sixth_swipe_finish_x_cor float, 
		sixth_swipe_finish_y_cor float, 
		seventh_swipe_direction blob,
		seventh_swipe_start_x_cor float,
		seventh_swipe_start_y_cor float,
		seventh_swipe_finish_x_cor float,
		seventh_swipe_finish_y_cor float,        
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
		total_swipe_finish_y_cor float,
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.location_models(
		id bigint, 
		client_id bigint, 
		project_id bigint, 
		customer_id bigint, 
		continent smallint,
		country smallint, 
		city smallint,    
		region smallint,  
		org smallint, 
		status boolean, 
		PRIMARY KEY((client_id, project_id)))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.type_standardization_models(
		key text,
		value smallint,
		PRIMARY KEY(key))`,
}

func GetTableQueries() [9]string {
	return tableQueries
}
