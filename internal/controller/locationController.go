package controller

import (
	locationManager "FilterWorkerService/internal/manager"
	kafkaAdapter "FilterWorkerService/pkg/kafka"
	"context"
	"fmt"
)

func ListenLocations() {
	ctx := context.Background()
	fmt.Println("ListenLocations is listening")
	kafkaAdapter.Consume(ctx, "LocationDataModel", "LocationDataModel_Filter_ContientConsumerGroup",
		locationManager.AddOrUpdateContient)
	// kafkaAdapter.Consume("LocationDataModel", "LocationDataModel_Filter_CountryConsumerGroup",
	// 	locationManager.AddOrUpdateCountry)
	// kafkaAdapter.Consume("LocationDataModel", "LocationDataModel_Filter_CityConsumerGroup",
	// 	locationManager.AddOrUpdateCity)
	// kafkaAdapter.Consume("LocationDataModel", "LocationDataModel_Filter_CityConsumerGroup",
	// 	locationManager.AddOrUpdateCity)
	// kafkaAdapter.Consume("LocationDataModel", "LocationDataModel_Filter_RegionConsumerGroup",
	// 	locationManager.AddOrUpdateRegion)
	// kafkaAdapter.Consume("LocationDataModel", "LocationDataModel_Filter_OrgConsumerGroup",
	// 	locationManager.AddOrUpdateOrg)
}
