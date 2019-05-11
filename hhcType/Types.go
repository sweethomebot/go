package hhcType

type Driver struct {
	Plugin string
	Icon   string
	Url    string
	Title  string
	Form   string
}


type DeviceCreateRequest struct {
	DvcId		string
	FormData	map[string]string
}

type Device struct {
	DvcId	string
	HostId	string
	DrvUrl	string
	Title	string
	Type	DeviceType
	Config	string
	State	string
}

type DeviceType string

const (
	DeviceTypeText		= DeviceType("Text")
	DeviceTypeButtons	= DeviceType("Buttons")
	DeviceTypeOnOff		= DeviceType("OnOff")
	DeviceTypeDim		= DeviceType("Dim")
	DeviceTypeColor		= DeviceType("Color")
	DeviceTypeData		= DeviceType("Data")
)