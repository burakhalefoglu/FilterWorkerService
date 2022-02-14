package cassandraTables

import "github.com/scylladb/gocqlx/v2/table"

var hardwareInformationMetadata = table.Metadata{
	Name: "MLDatabase.hardwareInformationModels",
	Columns: []string{
		"project_id",
		"client_id",
		"customer_id",
		"device_type",        
		"graphics_device_type",
		"graphics_memory_size",
		"operating_system",  
		"processor_count",   
		"processor_type",    
		"system_memory_size",
	},
	PartKey: []string{"project_id"},
	SortKey: []string{"client_id"},
}

var HardwareInformationTable = table.New(hardwareInformationMetadata)