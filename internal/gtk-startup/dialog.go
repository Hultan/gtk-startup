package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
)

func (m MainForm) openDialog() {
	if m.dialog == nil {
		// Get the dialog window from glade
		m.dialog = m.builder.getObject("settings_dialog").(*gtk.Dialog)

		m.dialog.SetTitle("settings dialog")
		m.dialog.SetTransientFor(m.window)
		m.dialog.SetModal(true)

		// Hook up the destroy event
		m.dialog.Connect(
			"destroy", func(w *gtk.Dialog) {
				m.dialog.Destroy()
				m.dialog = nil
			},
		)
	}

	responseId := m.dialog.Run()

	if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
		m.dialog.Hide()
	} else if responseId == gtk.RESPONSE_ACCEPT || responseId == gtk.RESPONSE_OK || responseId == gtk.RESPONSE_YES {
		// Save settings here
		m.dialog.Hide()
	}
}
