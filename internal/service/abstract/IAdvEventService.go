package abstract

type IAdvEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}
