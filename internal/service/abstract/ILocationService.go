package abstract

type ILocationService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
