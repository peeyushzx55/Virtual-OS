package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGallery(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Gallery")
	w.Resize(fyne.NewSize(800, 600))
	root_src := "C:\\Users\\Om sh Ganeshya namah\\Pictures"
	files, err := ioutil.ReadDir(root_src)

	if err != nil {
        log.Fatal(err)
    }


	tabs := container.NewAppTabs()

	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg" || extension == "jpg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
			image.FillMode = canvas.ImageFillContain
			tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
    }

	tabs.SetTabLocation(container.TabLocationBottom)
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(tabs)
	w.Show()
}