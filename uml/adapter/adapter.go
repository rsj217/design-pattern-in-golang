package adapter

type ClientInterface interface {
	method(data interface{})
}

type Adapter struct {
	adapter Service
}

func (a Adapter) method(data interface{}) {

}

type Service struct {
}

func serviceMethod(specialData interface{}) {

}

type Client struct {
}
