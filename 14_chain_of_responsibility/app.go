package chain_of_responsibility

import (
	"errors"
	"fmt"
)

type Filter struct {
	err error
}

func (f *Filter) CheckAdiu(adiu string) *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckAdiu")
	return f
}

func (f *Filter) CheckPoiid(poiid string) *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckPoiid")
	return f
}

func (f *Filter) CheckActTime() *Filter {
	if f.err != nil {
		return f
	}
	f.err = errors.New("act err")
	return f
}

func (f *Filter) CheckBenefitType() *Filter {
	if f.err != nil {
		return f
	}
	fmt.Println("CheckBenefitType")
	return f
}
func (f *Filter) Err() error {
	return f.err
}

func Client() {
	f := Filter{}
	err := f.CheckAdiu("adiu").
		CheckPoiid("poiid").
		CheckActTime().
		CheckBenefitType().
		Err()
	fmt.Println(err)

	err = f.CheckActTime().CheckAdiu("adiu").CheckPoiid("poiid").CheckBenefitType().Err()
	fmt.Println(err)

}
