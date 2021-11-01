package main

import (
	"io/ioutil"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor() {
	w := myApp.NewWindow("Text Editor")
	// a := app.New()
	// w := a.NewWindow("Text Editor")

	// w.Resize(fyne.NewSize(600, 600))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("My Text Editor"),
		),
	)

	var count int = 1

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save Text File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)
		
		saveFileDialog.SetFileName("New File" + strconv.Itoa(count - 1) + ".txt")

		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)
				output := fyne.NewStaticResource("New File", readData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))

					w.SetContent(container.NewScroll(viewData))

					w.Resize(fyne.NewSize(400, 400))

					w.Show()
			}, w)

			openFileDialog.SetFilter(
				storage.NewExtensionFileFilter([] string{".txt"}))
			openFileDialog.Show()

	})

	textContainer := container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
	)

	w.Resize(fyne.NewSize(600, 600))

	w.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, textContainer),
	)
	w.Show()
}