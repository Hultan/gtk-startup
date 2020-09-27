package main

import (
	gtkStartup "github.com/hultan/gtk-startup/internal/gtk-startup"
	"github.com/hultan/gtk-startup/pkg/tools"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	ApplicationId    = "se.softteam.gtk-startup"
	ApplicationFlags = glib.APPLICATION_FLAGS_NONE
)

func main() {
	// Create a new application
	application, err := gtk.ApplicationNew(ApplicationId, ApplicationFlags)
	tools.ErrorCheckWithPanic(err, "Failed to create GTK Application")

	mainForm := gtkStartup.NewMainForm()
	// Hook up the activate event handler
	_, err = application.Connect("activate", mainForm.OpenMainForm)
	tools.ErrorCheckWithPanic(err, "Failed to connect Application.Activate event")

	// Start the application (and exit when it is done)
	os.Exit(application.Run(nil))
}
