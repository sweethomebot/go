package hhcType

const (
	WpFuncNACK = iota
	WpFuncACK
	WpFuncGetFuncList
	WpFuncSetGPIO
	WpFuncGetGPIO
	WpFuncSendGPIO
)

var WpFuncString = []string{
	"WpFuncNACK",
	"WpFuncACK",
	"WpFuncGetFuncList",
	"WpFuncSetGPIO",
	"WpFuncGetGPIO",
	"WpFuncSendGPIO",
}