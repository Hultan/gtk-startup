package gtkStartup

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (m *MainForm) openAboutDialog() {
	if m.AboutDialog == nil {
		about := m.builder.getObject("about_dialog").(*gtk.AboutDialog)
		about.SetDestroyWithParent(true)
		about.SetTransientFor(m.Window)
		about.SetProgramName(applicationTitle)
		about.SetComments("An application...")
		about.SetVersion(applicationVersion)
		about.SetCopyright(applicationCopyRight)
		path, err := getResourcePath("application.png")
		ErrorCheckWithPanic(err, "failed to get application icon")
		image, err := gdk.PixbufNewFromFile(path)
		if err == nil {
			about.SetLogo(image)
		}
		about.SetModal(true)
		about.SetPosition(gtk.WIN_POS_CENTER)

		_, err = about.Connect("response", func(dialog *gtk.AboutDialog, responseId gtk.ResponseType) {
			if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
				about.Hide()
			}
		})
		ErrorCheckWithoutPanic(err,"failed to connect about_dialog.response signal")

		m.AboutDialog = about
	}

	m.AboutDialog.Present()
}
