package abstract

type IHardwareService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
