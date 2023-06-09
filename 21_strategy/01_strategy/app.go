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
	return fmt.Sprintf("%d-ios-14-110100", uid)
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

type PushAndSMS struct {
	uid int64
	*Push
	*SMS
}

func NewPushAndSMS(uid int64) *PushAndSMS {
	return &PushAndSMS{
		uid,
		&Push{uid: uid},
		&SMS{uid: uid},
	}
}

func (ps *PushAndSMS) Send(msg string) bool {
	ps.Push.Send(msg)
	ps.SMS.Send(msg)
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

func client() {
	flag := "both"
	notifyCtx := NotifyCtx{}
	msg := "hello world"

	switch flag {
	case "push":
		push := &Push{110110}
		notifyCtx.SetNotifyer(push)
		notifyCtx.Send(msg)
	case "sms":
		sms := &SMS{110110}
		notifyCtx.SetNotifyer(sms)
		notifyCtx.Send(msg)
	default:
		both := NewPushAndSMS(110110)
		notifyCtx.SetNotifyer(both)
		notifyCtx.Send(msg)
	}
}
