package abstract

type IGameSessionService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
