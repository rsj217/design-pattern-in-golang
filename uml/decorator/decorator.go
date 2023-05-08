package decorator

type Component interface {
	execute()
}

type ConcreteComponent struct {
}

func (cc ConcreteComponent) execute() {
}

type BaseDecorator struct {
	wrappee Component
}

func (bd BaseDecorator) execute() {
}

type ConcreteDecorators struct {
}

func (cd ConcreteDecorators) execute() {
}

func (cc ConcreteComponent) extra() {
}
