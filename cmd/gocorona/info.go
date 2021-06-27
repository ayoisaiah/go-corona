package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// CoronavirusInfo represents the widget
// that provides info about the coronavirus pandemic.
type CoronavirusInfo struct {
	Widget *widgets.Paragraph
}

// Construct creates a widget containing
// info about the coronavirus pandamic and how
// to protect oneself.
func (ci *CoronavirusInfo) Construct() {
	widget := widgets.NewParagraph()

	widget.Title = "🤒 Learn about the Coronavirus Pandemic"
	widget.Text = `
COVID-19 is a new illness that can affect your lungs and airways. It's caused by a virus called coronavirus.

[ How it spreads ](fg:white,bg:yellow,fg:bold)

The virus is thought to spread mainly from person-to-person contact through respiratory droplets produced when an infected person coughs or sneezes.

[ Symptoms ](fg:white,bg:yellow,fg:bold)

The main symptoms of coronavirus are:
👉 a high temperature – this means you feel hot to touch on your chest or back (you do not need to measure your temperature).
👉 a new, continuous cough – this means coughing a lot for more than an hour, or 3 or more coughing episodes in 24 hours (if you usually have a cough, it may be worse than usual).

You can usually treat mild coronavirus (COVID-19) symptoms at home. If your symptoms are severe, you may need medical care until you recover.

[ Prevention ](fg:white,bg:yellow,fg:bold)

To help stop the spread of coronavirus (COVID-19), avoid close contact with anyone you do not live with and wash your hands regularly.

To stop the spread of coronavirus, you should:
👉 wear a face mask when you are outdoors.
👉 wash your hands with soap and water often – for at least 20 seconds.
👉 wash your hands as soon as you get home.
👉 cover your mouth and nose with a tissue when you cough or sneeze.
👉 put used tissues in the bin immediately and wash your hands.
👉 not touch your face if your hands are not clean.

Learn more at [https://www.nhs.uk/conditions/coronavirus-covid-19/](fg:blue)
	`
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)
	widget.TextStyle = ui.NewStyle(ui.ColorClear)

	ci.Widget = widget
}
