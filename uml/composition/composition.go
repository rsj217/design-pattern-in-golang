package composition

type Card struct {
	userCoupons []*UserCoupon
}

type UserCoupon struct {
}
