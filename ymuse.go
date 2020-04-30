package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/op/go-logging"
	"github.com/yktoo/ymuse/internal/player"
	"github.com/yktoo/ymuse/internal/util"
	"os"
)

var log *logging.Logger

func main() {
	// Init logging
	logging.SetFormatter(logging.MustStringFormatter(`%{time:15:04:05.000} %{level:-5s} %{module} %{message}`))
	log = logging.MustGetLogger("main")
	logging.SetLevel(util.GetConfig().LogLevel, "main")

	// Start the app
	log.Info("Ymuse version", util.AppVersion)

	// Create Gtk Application, change appID to your application domain name reversed.
	application, err := gtk.ApplicationNew(util.AppID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application", err)
	}

	// Setup the application
	if _, err = application.Connect("activate", func() { onActivate(application) }); err != nil {
		log.Fatal("Failed to connect activation signal", err)
	}

	// Run the application
	os.Exit(application.Run(nil))
}

func onActivate(application *gtk.Application) {
	// Create the main window
	if window, err := player.NewMainWindow(application, ":6600"); err != nil {
		log.Fatal("Could not create application window", err)
	} else {
		window.Show()
	}
}
