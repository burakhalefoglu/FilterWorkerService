package abstract

type IAdvEventService interface {
	AddAdvEvent(data *[]byte) (s bool, m string)
}