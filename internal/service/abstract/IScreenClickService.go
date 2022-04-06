package abstract

type IScreenClickService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
