package abstract



type IGameSessionEveryLoginService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
