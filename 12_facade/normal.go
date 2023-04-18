package facade

import "fmt"

func checkStock(actId int64) bool {
	fmt.Printf("check actId=%d left stock\n", actId)
	return true
}

func checkLogin(session string) bool {
	fmt.Printf("check session=%s login status\n", session)
	return true
}

func checkRisk(device string) bool {
	fmt.Printf("check device=%s risk status\n", device)
	return true
}

func SendCoupon(actId int64, session string, device string) bool {
	if !checkLogin(session) {
		return false
	}
	if !checkStock(actId) {
		return false
	}
	if !checkRisk(device) {
		return false
	}
	return true
}

func App() bool {
	actId := int64(101)
	session := "session217"
	device := "ios171"
	SendCoupon(actId, session, device)
	fmt.Printf("send coupon")
	return true
}
