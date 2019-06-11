package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"time"
)

func Draw(s string) error {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}
	win.SetDecorated(false)
	win.SetTitle("go-lingua")
	_, err = win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		return err
	}
	l, err := gtk.LabelNew("")
	if err != nil {
		return err
	}
	l.SetMarkup(fmt.Sprintf("<span size='10000'>%s</span>", s))
	win.Add(l)
	win.SetDefaultSize(150, 200)
	win.Move(0, 0)
	go func() {
		time.Sleep(5 * time.Second)
		win.Close()
	}()
	win.ShowAll()
	gtk.Main()

	return nil
}
