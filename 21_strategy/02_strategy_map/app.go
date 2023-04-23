package strategy

import "fmt"

func PushSend(deviceId, msg string) bool {
	fmt.Printf("%s Send push %s\n", deviceId, msg)
	return true
}

func SMSSend(phone, msg string) bool {
	fmt.Printf("%s Send sms %s\n", phone, msg)
	return true
}

type NStrategy string

const (
	PushStrategy NStrategy = "push"
	SMSStrategy            = "sms"
)

var NotifyCtx = map[NStrategy]func(identify, msg string) bool{
	PushStrategy: PushSend,
	SMSStrategy:  SMSSend,
}

func client() {
	deviceId := "ios-14-10010"
	phone := "18519191001"
	msg := "hello world"

	flag := PushStrategy
	NotifyCtx[flag](deviceId, msg)

	flag = SMSStrategy
	NotifyCtx[flag](phone, msg)
}
