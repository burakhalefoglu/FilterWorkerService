package abstract

type IBuyingEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
