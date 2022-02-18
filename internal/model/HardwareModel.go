package model

import "time"

type HardwareResponseModel struct {
	Id                 int64
	ClientId           int64
	ProjectId          int64
	CustomerId         int64
	DeviceType         int16
	GraphicsDeviceType int16
	GraphicsMemorySize int16
	OperatingSystem    int16
	ProcessorCount     int16
	ProcessorType      int16
	SystemMemorySize   int16
	Status             bool
}

type HardwareModel struct {
	Id                    int64
	ClientId              int64
	ProjectId             int64
	CustomerId            int64
	DeviceModel           string
	DeviceName            string
	DeviceType            int
	GraphicsDeviceName    string
	GraphicsDeviceType    int
	GraphicsDeviceVendor  string
	GraphicsDeviceVersion string
	GraphicsMemorySize    int
	OperatingSystem       string
	ProcessorCount        int
	ProcessorType         string
	SystemMemorySize      int
	DateTime              time.Time
	Status                bool
}
