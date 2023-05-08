package facade

type Client struct {
}

type Facade struct {
	linksToSubsystemObjects interface{}
	opionalAdditionalFacade interface{}
}

func (f Facade) subsystemOperation() {

}

type AdditionalFacde struct {
}

func (af AdditionalFacde) anotherOperation() {

}
