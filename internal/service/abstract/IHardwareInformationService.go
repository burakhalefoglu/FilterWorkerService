package abstract

type IHardwareService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}