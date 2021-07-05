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
	extraWindow := builder.getObject("extra_window").(*gtk.Window)

	// Set up the extra window
	extraWindow.SetTitle("extra form")
	extraWindow.HideOnDelete()
	extraWindow.SetModal(true)
	extraWindow.SetKeepAbove(true)
	extraWindow.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	// Hook up the destroy event
	_, err := extraWindow.Connect("destroy", extraWindow.Destroy)
	ErrorCheckWithPanic(err, "Failed to connect the extraWindow.destroy event")

	// Close button
	button := builder.getObject("extra_window_close_button").(*gtk.Button)
	_, err = button.Connect("clicked", extraWindow.Destroy)
	ErrorCheckWithPanic(err, "Failed to connect the extra_window_close_button.clicked event")

	e.Window = extraWindow

	// Show the window
	extraWindow.ShowAll()
}