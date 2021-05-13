package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
)

type Dialog struct {
	Dialog *gtk.Dialog
}

func NewDialog() *Dialog {
	return new(Dialog)
}

func (d *Dialog) OpenDialog(parent gtk.IWindow) {
	// Create a new softBuilder
	builder := newSoftBuilder("main.glade")

	// Get the dialog window from glade
	dialog := builder.getObject("settings_dialog").(*gtk.Dialog)

	dialog.SetTitle("settings dialog")
	dialog.SetTransientFor(parent)
	dialog.SetModal(true)

	_, err := dialog.Connect("response", func(dialog *gtk.Dialog, responseId gtk.ResponseType) {
		if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
			dialog.Hide()
		} else if responseId == gtk.RESPONSE_ACCEPT || responseId == gtk.RESPONSE_OK || responseId == gtk.RESPONSE_YES {
			// Save
			dialog.Hide()
		}
	})
	ErrorCheckWithoutPanic(err,"failed to connect about_dialog.response signal")

	dialog.Present()
}