package gtkStartup

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/hultan/gtk-startup/pkg/tools"
	gtkHelper "github.com/hultan/softteam-tools/pkg/gtk-helper"
	"github.com/hultan/softteam-tools/pkg/resources"
	"os"
)

const applicationTitle = "gtk-startup"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	Helper      *gtkHelper.GtkHelper
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

	// Create a new gtk helper
	builder, err := gtk.BuilderNewFromFile(tools.GetResourcePath("../assets", "main.glade"))
	tools.ErrorCheckWithPanic(err, "Failed to create builder")
	helper := gtkHelper.GtkHelperNew(builder)
	m.Helper = helper

	// Get the main window from the glade file
	window, err := helper.GetApplicationWindow("main_window")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window")

	m.Window = window

	// Set up main window
	window.SetApplication(app)
	window.SetTitle("gtk-startup main window")

	// Hook up the destroy event
	_, err = window.Connect("destroy", window.Close)
	tools.ErrorCheckWithPanic(err, "Failed to connect the mainForm.destroy event")

	// Quit button
	button, err := helper.GetToolButton("main_window_quit_button")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window_quit_button")
	_, err = button.Connect("clicked", window.Close)
	tools.ErrorCheckWithPanic(err, "Failed to connect the main_window_quit_button.clicked event")

	// Status bar
	statusBar, err := helper.GetStatusBar("main_window_status_bar")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window_status_bar")
	statusBar.Push(statusBar.GetContextId("gtk-startup"), "gtk-startup : version 0.1.0")

	// Open form button
	openFormButton, err := helper.GetButton("main_window_open_form_button")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window_open_form_button")
	_, err = openFormButton.Connect("clicked", m.OpenForm)
	tools.ErrorCheckWithPanic(err, "Failed to connect the main_window_open_form_button.clicked event")

	// Open dialog button
	openDialogButton, err := helper.GetButton("main_window_open_dialog_button")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window_open_dialog_button")
	_, err = openDialogButton.Connect("clicked", m.OpenDialog)
	tools.ErrorCheckWithPanic(err, "Failed to connect the main_window_open_dialog_button.clicked event")

	// Menu
	m.setupMenu(window)

	// Show the main window
	window.ShowAll()
}

func (m *MainForm) OpenForm() {
	extraForm := NewExtraForm()
	extraForm.OpenForm()
}

func (m *MainForm) OpenDialog() {
	dialog := NewDialog()
	dialog.OpenDialog(m.Window)
}

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

func (m *MainForm) setupMenu(window *gtk.ApplicationWindow) {
	menuQuit, err := m.Helper.GetMenuItem("menu_file_quit")
	tools.ErrorCheckWithPanic(err, "failed to find menu item menu_file_quit")
	_, err = menuQuit.Connect("activate", window.Close)
	tools.ErrorCheckWithoutPanic(err,"failed to connect menu_file_quit.activate signal")

	menuHelpAbout, err := m.Helper.GetMenuItem("menu_help_about")
	tools.ErrorCheckWithPanic(err, "failed to find menu item menu_help_about")
	_, err = menuHelpAbout.Connect("activate", m.openAboutDialog)
	tools.ErrorCheckWithoutPanic(err,"failed to connect menu_help_about.activate signal")

}