package ui


type LineChart struct {
	Type  		string
	Title  		string
	Scales 		int
	FillLine 	bool
	Labels 		[]string
	Lines 		[]ChartItemLine
}

func NewLineChart(title string) LineChart {
	return LineChart{Type: "chart_line", Scales: 1, FillLine: false, Title: title}
}

type ChartItemLine struct {
	Type  	string
	Color 	string
	Title 	string
	Values  []string
}

func NewChartItemLine(title string) ChartItemLine {
	return ChartItemLine{Type: "chart_item_line", Title: title}
}




type BarChart struct {
	Type  	string
	Title  	string
	Bars 	[]ChartItemBar
}

func NewBarChart(title string) BarChart {
	return BarChart{Type: "chart_bar", Title: title}
}

type ChartItemBar struct {
	Type  	string
	Color 	string
	Title 	string
	Value 	string
}

func NewChartItemBar(title, value string) ChartItemBar {
	return ChartItemBar{Type: "chart_item_bar", Title: title, Value: value}
}
