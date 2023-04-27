package mediator

import "fmt"

type Broker interface {
	SendMsg(Consumer, string)
}

type Consumer interface {
	GetName() string
	RegBroker(Broker)
	Recv(msg string)
	Send(msg string)
}

var _ Broker = (*HsfBroker)(nil)

type HsfBroker struct {
	consumer map[string]Consumer
}

func (hb HsfBroker) SendMsg(consumer Consumer, msg string) {
	for k, v := range hb.consumer {
		if k != consumer.GetName() {
			v.Recv(msg)
		}
	}
}

var _ Consumer = (*ICSServer)(nil)

type ICSServer struct {
	name   string
	broker Broker
}

func (i *ICSServer) GetName() string {
	return i.name
}

func (i *ICSServer) RegBroker(broker Broker) {
	i.broker = broker
}

func (i *ICSServer) Recv(msg string) {
	fmt.Printf("%s received msg: %s\n", i.name, msg)
}

func (i *ICSServer) Send(msg string) {
	i.broker.SendMsg(i, msg)
}

var _ Consumer = (*IPCServer)(nil)

type IPCServer struct {
	name   string
	broker Broker
}

func (i *IPCServer) GetName() string {
	return i.name
}

func (i *IPCServer) RegBroker(broker Broker) {
	i.broker = broker
}

func (i *IPCServer) Recv(msg string) {
	fmt.Printf("%s received msg: %s\n", i.name, msg)
}

func (i *IPCServer) Send(msg string) {
	i.broker.SendMsg(i, msg)
}

func client() {
	broker := &HsfBroker{
		consumer: make(map[string]Consumer),
	}

	ics := &ICSServer{name: "ICS"}
	ipc := &IPCServer{name: "IPC"}

	broker.consumer[ics.GetName()] = ics
	broker.consumer[ipc.GetName()] = ipc

	ics.RegBroker(broker)
	ipc.RegBroker(broker)

	ics.Send("hello from ics")
	ipc.Send("world from ipc")

}
