package abstract

type ILocationService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}
