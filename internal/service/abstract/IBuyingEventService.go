package abstract

type IBuyingEventService interface {
	AddBuyingEvent(data *[]byte) (s bool, m string)
}