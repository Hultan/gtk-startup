package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
	"os"
)

const applicationTitle = "gtk-startup"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder      *SoftBuilder
	AboutDialog *gtk.AboutDialog
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	m.builder = newSoftBuilder("main.glade")

	// Get the main window from the glade file
	m.Window = m.builder.getObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("gtk-startup main window")

	// Hook up the destroy event
	_, err := m.Window.Connect("destroy", m.Window.Close)
	ErrorCheckWithPanic(err, "Failed to connect the mainForm.destroy event")

	// Quit button
	button := m.builder.getObject("main_window_quit_button").(*gtk.ToolButton)
	_, err = button.Connect("clicked", m.Window.Close)
	ErrorCheckWithPanic(err, "Failed to connect the main_window_quit_button.clicked event")

	// Status bar
	statusBar := m.builder.getObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("gtk-startup"), "gtk-startup : version 0.1.0")

	// Open form button
	openFormButton := m.builder.getObject("main_window_open_form_button").(*gtk.Button)
	_, err = openFormButton.Connect("clicked", m.OpenForm)
	ErrorCheckWithPanic(err, "Failed to connect the main_window_open_form_button.clicked event")

	// Open dialog button
	openDialogButton := m.builder.getObject("main_window_open_dialog_button").(*gtk.Button)
	_, err = openDialogButton.Connect("clicked", m.OpenDialog)
	ErrorCheckWithPanic(err, "Failed to connect the main_window_open_dialog_button.clicked event")

	// Menu
	m.setupMenu(m.Window)

	// Show the main window
	m.Window.ShowAll()
}

func (m *MainForm) OpenForm() {
	extraForm := NewExtraForm()
	extraForm.OpenForm()
}

func (m *MainForm) OpenDialog() {
	dialog := NewDialog()
	dialog.OpenDialog(m.Window)
}
//
//func (m *MainForm) openAboutDialog() {
//	if m.AboutDialog == nil {
//		about, err := m.Helper.GetAboutDialog("about_dialog")
//		tools.ErrorCheckWithPanic(err, "failed to find dialog about_dialog")
//		about.SetDestroyWithParent(true)
//		about.SetTransientFor(m.Window)
//		about.SetProgramName(applicationTitle)
//		about.SetComments("An application...")
//		about.SetVersion(applicationVersion)
//		about.SetCopyright(applicationCopyRight)
//		resource := resources.NewResources()
//		image, err := gdk.PixbufNewFromFile(resource.GetResourcePath("application.png"))
//		if err == nil {
//			about.SetLogo(image)
//		}
//		about.SetModal(true)
//		about.SetPosition(gtk.WIN_POS_CENTER)
//
//		_, err = about.Connect("response", func(dialog *gtk.AboutDialog, responseId gtk.ResponseType) {
//			if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
//				about.Hide()
//			}
//		})
//		tools.ErrorCheckWithoutPanic(err,"failed to connect about_dialog.response signal")
//
//		m.AboutDialog = about
//	}
//
//	m.AboutDialog.Present()
//}

func (m *MainForm) setupMenu(window *gtk.ApplicationWindow) {
	menuQuit := m.builder.getObject("menu_file_quit").(*gtk.MenuItem)
	_, err := menuQuit.Connect("activate", window.Close)
	ErrorCheckWithoutPanic(err,"failed to connect menu_file_quit.activate signal")

	menuHelpAbout := m.builder.getObject("menu_help_about").(*gtk.MenuItem)
	_, err = menuHelpAbout.Connect("activate", m.openAboutDialog)
	ErrorCheckWithoutPanic(err,"failed to connect menu_help_about.activate signal")
}