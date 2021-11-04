package abstract

type ILocationService interface {
	AddLocation(data *[]byte) (s bool, m string)
}
