package template_method

import "fmt"

type Notifyer interface {
	GetIdentify() string
	Send(string, string) bool
	Save(interface{})
}

var _ Notifyer = (*Push)(nil)

type Push struct {
	uid int64
}

func (p *Push) RecordLog() {
	//TODO implement me
	panic("implement me")
}

func (p *Push) GetIdentify() string {
	return fmt.Sprintf("%d-ios-14-110100", p.uid)
}

func (p *Push) Send(msg, identify string) bool {
	fmt.Printf("%s Send push %s\n", identify, msg)
	return true
}

func (p *Push) Save(info interface{}) {
	fmt.Printf("Push save  %v to db\n", info)
}

var _ Notifyer = (*SMS)(nil)

type SMS struct {
	uid int64
}

func (s *SMS) GetIdentify() string {
	return fmt.Sprintf("18519191001-%d", s.uid)
}

func (s *SMS) Send(msg, identify string) bool {
	fmt.Printf("%s Send sms %s\n", identify, msg)
	return true
}

func (s *SMS) Save(info interface{}) {
	fmt.Printf("SMS record info %v log\n", info)
}

func MakeNotify(msg string, notifyer Notifyer) bool {
	fmt.Println("Make Notify...")
	identify := notifyer.GetIdentify()
	ans := notifyer.Send(msg, identify)
	notifyer.Save(ans)
	return ans
}

func Client() {
	flag := "push"
	msg := "hello world"
	uid := int64(100100)

	switch flag {
	case "push":
		push := &Push{uid}
		MakeNotify(msg, push)
	case "sms":
		sms := &SMS{uid}
		MakeNotify(msg, sms)
	}
}
