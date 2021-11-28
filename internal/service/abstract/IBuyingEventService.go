package abstract

type IBuyingEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}
