package abstract

type IScreenSwipeService interface {
	ConvertRawModelToResponseModel(data *[]byte) (s bool, m string)
}
