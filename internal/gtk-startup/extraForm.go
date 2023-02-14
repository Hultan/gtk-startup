package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
)

func (m *MainForm) openForm() {
	if m.extraForm == nil {
		// Get the extra window from the glade file
		m.extraForm = m.builder.getObject("extra_window").(*gtk.Window)

		// Set up the extra window
		m.extraForm.SetTitle("extra form")
		m.extraForm.HideOnDelete()
		m.extraForm.SetModal(true)
		m.extraForm.SetTransientFor(m.window)
		m.extraForm.SetPosition(gtk.WIN_POS_CENTER_ON_PARENT)

		// Hook up the destroy event
		m.extraForm.Connect(
			"destroy", func() {
				m.extraForm.Destroy()
				m.extraForm = nil
			},
		)

		// Close button
		button := m.builder.getObject("extra_window_close_button").(*gtk.Button)
		button.Connect("clicked", m.extraForm.Hide)
	}

	// Show the window
	m.extraForm.Present()
}
