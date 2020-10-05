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

	_, err = dialog.Connect("response", func(dialog *gtk.Dialog, responseId gtk.ResponseType) {
		if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
			dialog.Hide()
		} else if responseId == gtk.RESPONSE_ACCEPT || responseId == gtk.RESPONSE_OK || responseId == gtk.RESPONSE_YES {
			// Save
			dialog.Hide()
		}
	})
	tools.ErrorCheckWithoutPanic(err,"failed to connect about_dialog.response signal")

	dialog.Present()
}