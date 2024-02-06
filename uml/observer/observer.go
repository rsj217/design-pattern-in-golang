package observer

type Subscriber interface {
	update(context interface{})
}

type ConcreteSubscribers struct {
}

func (cs ConcreteSubscribers) update(context interface{}) {

}

type Publisher struct {
	subscriblers []Subscriber
	mainState    interface{}
}

func (p Publisher) subscribe(s Subscriber) {
}

func (p Publisher) unsubscribe(s Subscriber) {
}
func (p Publisher) notifySubscribers() {
}
func (p Publisher) mainBusinessLogic() {
}

type Client struct {
}
