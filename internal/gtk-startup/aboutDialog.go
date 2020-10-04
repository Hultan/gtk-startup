package gtkStartup

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/hultan/gtk-startup/pkg/tools"
	"github.com/hultan/softteam-tools/pkg/resources"
)

func (m *MainForm) openAboutDialog() {
	if m.AboutDialog == nil {
		about, err := m.Helper.GetAboutDialog("about_dialog")
		tools.ErrorCheckWithPanic(err, "failed to find dialog about_dialog")
		about.SetDestroyWithParent(true)
		about.SetTransientFor(m.Window)
		about.SetProgramName(applicationTitle)
		about.SetComments("An application...")
		about.SetVersion(applicationVersion)
		about.SetCopyright(applicationCopyRight)
		resource := resources.NewResources()
		image, err := gdk.PixbufNewFromFile(resource.GetResourcePath("application.png"))
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
		tools.ErrorCheckWithoutPanic(err,"failed to connect about_dialog.response signal")

		m.AboutDialog = about
	}

	m.AboutDialog.Present()
}
