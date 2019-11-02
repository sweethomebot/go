package automate


type BlockType string

const (
	BlockTypeStart		= BlockType("Start")
	BlockTypeOk			= BlockType("Ok")
	BlockTypeYesNo		= BlockType("YesNo")
)

type BlockNextConn string

const (
	BlockNextConnOk		= BlockNextConn("AUT_BNC_Ok")
	BlockNextConnYes	= BlockNextConn("AUT_BNC_Yes")
	BlockNextConnNo		= BlockNextConn("AUT_BNC_No")
)



type BlockConnType struct {
	Ok 		string
	Yes	 	string
	No		string
}

type BlockPosType struct {
	X 		int
	Y	 	int
}

type Block struct {
	Type		BlockType
	Plugin 		string
	ProgramId	string
	Id	 		string
	Icon   		string
	Url			string
	Title  		string
	Output		map[string]string
	Config		map[string]string
	Pos			BlockPosType
	Conn		BlockConnType
}

type BlockExecAnswer struct {
	Next		BlockNextConn
	Output 		map[string]string
}



