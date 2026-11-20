package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Paleta Refletunes — escura, com acento verde-água.
var (
	colBg        = color.NRGBA{0x0D, 0x11, 0x1A, 0xFF}
	colCard      = color.NRGBA{0x16, 0x1D, 0x29, 0xFF}
	colCardHover = color.NRGBA{0x1C, 0x26, 0x35, 0xFF}
	colAccent    = color.NRGBA{0x2D, 0xD4, 0xBF, 0xFF}
	colAccentDim = color.NRGBA{0x1B, 0x7A, 0x6E, 0xFF}
	colText      = color.NRGBA{0xE6, 0xEB, 0xF2, 0xFF}
	colTextDim   = color.NRGBA{0x8B, 0x96, 0xA5, 0xFF}
	colOK        = color.NRGBA{0x34, 0xD1, 0x7B, 0xFF}
	colWarn      = color.NRGBA{0xF5, 0xB1, 0x4C, 0xFF}
	colErr       = color.NRGBA{0xF0, 0x6A, 0x6A, 0xFF}
	colHeader    = color.NRGBA{0x11, 0x1A, 0x26, 0xFF}
)

type refletunesTheme struct{ fyne.Theme }

func newTheme() fyne.Theme { return &refletunesTheme{Theme: theme.DarkTheme()} }

func (t *refletunesTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return colBg
	case theme.ColorNameButton:
		return colCardHover
	case theme.ColorNamePrimary:
		return colAccent
	case theme.ColorNameFocus:
		return colAccent
	case theme.ColorNameSelection:
		return colAccentDim
	case theme.ColorNameInputBackground:
		return colCard
	case theme.ColorNameForeground:
		return colText
	case theme.ColorNamePlaceHolder:
		return colTextDim
	case theme.ColorNameHover:
		return colCardHover
	case theme.ColorNameMenuBackground:
		return colCard
	case theme.ColorNameOverlayBackground:
		return colCard
	case theme.ColorNameSuccess:
		return colOK
	case theme.ColorNameWarning:
		return colWarn
	case theme.ColorNameError:
		return colErr
	}
	return t.Theme.Color(name, variant)
}
