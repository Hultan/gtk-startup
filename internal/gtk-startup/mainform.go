package gtkStartup

import (
	"github.com/gotk3/gotk3/gtk"
	gtkHelper "github.com/hultan/softteam/gtk"
	"gtk-startup/pkg/tools"
	"os"
)

type MainForm struct {
	Window *gtk.ApplicationWindow
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
	statusBar.Push(statusBar.GetContextId("gtk-startup"),"gtk-startup : version 0.1.0")

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