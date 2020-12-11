package main

import "fmt"

//["CNY", "USD", 0.14], //1 CNY = 0.14 USD
//["JPY", "CNY", 0.065],
//["CAD", "JPY", 81.25],
//["EUR", "HKD", 9.10]

var (
	relaMap = map[string]map[string]float64{
		"CNY": {"USD": 0.14},
		"JPY": {"CNY": 0.065},
		"CAD": {"JPY": 81.25},
		"EUR": {"HKD": 9.10},
	}
)

type ConvRate struct {
	from string
	to   string
	rate float64
}

func NewConvRate() *ConvRate {
	return &ConvRate{}
}

// 可在此处实现已知汇率的预处理
func (cr *ConvRate) CurrencyConv() {

}

// 在此处实现两个货币的换汇
func (cr *ConvRate) GetConvRate(from, to string) (res float64) {
	if res, ok := relaMap[from][to]; ok {
		return res
	}

	for k, v := range relaMap {
		for k1, v1 := range v {
			fmt.Println(k, k1, v1)
		}
	}

	return -1
}

func main() {
	from, to := "USD", "CAD" //"USD", "CAD" 1.37 "EUR", "JPY" -1
	rate := NewConvRate().GetConvRate(from, to)

	fmt.Println(rate)
}
