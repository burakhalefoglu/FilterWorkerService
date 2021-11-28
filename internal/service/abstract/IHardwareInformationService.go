package abstract

type IHardwareInformationService interface {
	AddHardwareInformation(data *[]byte) (v interface{}, s bool, m string)
}