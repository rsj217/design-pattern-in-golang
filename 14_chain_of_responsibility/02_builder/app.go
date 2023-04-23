package _2_builder

import (
	"errors"
	"fmt"
	"time"
)

type Filter struct {
	adiu        string
	poiid       string
	benefitType int64
	actTime     time.Time
	err         error
}

func (f *Filter) CheckAdiu(adiu string) *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckAdiu")
	if f.adiu != adiu {
		f.err = errors.New("adiu err")
	}
	return f
}

func (f *Filter) CheckPoiid(poiid string) *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckPoiid")
	if f.poiid != poiid {
		f.err = errors.New("poiid err")
	}
	return f
}

func (f *Filter) CheckActTime(actTime time.Time) *Filter {
	if f.err != nil {
		return f
	}
	if f.actTime != actTime {
		f.err = errors.New("actTime err")
	}
	return f
}

func (f *Filter) CheckBenefitType(benefitType int64) *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckBenefitType")
	if f.benefitType != benefitType {
		f.err = errors.New("benefitType err")
	}
	return f
}
func (f *Filter) Err() error {
	return f.err
}

func client() {
	f := Filter{
		adiu:        "adiu_1234",
		poiid:       "BX10010",
		benefitType: 1,
		actTime:     time.Now(),
	}
	err := f.CheckAdiu("adiu_1234").
		CheckPoiid("BX10010").
		CheckActTime(time.Now()).
		CheckBenefitType(int64(1)).
		Err()
	fmt.Println("first check ", err)

	err = f.CheckActTime(time.Now()).CheckAdiu("adiu_1234").CheckPoiid("BX10010").CheckBenefitType(int64(2)).Err()
	fmt.Println("second check", err)

}
