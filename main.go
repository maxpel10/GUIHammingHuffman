package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	SizeW = 600
	SizeH = 400
)

var family = "Segoe UI"

func main() {
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Principal",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Hamming",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							preHammingWindow(mw)
						},
					},
					PushButton{
						Text: "Huffman",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							preHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Hamming/Huffman",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							preHammingHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Estadisticas",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							preStatisticsWindow(mw)
						},
					},
				},
			},
			VSpacer{
				Size: 20,
			},
			PushButton{
				Text: "Salir",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					exitWindow(mw)
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func preHammingWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Hamming",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Proteger archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							hammingWindow(mw)
						},
					},
					PushButton{
						Text: "Desproteger archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							deHammingWindow(mw)
						},
					},
					PushButton{
						Text: "Introducir errores",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							introduceErrorsWindow(mw)
						},
					},
				},
			},
			PushButton{
				Text: "Volver",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func preHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Huffman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							huffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Descomprimir archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							deHuffmanWindow(mw)
						},
					},
				},
			},
			PushButton{
				Text: "Volver",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func preHammingHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Hamming/Huffman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir y proteger archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							hammingHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Desproteger y desproteger archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							deHammingHuffmanWindow(mw)
						},
					},
				},
			},
			PushButton{
				Text: "Volver",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func hammingWindow(window *walk.MainWindow) {
	window.Hide()
	var ano *walk.TextEdit
	var mes *walk.TextEdit
	var dia *walk.TextEdit
	var hora *walk.TextEdit
	var minutos *walk.TextEdit
	var segundos *walk.TextEdit
	var menuItems = []string{ // ComboBox項目リスト
		"Hamming 7 (Seguridad Muy Alta)",
		"Hamming 32 (Seguridad Alta)",
		"Hamming 1024 (Seguridad Medio)",
		"Hamming 32768 (Seguridad Baja)",
	}
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var comboBox *walk.ComboBox
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Hamming",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				Font:         Font{Family: family, PointSize: 11},
				AssignTo:     &comboBox,
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 10},
						AssignTo: &segundos,
						Text:     "Segundos",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Font: Font{Family: family, PointSize: 11},
						Text: "Proteger",
						OnClicked: func() {
							var day, month, year, hour, minutes, seconds int
							errs := make([]error, 6)
							size := comboBox.CurrentIndex()
							switch size {
							case 0:
								size = 7
							case 1:
								size = 32
							case 2:
								size = 1024
							case 3:
								size = 32768
							}
							fileName := url.Text()
							anoString := ano.Text()
							mesString := mes.Text()
							diaString := dia.Text()
							horaString := hora.Text()
							minutosString := minutos.Text()
							segundosString := segundos.Text()

							year, errs[2] = strconv.Atoi(anoString)
							month, errs[1] = strconv.Atoi(mesString)
							day, errs[0] = strconv.Atoi(diaString)
							hour, errs[3] = strconv.Atoi(horaString)
							minutes, errs[4] = strconv.Atoi(minutosString)
							seconds, errs[5] = strconv.Atoi(segundosString)
							//Check if the date have errors
							for i := 0; i < len(errs); i++ {
								if errs[i] != nil {
									showError(mw, "El formato de la fecha no es válido")
									return
								}
							}
							unixDate := convertDate(year, month, day, hour, minutes, seconds)
							err := preHamming(size, fileName, unixDate)
							if err != nil {
								showError(mw, err.Error())
							} else {
								showSuccess(mw, "El archivo fue protegido correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func deHammingWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var checkBox *walk.CheckBox
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "DeHamming",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Corregir Errores",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			CheckBox{
				Font:     Font{Family: family, PointSize: 12},
				AssignTo: &checkBox,
				Checked:  true,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Desproteger",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							fixErrors := checkBox.Checked()
							fileName := url.Text()
							err := preDeHamming(fileName, fixErrors)
							if err != nil {
								showError(mw, err.Error())
							} else {
								showSuccess(mw, "Archivo decodificado correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func introduceErrorsWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Introducir errores",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Font: Font{Family: family, PointSize: 11},
						Text: "Arrastrar archivo",
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Introducir errores",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							fileName := url.Text()
							err := introduceErrors(fileName)
							if err != nil {
								showError(mw, err.Error())
							} else {
								showSuccess(mw, "Errores introducidos correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func huffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var ano *walk.TextEdit
	var mes *walk.TextEdit
	var dia *walk.TextEdit
	var hora *walk.TextEdit
	var minutos *walk.TextEdit
	var segundos *walk.TextEdit
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Hufmman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &segundos,
						Text:     "Segundos",
					},
				},
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 50},
				Children: []Widget{
					PushButton{
						Text: "Comprimir",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/
							var day, month, year, hour, minutes, seconds int
							err := make([]error, 6)

							year, err[2] = strconv.Atoi(ano.Text())
							month, err[1] = strconv.Atoi(mes.Text())
							day, err[0] = strconv.Atoi(dia.Text())
							hour, err[3] = strconv.Atoi(hora.Text())
							minutes, err[4] = strconv.Atoi(minutos.Text())
							seconds, err[5] = strconv.Atoi(segundos.Text())

							//Check if the date have errors
							for i := 0; i < len(err); i++ {
								if err[i] != nil {
									showError(mw, "El formato de la fecha no es válido")
									return
								}
							}
							unixDate := convertDate(year, month, day, hour, minutes, seconds)
							errs := huffman(urlString, unixDate)
							if errs != nil {
								showError(mw, errs.Error())
							} else {
								showSuccess(mw, "El archivo fue comprimido correctamente")
							}

						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func deHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "DeHufmman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 50},
				Children: []Widget{
					PushButton{
						Text: "Descomprimir",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							errs := desHuffman(urlString)
							if errs != nil {
								showError(mw, errs.Error())
							} else {
								showSuccess(mw, "El archivo fue comprimido correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func hammingHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var ano *walk.TextEdit
	var mes *walk.TextEdit
	var dia *walk.TextEdit
	var hora *walk.TextEdit
	var minutos *walk.TextEdit
	var segundos *walk.TextEdit
	var comboBox *walk.ComboBox
	var menuItems = []string{ // ComboBox項目リスト
		"Hamming 7 (Seguridad Muy Alta)",
		"Hamming 32 (Seguridad Alta)",
		"Hamming 1024 (Seguridad Medio)",
		"Hamming 32768 (Seguridad Baja)",
	}
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Hamming/Huffman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				Font:         Font{Family: family, PointSize: 11},
				AssignTo:     &comboBox,
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &segundos,
						Text:     "Segundos",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir y proteger",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							var day, month, year, hour, minutes, seconds int
							errs := make([]error, 6)
							size := comboBox.CurrentIndex()
							switch size {
							case 0:
								size = 7
							case 1:
								size = 32
							case 2:
								size = 1024
							case 3:
								size = 32768
							}
							fileName := url.Text()
							anoString := ano.Text()
							mesString := mes.Text()
							diaString := dia.Text()
							horaString := hora.Text()
							minutosString := minutos.Text()
							segundosString := segundos.Text()

							year, errs[2] = strconv.Atoi(anoString)
							month, errs[1] = strconv.Atoi(mesString)
							day, errs[0] = strconv.Atoi(diaString)
							hour, errs[3] = strconv.Atoi(horaString)
							minutes, errs[4] = strconv.Atoi(minutosString)
							seconds, errs[5] = strconv.Atoi(segundosString)
							//Check if the date have errors
							for i := 0; i < len(errs); i++ {
								if errs[i] != nil {
									showError(mw, "El formato de la fecha no es válido")
									return
								}
							}
							unixDate := convertDate(year, month, day, hour, minutes, seconds)
							err := preHammingHuffman(size, fileName, unixDate)
							if err != nil {
								showError(mw, err.Error())
							} else {
								showSuccess(mw, "El archivo fue protegido correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func deHammingHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "DeHamming/DeHuffman",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño aplicado:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Desproteger y descomprimir",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							fileName := url.Text()
							err := preDeHammingDeHuffman(fileName)
							if err != nil {
								showError(mw, err.Error())
							} else {
								showSuccess(mw, "Archivo decodificado correctamente")
							}
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func preStatisticsWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Estadisticas de tamaño",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo original",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						Font:     Font{Family: family, PointSize: 11},
						AssignTo: &url,
					},
					PushButton{
						Text: "Arrastrar archivo",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 50},
				Children: []Widget{
					PushButton{
						Text: "Ver tamaños",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							statisticsWindow(mw, url.Text())
						},
					},
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 11},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()

}

func statisticsWindow(window *walk.MainWindow, url string) {
	window.Hide()
	var mw *walk.MainWindow
	answer := statistics(url)
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Estadisticas de tamaño",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[0],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[1],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[2],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[3],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[4],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[5],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[6],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[7],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[8],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[9],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[10],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[11],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[12],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[13],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[14],
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(255, 255, 255),
			},
			PushButton{
				Text: "Volver",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func dropFile(window *walk.MainWindow) string {
	window.SetEnabled(false)
	var mw *walk.MainWindow
	var ret string
	_ = MainWindow{
		Title:    "Seleccionar archivo",
		AssignTo: &mw,
		Layout:   VBox{},
		OnDropFiles: func(files []string) {
			ret = strings.Join(files, "\r\n")
			mw.Dispose()
		},
		Children: []Widget{
			Label{
				Text:      "Arrastre el archivo aquí abajo:",
				Font:      Font{Family: family, PointSize: 12, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			TextEdit{
				ReadOnly: true,
				Text:     "Suelte el archivo aquí",
				Font:     Font{Family: family, PointSize: 11},
			},
			PushButton{
				Text: "Cancelar",
				Font: Font{Family: family, PointSize: 11},
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()
	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-320)/2,
		(yScreen-240)/2,
		320,
		240,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
	window.SetEnabled(true)
	return ret
}

func convertDate(year int, month int, day int, hour int, minutes int, seconds int) []byte {
	//No error found then create the date
	parseMonth := time.Month(month)
	location, _ := time.LoadLocation("America/Argentina/Cordoba")
	auxDate := time.Date(year, parseMonth, day, hour, minutes, seconds, 0, location)
	auxUnixDate := auxDate.Unix()
	s := []byte(strconv.FormatInt(auxUnixDate, 10))
	unixDate := []byte(s)
	for i := len(unixDate); i < 10; i = len(unixDate) {
		unixDate = append([]byte{48}, unixDate...)
	}
	return unixDate
}

func showError(window *walk.MainWindow, text string) {
	window.SetEnabled(false)
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Error",
		AssignTo: &mw,
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Error!",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(238, 50, 19),
			},
			Label{
				Text:      text,
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(238, 50, 19),
			},
			PushButton{
				Text: "Volver",
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()
	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-480)/2,
		(yScreen-200)/2,
		480,
		200,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
	window.SetEnabled(true)
}

func showSuccess(window *walk.MainWindow, text string) {
	window.SetEnabled(false)
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Exito",
		AssignTo: &mw,
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Exito",
				Font:      Font{Family: family, PointSize: 20, Bold: true},
				TextColor: walk.RGB(25, 167, 40),
			},
			Label{
				Text:      text,
				Font:      Font{Family: family, PointSize: 11},
				TextColor: walk.RGB(25, 167, 40),
			},
			PushButton{
				Text: "Volver",
				OnClicked: func() {
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()
	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-480)/2,
		(yScreen-200)/2,
		480,
		200,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
	window.SetEnabled(true)
}

func exitWindow(window *walk.MainWindow) {
	window.Hide()
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{Width: 600, Height: 400},
		MaxSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{

			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "Equipo de desarrollo",
				Font:          Font{Family: family, PointSize: 12},
				TextColor:     walk.RGB(255, 255, 255),
			},
			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "DECENA, Facundo Matías --- facundo.decena@gmail.com",
				Font:          Font{Family: family, PointSize: 10},
				TextColor:     walk.RGB(255, 255, 255),
			},
			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "PELLEGRINO, Maximiliano --- maxi.101997@gmail.com",
				Font:          Font{Family: family, PointSize: 10},
				TextColor:     walk.RGB(255, 255, 255),
			},

			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "VERGES, Federico --- fede_16_98@hotmail.com",
				Font:          Font{Family: family, PointSize: 10},
				TextColor:     walk.RGB(255, 255, 255),
			},
			VSpacer{
				Size: 20,
			},
			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "Profesores a cargo",
				Font:          Font{Family: family, PointSize: 11},
				TextColor:     walk.RGB(255, 255, 255),
			},
			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "SILVESTRI, Mario Alfredo",
				Font:          Font{Family: family, PointSize: 9},
				TextColor:     walk.RGB(255, 255, 255),
			},
			TextLabel{
				TextAlignment: AlignHNearVNear,
				Text:          "MONTEJANO, German Antonio",
				Font:          Font{Family: family, PointSize: 9},
				TextColor:     walk.RGB(255, 255, 255),
			},
			VSpacer{
				Size: 40,
			},
			Label{
				Text:      "¿Esta seguro que desea salir?",
				Font:      Font{Family: family, PointSize: 16, Bold: true},
				TextColor: walk.RGB(255, 255, 255),
			},
			VSpacer{
				Size: 50,
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Volver",
						Font: Font{Family: family, PointSize: 16},
						OnClicked: func() {
							mw.Dispose()
							window.Show()
						},
					},
					PushButton{
						Text: "Salir",
						Font: Font{Family: family, PointSize: 16},
						OnClicked: func() {
							os.Exit(3)
						},
					},
				},
			},
			Label{
				Text:          "Universidad Nacional de San Luis, 2019",
				Font:          Font{Family: family, PointSize: 8},
				TextColor:     walk.RGB(255, 255, 255),
				TextAlignment: 3,
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-SizeW)/2,
		(yScreen-SizeH)/2,
		SizeW,
		SizeH,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()

}
