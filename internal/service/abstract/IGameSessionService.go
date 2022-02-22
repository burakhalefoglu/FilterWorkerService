package abstract

type IGameSessionService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}
