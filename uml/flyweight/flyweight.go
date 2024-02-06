package flyweight

type Context struct {
	uniqueState interface{}
	flyweight   Flyweight
}

func (c Context) operation() {

}

type Flyweight struct {
	repeatState interface{}
}

func (c Flyweight) operation() {

}

type FlyweightFactory struct {
	cache []Flyweight
}

func (ff FlyweightFactory) getFlyweight() {

}
