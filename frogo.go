package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gui() {
	app := app.New()
	window := app.NewWindow("FRoGo (Zum Beenden: Strg + Shift + q)")
	label := widget.NewLabel("Minuten:")
	entry := widget.NewEntry()
	button := widget.NewButton("Start", func() {
		cli(strconv.Atoi(entry.Text))
	})
	grid := container.New(layout.NewGridLayout(3), label, entry, button)
	window.SetIcon(theme.FyneLogo())
	window.SetContent(grid)
	window.Resize(fyne.NewSize(500, 0))
	window.SetFixedSize(true)
	window.CenterOnScreen()
	go func() {
		hotkey := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyQ)
		hotkey.Register()
		for range hotkey.Keydown() {
			os.Exit(0)
		}
	}()
	window.ShowAndRun()
}

func cli(in int, er error) {
	timeout := time.After(time.Duration(in) * time.Minute)
	if in <= 0 || er != nil {
		timeout = time.After(24 * time.Hour)
	}
	for {
		select {
			case <-timeout:
				os.Exit(0)
			default:
				robotgo.MouseSleep = 100
				width, height := robotgo.GetScreenSize()
				rand.Seed(time.Now().UnixNano())
				robotgo.MoveSmooth(rand.Intn(width), rand.Intn(height), 1.1/3, 1.1)
				robotgo.Click("right")
				time.Sleep(2 * time.Second)
				robotgo.KeyTap("esc")
		}
	}
}

func main() {
	gui()
}
