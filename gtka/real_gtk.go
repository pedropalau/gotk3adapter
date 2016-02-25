package gtka

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/gotk3adapter/gdki"
	"github.com/twstrike/gotk3adapter/glibi"
	"github.com/twstrike/gotk3adapter/gtki"
)

type RealGtk struct{}

var Real = &RealGtk{}

func (*RealGtk) AboutDialogNew() (gtki.AboutDialog, error) {
	return wrapAboutDialog(gtk.AboutDialogNew())
}

func (*RealGtk) AccelGroupNew() (gtki.AccelGroup, error) {
	return wrapAccelGroup(gtk.AccelGroupNew())
}

func (*RealGtk) AcceleratorParse(acc string) (uint, gdki.ModifierType) {
	res, res2 := gtk.AcceleratorParse(acc)
	return res, gdki.ModifierType(res2)
}

func (*RealGtk) AddProviderForScreen(s gdki.Screen, provider gtki.StyleProvider, prio uint) {
	gtk.AddProviderForScreen(s, provider, prio)
}

func (*RealGtk) ApplicationNew(appId string, flags glibi.ApplicationFlags) (gtki.Application, error) {
	return wrapApplication(gtk.ApplicationNew(appId, flags))
}

func (*RealGtk) BuilderNew() (gtki.Builder, error) {
	return wrapBuilder(gtk.BuilderNew())
}

func (*RealGtk) CellRendererTextNew() (gtki.CellRendererText, error) {
	return wrapCellRendererText(gtk.CellRendererTextNew())
}

func (*RealGtk) CheckButtonNewWithMnemonic(label string) (gtki.CheckButton, error) {
	return wrapCheckButton(gtk.CheckButtonNewWithMnemonic(label))
}

func (*RealGtk) CssProviderNew() (gtki.CssProvider, error) {
	return wrapCssProvider(gtk.CssProviderNew())
}

func (*RealGtk) EntryNew() (gtki.Entry, error) {
	return wrapEntry(gtk.EntryNew())
}

func (*RealGtk) FileChooserDialogNewWith2Buttons(title string, parent gtki.Window, action gtki.FileChooserAction, first_button_text string, first_button_id gtki.ResponseType, second_button_text string, second_button_id gtki.ResponseType) (gtki.FileChooserDialog, error) {
	return wrapFileChooserDialog(gtk.FileChooserDialogNewWith2Buttons(title, parent, action, first_button_text, first_button_id, second_button_text, second_button_id))
}

func (*RealGtk) Init(args *[]string) {
	gtk.Init(args)
}

func (*RealGtk) LabelNew(str string) (gtki.Label, error) {
	return wrapLabel(gtk.LabelNew(str))
}

func (*RealGtk) ListStoreNew(types ...glibi.Type) (gtki.ListStore, error) {
	return wrapListStore(gtk.ListStoreNew(types...))
}

func (*RealGtk) MenuItemNew() (gtki.MenuItem, error) {
	return wrapMenuItem(gtk.MenuItemNew())
}

func (*RealGtk) MenuItemNewWithMnemonic(label string) (gtki.MenuItem, error) {
	return wrapMenuItem(gtk.MenuItemNewWithMnemonic(label))
}

func (*RealGtk) MenuNew() (gtki.Menu, error) {
	return wrapMenu(gtk.MenuNew())
}

func (*RealGtk) SeparatorMenuItemNew() (gtki.SeparatorMenuItem, error) {
	return wrapSeparatorMenuItem(gtk.SeparatorMenuItemNew())
}

func (*RealGtk) TextBufferNew(table gtki.TextTagTable) (gtki.TextBuffer, error) {
	return wrapTextBuffer(gtk.TextBufferNew(table))
}

func (*RealGtk) TextTagNew(name string) (gtki.TextTag, error) {
	return wrapTextTag(gtk.TextTagNew(name))
}

func (*RealGtk) TextTagTableNew() (gtki.TextTagTable, error) {
	return wrapTextTagTable(gtk.TextTagTableNew())
}

func (*RealGtk) TextViewNew() (gtki.TextView, error) {
	return wrapTextView(gtk.TextViewNew())
}

func (*RealGtk) WindowSetDefaultIcon(icon gdki.Pixbuf) {
	gtk.WindowSetDefaultIcon(icon)
}