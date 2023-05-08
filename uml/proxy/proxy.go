package proxy

type ServiceInterface interface {
	operation()
}

type Proxy struct {
	realService Service
}

func (p Proxy) checkAccess() {

}

func (p Proxy) operation() {

}

type Service struct {
}

func (s Service) operation() {

}

type Client struct {
}
