package strategy

import "fmt"

type Notifyer interface {
	Send(msg string) bool
}

var _ Notifyer = (*Push)(nil)

type Push struct {
	uid int64
}

func (p *Push) getDeviceIdBy(uid int64) string {
	return fmt.Sprintf("%d-ios-14-10010", uid)
}

func (p *Push) Send(msg string) bool {
	deviceId := p.getDeviceIdBy(p.uid)
	fmt.Printf("%s Send push %s\n", deviceId, msg)
	return true
}

var _ Notifyer = (*SMS)(nil)

type SMS struct {
	uid int64
}

func (s *SMS) getPhoneBy(uid int64) string {
	return "18519191001"
}

func (s *SMS) Send(msg string) bool {
	phone := s.getPhoneBy(s.uid)
	fmt.Printf("%s Send sms %s\n", phone, msg)
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
	push := &Push{110110}
	notifyCtx := NotifyCtx{}
	notifyCtx.SetNotifyer(push)
	notifyCtx.Send("hello world")

	sms := &SMS{110110}
	notifyCtx.SetNotifyer(sms)
	notifyCtx.Send("hello world")
}
