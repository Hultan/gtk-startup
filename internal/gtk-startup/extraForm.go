package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
)

type ExtraForm struct {
	Window *gtk.Window
}

func NewExtraForm() *ExtraForm {
	return new(ExtraForm)
}

func (e *ExtraForm) OpenForm() {
	// Create a new gtk helper
	builder := newSoftBuilder("main.glade")

	// Get the extra window from glade
	extraForm := builder.getObject("extra_window").(*gtk.Window)

	// Set up the extra window
	extraForm.SetTitle("extra form")
	extraForm.HideOnDelete()
	extraForm.SetModal(true)
	extraForm.SetKeepAbove(true)
	extraForm.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	// Hook up the destroy event
	_, err := extraForm.Connect("destroy", extraForm.Destroy)
	ErrorCheckWithPanic(err, "Failed to connect the extraForm.destroy event")

	// Close button
	button := builder.getObject("extra_window_close_button").(*gtk.Button)
	_, err = button.Connect("clicked", extraForm.Destroy)
	ErrorCheckWithPanic(err, "Failed to connect the extra_window_close_button.clicked event")

	// Show the window
	extraForm.ShowAll()
}