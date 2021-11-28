package abstract

type IScreenSwipeService interface {
	ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string)
}
