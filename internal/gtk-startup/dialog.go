package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/hultan/gtk-startup/pkg/tools"
	gtkHelper "github.com/hultan/softteam-tools/pkg/gtk-helper"
)

type Dialog struct {
	Dialog *gtk.Dialog
}

func NewDialog() *Dialog {
	return new(Dialog)
}

func (d *Dialog) OpenDialog(parent gtk.IWindow) {
	// Create a new gtk helper
	builder, err := gtk.BuilderNewFromFile(tools.GetResourcePath("../assets", "main.glade"))
	tools.ErrorCheckWithPanic(err, "Failed to create builder")
	helper := gtkHelper.GtkHelperNew(builder)

	// Get the dialog window from glade
	dialog, err := helper.GetDialog("settings_dialog")
	tools.ErrorCheckWithPanic(err,"Failed to open settings_dialog")

	dialog.SetTitle("settings dialog")
	dialog.SetTransientFor(parent)
	dialog.SetModal(true)

	// Save button
	saveButton, err := helper.GetButton("dialog_save_button")
	tools.ErrorCheckWithPanic(err, "Failed to find dialog_save_button")
	_, err = saveButton.Connect("clicked", dialog.Destroy)
	tools.ErrorCheckWithPanic(err, "Failed to connect the dialog_save_button.clicked event")

	// Cancel button
	cancelButton, err := helper.GetButton("dialog_cancel_button")
	tools.ErrorCheckWithPanic(err, "Failed to find dialog_cancel_button")
	_, err = cancelButton.Connect("clicked", dialog.Destroy)
	tools.ErrorCheckWithPanic(err, "Failed to connect the dialog_cancel_button.clicked event")

	dialog.Show()
}