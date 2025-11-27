package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.NewWithID("com.notepad.app")
	// استفاده از تم سفارشی زیبا
	myApp.Settings().SetTheme(NewCustomTheme())

	window := myApp.NewWindow("Notepad App - دفترچه یادداشت")
	window.Resize(fyne.NewSize(900, 700))
	window.CenterOnScreen()
	window.SetFixedSize(false)

	// ایجاد تب‌ها با نام‌های انگلیسی برای جلوگیری از مشکل encoding
	tabs := container.NewAppTabs(
		&container.TabItem{
			Text:    "Notes",
			Content: NewNotesTab(),
		},
		&container.TabItem{
			Text:    "Todos",
			Content: NewTodoTab(),
		},
		&container.TabItem{
			Text:    "Calculator",
			Content: NewCalculatorTab(),
		},
	)

	tabs.SetTabLocation(container.TabLocationTop)

	window.SetContent(tabs)
	window.ShowAndRun()
}

