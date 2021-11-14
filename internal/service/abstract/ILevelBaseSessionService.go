package abstract



type ILevelBaseSessionService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
