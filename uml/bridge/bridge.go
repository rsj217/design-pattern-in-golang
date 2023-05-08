package bridge

type Implementation interface {
	method1()
	method2()
	method3()
}

type Abstraction struct {
	i Implementation
}

func (a Abstraction) feature1() {
}
func (a Abstraction) feature2() {
}

type RefinedAbstraction struct {
	Abstraction
}

func (ra RefinedAbstraction) featureN() {
}

type ConcreteImplementation struct {
}

func (c ConcreteImplementation) method1() {
}

func (c ConcreteImplementation) method2() {
}

func (c ConcreteImplementation) method3() {
}
