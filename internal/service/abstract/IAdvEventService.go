package abstract


type IAdvEventService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
