package mongodb

import (
	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "Client", options.Client().ApplyURI("mongodb://localhost:27017"))
	if(err != nil){
		panic(err)
	}
}

