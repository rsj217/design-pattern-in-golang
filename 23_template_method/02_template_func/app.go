package template_method

import "fmt"

type GetIdentifyer func(int642 int64) string
type Sender func(string, string) bool
type Saver func(interface{})

func PushGetIdentify(uid int64) string {
	return fmt.Sprintf("%d-ios-14-110100", uid)
}

func PushSend(msg, identify string) bool {
	fmt.Printf("%s Send push %s\n", identify, msg)
	return true
}

func PushSave(info interface{}) {
	fmt.Printf("Push save  %v to db\n", info)
}

func SMSGetIdentify(uid int64) string {
	return fmt.Sprintf("18519191001-%d", uid)
}

func SMSSend(msg, identify string) bool {
	fmt.Printf("%s Send sms %s\n", identify, msg)
	return true
}

func SMSSave(info interface{}) {
	fmt.Printf("SMS record info %v log\n", info)
}

func MakeNotify(uid int64, msg string, getIdentify GetIdentifyer, sender Sender, saver Saver) bool {
	fmt.Println("Make Notify...")
	identify := getIdentify(uid)
	ans := sender(msg, identify)
	saver(ans)
	return ans
}

func Client() {
	flag := "push"
	msg := "hello world"
	uid := int64(100100)

	switch flag {
	case "push":
		MakeNotify(uid, msg, PushGetIdentify, PushSend, PushSave)
	case "sms":
		MakeNotify(uid, msg, SMSGetIdentify, SMSSend, SMSSave)
	}
}
