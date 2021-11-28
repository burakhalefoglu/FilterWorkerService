package abstract

type ILocationService interface {
	AddLocation(data *[]byte) (v interface{}, s bool, m string)
}
