package mediator

type Mediator interface {
	notify(sender interface{})
}

type ConcreteMediator struct {
	ComponentA
	ComponentB
	ComponentC
	ComponentD
}

func (c ConcreteMediator) notify(sender interface{}) {
}

func (c ConcreteMediator) reactOnA() {
}

func (c ConcreteMediator) reactOnB() {
}

func (c ConcreteMediator) reactOnC() {
}

func (c ConcreteMediator) reactOnD() {
}

type ComponentA struct {
	m Mediator
}

func (c ComponentA) operationA() {
}

type ComponentB struct {
	m Mediator
}

func (c ComponentB) operationB() {
}

type ComponentC struct {
	m Mediator
}

func (c ComponentC) operationC() {
}

type ComponentD struct {
	m Mediator
}

func (c ComponentD) operationD() {
}
