package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
	gtkHelper "github.com/hultan/softteam/gtk"
	"gtk-startup/pkg/tools"
)

type ExtraForm struct {
	Window *gtk.Window
}

func NewExtraForm() *ExtraForm {
	return new(ExtraForm)
}

func (e *ExtraForm) OpenForm() {
	// Create a new gtk helper
	builder, err := gtk.BuilderNewFromFile(tools.GetResourcePath("../assets", "main.glade"))
	tools.ErrorCheckWithPanic(err, "Failed to create builder")
	helper := gtkHelper.GtkHelperNew(builder)

	// Get the extra window from glade
	extraForm, err := helper.GetWindow("extra_window")
	tools.ErrorCheckWithPanic(err,"Failed to open extra_window")

	// Set up the extra window
	extraForm.SetTitle("extra form")
	extraForm.HideOnDelete()
	extraForm.SetModal(true)
	extraForm.SetKeepAbove(true)
	extraForm.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	// Hook up the destroy event
	_, err = extraForm.Connect("destroy", extraForm.Destroy)
	tools.ErrorCheckWithPanic(err, "Failed to connect the extraForm.destroy event")

	// Close button
	button, err := helper.GetButton("extra_window_close_button")
	tools.ErrorCheckWithPanic(err, "Failed to find extra_window_close_button")
	_, err = button.Connect("clicked", extraForm.Destroy)
	tools.ErrorCheckWithPanic(err, "Failed to connect the extra_window_close_button.clicked event")

	// Show the window
	extraForm.ShowAll()
}