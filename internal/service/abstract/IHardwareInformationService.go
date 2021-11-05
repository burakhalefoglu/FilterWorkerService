package abstract

type IHardwareInformationService interface {
	AddHardwareInformation(data *[]byte) (s bool, m string)
}