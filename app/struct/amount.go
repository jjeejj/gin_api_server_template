package _struct

import "github.com/shopspring/decimal"

// 金额处理, 和浮点数直接相互转换

type Amount int64

func NewAAmount(amount int64) Amount {
	return Amount(amount)
}

// ToFloat64ByDivRate 除于给定精度返回对应的浮点数
func (a Amount) ToFloat64ByDivRate(rato int32) float64 {
	result, _ := decimal.NewFromUint64(uint64(a)).Div(decimal.NewFromInt32(rato)).Float64()
	return result
}
