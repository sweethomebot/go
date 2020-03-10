package ui


import (
	"github.com/sweethomebot/go/core/v1"
)

type Form struct {
	Type        string
	FormId      string
	ApiEndpoint string
	Body        []interface{}
}

func NewForm(formId string, apiEndpoint string) Form {
	return Form{Type: "form", FormId: formId, ApiEndpoint: apiEndpoint}
}


// Used by HHC Dvc, Setting, Permission
type FormInput struct {
	Type  		string
	Title  		string
	Name  		string
	BspValue 	string
	Info  		string
	Options		[]FormOption
	Rule		string
	Disabled	bool
}

type FormOption struct {
	Title  	string
	Value	string
}



func NewFormInput(ch core.Channel, transModul string, input FormInput, value string) interface{} {
	switch input.Type {
	case "select":
		choice := NewInputSelect(core.TransPlugin(ch, transModul, input.Title), input.Name, value)
		for _, option := range input.Options  {
			choice.Options = append(choice.Options, NewInputOption(core.TransPlugin(ch, transModul, option.Title), option.Value))
		}
		if input.Disabled {
			choice.Disabled = true
		}
		choice.Info = core.TransPlugin(ch, transModul, input.Info)
		choice.Rule = input.Rule
		return choice
		// TODO : implement line
	case "bool":
		checkbox := NewInputCheckbox(core.TransPlugin(ch, transModul, input.Title), input.Name, value == "1")
		if input.Disabled {
			checkbox.Disabled = true
		}
		checkbox.Info = core.TransPlugin(ch, transModul, input.Info)
		checkbox.Rule = input.Rule
		return checkbox
	default:
		textInput := NewInput(core.TransPlugin(ch, transModul, input.Title), input.Name, value)
		textInput.InputType = input.Type
		textInput.Info = core.TransPlugin(ch, transModul, input.Info)
		textInput.BspValue = input.BspValue
		if input.Disabled {
			textInput.Disabled = true
		}
		textInput.Rule = input.Rule
		return textInput
	}
}

func NewFormInputs(ch core.Channel, transModul string, formData []FormInput, values map[string]string) []interface{} {
	var retData []interface{}
	for _, input := range formData  {
		value, ok := values[input.Name]
		if !ok {
			value = ""
		}
		retData = append(retData, NewFormInput(ch, transModul, input, value) )
	}
	return retData
}






type Input struct {
	Type      string
	InputType string
	Name      string
	Title     string
	Value     string
	BspValue  string
	Info	  string
	Rule	  string
	Disabled  bool
}

func NewInput(title string, name string, value string) Input {
	return Input{Type: "input", InputType: "text", Disabled: false, Title: title, Name: name, Value: value}
}
func NewHiddenInput( name string, value string) Input {
	return Input{Type: "input", InputType: "hidden", Name: name, Value: value}
}

type InputSelect struct {
	Type      string
	Name      string
	Title     string
	Value     string
	Info	  string
	Rule	  string
	Options   []InputOption
	Disabled  bool
}

type InputOption struct {
	Type      string
	Title     string
	Value	  string
	Disabled  bool
}

func NewInputSelect(title string, name string, value string) InputSelect {
	return InputSelect{Type: "input_select", Disabled: false, Title: title, Name: name, Value: value}
}
func NewInputOption(title string, value string) InputOption {
	return InputOption{Type: "input_option", Disabled: false, Title: title, Value: value}
}

type InputCheckbox struct {
	Type      string
	Name      string
	Title     string
	Info	  string
	Rule	  string
	Checked	  bool
	Disabled  bool
}

func NewInputCheckbox(title string, name string, checked bool) InputCheckbox {
	return InputCheckbox{Type: "input_checkbox", Disabled: false, Title: title, Name: name,  Checked: checked}
}