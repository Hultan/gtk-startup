package gtkStartup

import (
	_ "embed"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

const applicationTitle = "gtk-startup"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	window      *gtk.ApplicationWindow
	builder     *gtkBuilder
	aboutDialog *gtk.AboutDialog
	extraForm   *gtk.Window
	dialog      *gtk.Dialog
}

//go:embed assets/main.glade
var gladeFile string

//go:embed assets/application.png
var applicationIcon []byte

// NewMainForm : Creates a new MainForm object
func NewMainForm() MainForm {
	return MainForm{}
}

// OpenMainForm : Opens the MainForm window
func (m MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	builder, err := newBuilder(gladeFile)
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.window = m.builder.getObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.window.SetApplication(app)
	m.window.SetTitle("gtk-startup main window")

	// Hook up the destroy event
	m.window.Connect("destroy", m.window.Destroy)

	// Quit button
	button := m.builder.getObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.window.Destroy)

	// Status bar
	statusBar := m.builder.getObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("gtk-startup"), "gtk-startup : version 0.1.0")

	// Open form button
	openFormButton := m.builder.getObject("main_window_open_form_button").(*gtk.Button)
	openFormButton.Connect(
		"clicked", func() {
			m.openForm()
		},
	)

	// Open dialog button
	openDialogButton := m.builder.getObject("main_window_open_dialog_button").(*gtk.Button)
	openDialogButton.Connect(
		"clicked", func() {
			m.openDialog()
		},
	)

	// Menu
	m.setupMenu()

	// Show the main window
	m.window.ShowAll()
}

func (m MainForm) setupMenu() {
	menuQuit := m.builder.getObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.window.Destroy)

	menuHelpAbout := m.builder.getObject("menu_help_about").(*gtk.MenuItem)
	menuHelpAbout.Connect(
		"activate", func() {
			m.openAboutDialog()
		},
	)
}
