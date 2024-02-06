package command

type Command interface {
	execute()
}

type Invoker struct {
	command Command
}

func (i Invoker) setCommand(command Command) {
}

func (i Invoker) executeCommand() {
}

type Receiver struct {
}

func (r Receiver) operation(a, b, c interface{}) {
}

type ConcreteCommand1 struct {
	receiver Receiver
	params   interface{}
}

func (cc ConcreteCommand1) execute() {
}

type ConcreteCommand2 struct {
	receiver Receiver
	params   interface{}
}

func (cc ConcreteCommand2) execute() {
}
