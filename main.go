package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindown    fyne.Window
	PriceContainer *fyne.Container
}

var myApp Config

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database

	// create a database repository

	// create and size a fyne window
	myApp.MainWindown = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindown.Resize(fyne.NewSize(770, 410))
	myApp.MainWindown.SetFixedSize(true)
	myApp.MainWindown.SetMaster()

	// show and run the application
	myApp.makeUI()
	myApp.MainWindown.ShowAndRun()
}
