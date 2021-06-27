package gocorona

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Credits represents the creates widget.
type Credits struct {
	Widget *widgets.Paragraph
}

// Construct creates the credits widget
// containing info about the project.
func (c *Credits) Construct() {
	widget := widgets.NewParagraph()

	widget.Title = "💪 About Gocorona"
	widget.Text = `
Worldwide COVID-19 tracker for your terminal.

©2020 Ayooluwa Isaiah and other contributors.

Gocorona is open source software made available under the terms of the MIT license.

Star the repo or contribute on GitHub: [https://github.com/ayoisaiah/gocorona](fg:blue)

Gocorona relies heavily on other open source software listed below:

👉 Termui: [https://github.com/gizak/termui](fg:blue)
👉 Disease.sh API: [https://github.com/disease-sh/API](fg:blue)
	`
	widget.TextStyle = ui.NewStyle(ui.ColorClear)
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)

	c.Widget = widget
}
