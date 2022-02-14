package cassandraTables

import "github.com/scylladb/gocqlx/v2/table"

var typeStandardizationMetadata = table.Metadata{
	Name: "MLDatabase.typeStandardizationModels",
	Columns: []string{
		"key",
		"value",		
	},
	PartKey: []string{"key"},
	SortKey: []string{"key"},
}

var TypeStandardizationTable = table.New(typeStandardizationMetadata)
