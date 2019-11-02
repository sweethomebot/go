package uiAutomate


type AutomateGrid struct {
	Type  		string
	Id  		string
	Title  		string
	Json  		string
}

func NewAutomateGrid(id, title, json string) AutomateGrid {
	return AutomateGrid{Type: "automate_grid", Id: id, Title: title, Json: json}
}
