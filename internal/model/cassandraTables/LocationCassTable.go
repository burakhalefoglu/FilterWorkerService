package cassandraTables

import "github.com/scylladb/gocqlx/v2/table"

var locationMetadata = table.Metadata{
	Name: "MLDatabase.locationModels",
	Columns: []string{
		"project_id",
		"client_id",
		"customer_id",
		"continent",               
		"country",                
		"city",               
		"region",                 
	},
	PartKey: []string{"project_id"},
	SortKey: []string{"client_id"},
}

var LocationTable = table.New(locationMetadata)