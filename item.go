package bento

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
)

type ItemLength int

const (
	LengthShort ItemLength = iota
	LengthLong
	LengthIndefinite
)

var _ fyne.CanvasObject = (*Item)(nil)

type Item struct {
	widget.BaseWidget

	message      string
	actionTitle  string
	actionAction func()
	length       ItemLength

	closeAction func()

	backgroundColor color.Color
}

func (i *Item) CreateRenderer() fyne.WidgetRenderer {
	messageLbl := widget.NewLabel(i.message)
	messageLbl.Wrapping = fyne.TextWrapWord
	background := canvas.NewRectangle(i.backgroundColor)
	actionButton := widget.NewButton("Test", i.closeTapped)
	//actionButton.Importance = widget.LowImportance
	actionButton.Hide()
	closeButton := widget.NewButton("X", i.closeTapped)
	closeButton.Importance = widget.LowImportance

	return &itemRenderer{
		item:         i,
		messageLbl:   messageLbl,
		background:   background,
		actionButton: actionButton,
		closeButton:  closeButton,
	}
}

func (i *Item) closeTapped() {
	i.closeAction()
}

func (i *Item) SetBackgroundColor(c color.Color) {
	i.backgroundColor = c
	i.Refresh()
}

func (i *Item) AddAction(title string, action func()) {
	i.actionTitle = title
	i.actionAction = action
	i.Refresh()
}

func newDefaultItem(length ItemLength) *Item {
	background := colornames.Grey
	itm := &Item{
		length:          length,
		backgroundColor: background,
	}

	itm.ExtendBaseWidget(itm)

	return itm
}

func NewItemWithMessage(message string, length ItemLength) *Item {
	itm := newDefaultItem(length)
	itm.message = message

	return itm
}

type itemRenderer struct {
	item *Item

	background   *canvas.Rectangle
	messageLbl   *widget.Label
	actionButton *widget.Button
	closeButton  *widget.Button
}

func (i *itemRenderer) Destroy() {

}

func (i *itemRenderer) Layout(size fyne.Size) {
	messageSize := i.messageLbl.MinSize()
	actionSize := fyne.NewSize(0, 0)
	if i.actionButton.Visible() {
		actionSize = i.actionButton.MinSize()
	}
	closeSize := i.closeButton.MinSize()

	messagePos := fyne.NewPos(theme.Padding(), theme.Padding())
	i.messageLbl.Move(messagePos)
	i.messageLbl.Resize(fyne.NewSize(size.Width-actionSize.Width-closeSize.Width-4*theme.Padding(), messageSize.Height))

	closePos := fyne.NewPos(size.Width-theme.Padding()-closeSize.Width, theme.Padding())
	i.closeButton.Move(closePos)
	i.closeButton.Resize(closeSize)

	if i.actionButton.Visible() {
		actionPos := closePos.SubtractXY(theme.Padding()+actionSize.Width, 0)
		i.actionButton.Move(actionPos)
		i.actionButton.Resize(actionSize)
	}

	i.background.Resize(size)
	i.background.Move(fyne.NewPos(0, 0))

}

func (i *itemRenderer) MinSize() fyne.Size {
	messageSize := i.messageLbl.MinSize()
	actionSize := fyne.NewSize(0, 0)
	if i.actionButton.Visible() {
		actionSize = i.actionButton.MinSize()
	}
	closeSize := i.closeButton.MinSize()

	return fyne.NewSize(messageSize.Width+actionSize.Width+closeSize.Width+4*theme.Padding(), fyne.Max(messageSize.Height, closeSize.Height)+2*theme.Padding())
}

func (i *itemRenderer) Objects() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{i.background, i.messageLbl, i.closeButton}
	if i.actionButton != nil {
		objs = append(objs, i.actionButton)
	}
	return objs
}

func (i *itemRenderer) Refresh() {
	i.messageLbl.Text = i.item.message
	i.messageLbl.Refresh()
	i.background.FillColor = i.item.backgroundColor
	i.background.Refresh()
	if i.item.actionTitle != "" {
		i.actionButton.Show()
		i.actionButton.SetText(i.item.actionTitle)
		i.actionButton.OnTapped = i.item.actionAction
		i.actionButton.Refresh()
	}
}
