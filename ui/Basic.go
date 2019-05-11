package ui


type Array struct {
	Type  string
	Body  []interface{}
}

func NewArray() Array {
	return Array{Type: "array"}
}

type String struct {
	Type  string
	Title string
	Bold  bool
}

func NewString(title string) String {
	return String{Type: "string", Title: title}
}

type Command struct {
	Type  string
	Cmd   CommandType
}

func NewCommand(cmd CommandType) Command {
	return Command{Type: "command", Cmd: cmd}
}

type Alert struct {
	Type  string
	Title string
	Msg   string
}

func NewAlert(title string, msg string) Alert {
	return Alert{Type: "alert", Title: title, Msg: msg}
}

type Button struct {
	Type  string
	Title string
	Color ColorType
	Icon string
	Disabled bool
	Function Function
}

func NewButton(title string, function Function) Button {
	return Button{Type: "button", Title: title, Function: function, Disabled: false}
}
func NewButtonColor(title string, function Function, color ColorType) Button {
	return Button{Type: "button", Title: title, Function: function, Disabled: false, Color: color}
}

type Function struct {
	Type string
	Name FunctionNameType
	Data interface{}
}

func NewFunction(functionName FunctionNameType, data interface{}) Function {
	return Function{Type: "function", Name: functionName, Data: data}
}

type Box struct {
	Type  string
	Icon  string
	Size  string
	Body  []interface{}
}

func NewBox() Box {
	return Box{Type: "box", Size: "l"}
}

type Menu struct {
	Type     string
	MenuType string
	Icon     string
	Title    string
	Submenus []Menu
	Function Function
}

func NewMenu(title string, function Function) Menu {
	return Menu{Type: "menu", MenuType:"menu", Title: title, Function: function}
}

func NewIconMenu(icon string, title string, function Function) Menu {
	return Menu{Type: "menu", MenuType:"menu", Icon: icon, Title: title, Function: function}
}


type Dialog struct {
	Type           string
	DialogId       string
	Title      	   string
	HasCloseButton bool
	Buttons        []Button
	Body           []interface{}
}

func NewDialog(dialogId string, title string) Dialog {
	return Dialog{Type: "dialog", HasCloseButton: true, DialogId: dialogId, Title: title}
}


type EntryButtons struct {
	Type      string
	Title     string
	Buttons   []Button
}

func NewEntryButtons(title string) EntryButtons {
	return EntryButtons{Type: "entry_buttons", Title: title}
}

type Markdown struct {
	Type     string
	Body	 string
}

func NewMarkdown(body string) Markdown {
	return Markdown{Type: "markdown", Body: body}
}


type CardSwitch struct {
	Type      	string
	Title     	string
	Description string
	ButtonOn   	Button
	ButtonOff  	Button
}

func NewCardSwitch(title string, buttonOn, buttonOff Button) CardSwitch {
	return CardSwitch{Type: "card_switch", Title: title, ButtonOn: buttonOn, ButtonOff: buttonOff}
}

type CardButtons struct {
	Type      	string
	Title     	string
	Description string
	Buttons   	[]Button
}

func NewCardButtons(title string) CardButtons {
	return CardButtons{Type: "card_buttons", Title: title}
}

type CardInfo struct {
	Type      	string
	Title     	string
	Description string
	Button   	Button
}

func NewCardInfo(title string) CardInfo {
	return CardInfo{Type: "card_info", Title: title}
}