package cor

type Handler interface {
	setNext(handler Handler)
	handle(request interface{})
}

type BaseHandler struct {
	next Handler
}

func (bh BaseHandler) setNext(handler Handler) {

}

func (bh BaseHandler) handle(request interface{}) {

}

type ConcreteHandlers struct {
}

func (bh ConcreteHandlers) handle() {

}
