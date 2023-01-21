package bento

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
	"log"
)

type Box struct {
	widget.BaseWidget

	items []*Item

	bottomOffset float32
}

func (b *Box) UpdateBottomOffset(offset float32) {
	b.bottomOffset = offset
	b.Refresh()
}

func (b *Box) CreateRenderer() fyne.WidgetRenderer {
	return &boxRenderer{box: b}
}

func NewBox() *Box {
	bx := &Box{}
	bx.ExtendBaseWidget(bx)

	return bx
}

func (b *Box) AddItem(i *Item) {
	b.items = append(b.items, i)
	i.closeAction = func() {
		idx := slices.Index(b.items, i)
		log.Println("Removing item:", idx)
		b.items = slices.Delete(b.items, idx, idx+1)
		b.Refresh()
	}
	b.Refresh()
}

type boxRenderer struct {
	box *Box
}

func (b *boxRenderer) Destroy() {

}

func (b *boxRenderer) Layout(size fyne.Size) {
	maxWidth := size.Width * 0.6667
	runningHeight := size.Height - theme.Padding() - b.box.bottomOffset
	for _, i := range b.box.items {
		iSize := i.MinSize()
		i.Resize(fyne.NewSize(maxWidth, iSize.Height))
		i.Move(fyne.NewPos(size.Width/2-maxWidth/2, runningHeight-iSize.Height))
		runningHeight -= theme.Padding() + iSize.Height
	}
}

func (b *boxRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (b *boxRenderer) Objects() []fyne.CanvasObject {
	var cos []fyne.CanvasObject
	for _, i := range b.box.items {
		cos = append(cos, i)
	}

	return cos
}

func (b *boxRenderer) Refresh() {

}
