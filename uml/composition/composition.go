package composition

type Component interface {
	execute()
}

type Composite struct {
	children []*Component
}

func (cc Composite) execute() {
}

func (cc Composite) add() {
}

func (cc Composite) remove() {
}
func (cc Composite) getChildren() []Component {
	return nil
}

type Leaf struct {
}

func (l Leaf) execute() {
}

type Client struct {
}
