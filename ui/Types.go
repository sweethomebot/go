package ui

type CommandType string
const (
	CommandReload		= CommandType("reload")
	CommandReloadSite	= CommandType("reloadSite")
	CommandReloadRoom	= CommandType("reloadRoom")
)

type ColorType string
const (
	ColorGreen			= ColorType("green")
	ColorGrey			= ColorType("grey")
	ColorWhite			= ColorType("white")
	ColorRed			= ColorType("red")
	ColorBlue			= ColorType("blue")
	ColorLightBlue		= ColorType("lightblue")
	ColorYellow			= ColorType("yellow")
	ColorOrange			= ColorType("orange")
)

type FunctionNameType string
const (
	FunctionDriverButton	= FunctionNameType("driverButton")
	FunctionLoad			= FunctionNameType("load")
	FunctionOpenDialog		= FunctionNameType("openDialog")
	FunctionLoadOpenDialog	= FunctionNameType("loadOpenDialog")
	FunctionSendForm		= FunctionNameType("sendForm")
	FunctionCallApi			= FunctionNameType("callApi")
	FunctionOpenStore		= FunctionNameType("openStore")
)
