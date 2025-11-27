package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

var _ fyne.Theme = (*CustomTheme)(nil)

func NewCustomTheme() fyne.Theme {
	return &CustomTheme{}
}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 245, G: 245, B: 250, A: 255}
	case theme.ColorNameButton:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 255}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 150, G: 150, B: 150, A: 255}
	case theme.ColorNameError:
		return color.NRGBA{R: 239, G: 68, B: 68, A: 255}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 200}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 30, G: 30, B: 50, A: 255}
	case theme.ColorNameHover:
		return color.NRGBA{R: 120, G: 113, B: 248, A: 255}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 200, G: 200, B: 210, A: 255}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 150, G: 150, B: 160, A: 255}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 79, G: 70, B: 229, A: 255}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 255}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 200, G: 200, B: 210, A: 150}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 100}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 220, G: 220, B: 230, A: 255}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0, G: 0, B: 0, A: 50}
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 34, G: 197, B: 94, A: 255}
	case theme.ColorNameWarning:
		return color.NRGBA{R: 251, G: 191, B: 36, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameScrollBar:
		return 12
	case theme.SizeNameScrollBarSmall:
		return 6
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameInputBorder:
		return 2
	case theme.SizeNameSelectionRadius:
		return 4
	case theme.SizeNameInputRadius:
		return 6
	default:
		return theme.DefaultTheme().Size(name)
	}
}

