package model

type HardwareInformationResponseModel struct {
	ProjectId          string
	ClientId           string
	CustomerId         string
	DeviceType         int64
	GraphicsDeviceType int64
	GraphicsMemorySize int64
	OperatingSystem    int64
	ProcessorCount     int64
	ProcessorType      int64
	SystemMemorySize   int64
}

type HardwareInformationModel struct {
	ProjectId          string
	ClientId           string
	CustomerId         string
	DeviceType         int
	GraphicsDeviceType int
	GraphicsMemorySize int
	OperatingSystem    string
	ProcessorCount     int
	ProcessorType      string
	SystemMemorySize   int
}
