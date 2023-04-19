package strategy

import "fmt"

type Notifyer interface {
	Send(msg string) bool
}

var _ Notifyer = (*Push)(nil)

type Push struct {
	deviceId string
}

func (p *Push) Send(msg string) bool {
	fmt.Printf("%s Send push %s\n", p.deviceId, msg)
	return true
}

var _ Notifyer = (*SMS)(nil)

type SMS struct {
	phone string
}

func (s *SMS) Send(msg string) bool {
	fmt.Printf("%s Send sms %s\n", s.phone, msg)
	return true
}

type NotifyCtx struct {
	notifyer Notifyer
}

func (nc *NotifyCtx) SetNotifyer(notifyer Notifyer) {
	nc.notifyer = notifyer
}

func (nc *NotifyCtx) Send(msg string) bool {
	return nc.notifyer.Send(msg)
}

func Client() {
	push := &Push{
		"ios-14-10010",
	}
	notifyCtx := NotifyCtx{}
	notifyCtx.SetNotifyer(push)
	notifyCtx.Send("hello world")

	sms := &SMS{
		"18519191001",
	}
	notifyCtx.SetNotifyer(sms)
	notifyCtx.Send("hello world")
}
