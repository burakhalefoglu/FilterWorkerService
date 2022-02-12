package model

type HardwareInformationResponseModel struct {
	ProjectId          string
	ClientId           string
	CustomerId         string
	DeviceType         int16
	GraphicsDeviceType int16
	GraphicsMemorySize int16
	OperatingSystem    int16
	ProcessorCount     int16
	ProcessorType      int16
	SystemMemorySize   int16
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
