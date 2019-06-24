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
	SIZE_W = 600
	SIZE_H = 400
)

func main() {
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Principal",
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Hamming",
						OnClicked: func() {
							preHammingWindow(mw)
						},
					},
					PushButton{
						Text: "Huffman",
						OnClicked: func() {
							preHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Hamming/Huffman",
						OnClicked: func() {
							preHammingHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Estadisticas de tamaño",
						OnClicked: func() {
							preStatisticsWindow(mw)
						},
					},
				},
			},
			PushButton{
				Text: "Salir",
				OnClicked: func() {
					os.Exit(3)
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Proteger archivo",
						OnClicked: func() {
							hammingWindow(mw)
						},
					},
					PushButton{
						Text: "Desproteger archivo",
						OnClicked: func() {
							deHammingWindow(mw)
						},
					},
					PushButton{
						Text: "Introducir errores",
						OnClicked: func() {
							introduceErrorsWindow(mw)
						},
					},
				},
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Menu Huffman",
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir archivo",
						OnClicked: func() {
							huffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Descomprimir archivo",
						OnClicked: func() {
							deHuffmanWindow(mw)
						},
					},
				},
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir y proteger archivo",
						OnClicked: func() {
							hammingHuffmanWindow(mw)
						},
					},
					PushButton{
						Text: "Desproteger y desproteger archivo",
						OnClicked: func() {
							deHammingHuffmanWindow(mw)
						},
					},
				},
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
		"Hamming 7",
		"Hamming 32",
		"Hamming 1024",
		"Hamming 32768",
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				AssignTo:     &comboBox,
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
						AssignTo: &segundos,
						Text:     "Segundos",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
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
									break
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func deHammingWindow(window *walk.MainWindow) {
	window.Hide()
	var menuItems = []string{ // ComboBox項目リスト
		"Hamming 7",
		"Hamming 32",
		"Hamming 1024",
		"Hamming 32768",
	}
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var comboBox *walk.ComboBox
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño aplicado:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				AssignTo:     &comboBox,
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Corregir Errores",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			CheckBox{
				AssignTo: &checkBox,
				Checked:  true,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func introduceErrorsWindow(window *walk.MainWindow) {
	window.Hide()
	var menuItems = []string{ // ComboBox項目リスト
		"Hamming 7",
		"Hamming 32",
		"Hamming 1024",
		"Hamming 32768",
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
				Text:      "Introducir errores",
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño aplicado:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
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
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/

						},
					},
					PushButton{
						Text: "Volver",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
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
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/
						},
					},
					PushButton{
						Text: "Volver",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
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
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/
						},
					},
					PushButton{
						Text: "Volver",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
		"Hamming 7",
		"Hamming 32",
		"Hamming 1024",
		"Hamming 32768",
	}
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "Hamming/Huffman",
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
						OnClicked: func() {
							urlString = dropFile(mw)
							_ = url.SetText(urlString)
						},
					},
				},
			},
			Label{
				Text:      "Seleccione la fecha de decodificacion:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &ano,
						Text:     "Año",
					},
					TextEdit{
						AssignTo: &mes,
						Text:     "Mes",
					},
					TextEdit{
						AssignTo: &dia,
						Text:     "Dia",
					},
					TextEdit{
						AssignTo: &hora,
						Text:     "Hora",
					},
					TextEdit{
						AssignTo: &minutos,
						Text:     "Minutos",
					},
					TextEdit{
						AssignTo: &segundos,
						Text:     "Segundos",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Comprimir y proteger",
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/
						},
					},
					PushButton{
						Text: "Volver",
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
								}
							}
							unixDate := convertDate(year, month, day, hour, minutes, seconds)
							err := preHamming(size, fileName, unixDate)
							if err != nil {
								showError(mw, err.Error())
							}
							showSuccess(mw, "El archivo fue protegido correctamente")
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
	mw.Run()
}

func deHammingHuffmanWindow(window *walk.MainWindow) {
	window.Hide()
	var menuItems = []string{ // ComboBox項目リスト
		"Hamming 7",
		"Hamming 32",
		"Hamming 1024",
		"Hamming 32768",
	}
	var mw *walk.MainWindow
	var url *walk.TextEdit
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:      "DeHamming/DeHuffman",
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione el tamaño aplicado:",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			ComboBox{
				Model:        menuItems,
				CurrentIndex: 0,
			},
			Label{
				Text:      "Seleccione la ruta del archivo",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
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
						OnClicked: func() {
							/*mw.Dispose()
							window.Show()
							*/
						},
					},
					PushButton{
						Text: "Volver",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      "Seleccione la ruta del archivo original",
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			HSplitter{
				MaxSize: Size{Width: 600, Height: 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
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
						OnClicked: func() {
							statisticsWindow(mw, url.Text())
						},
					},
					PushButton{
						Text: "Volver",
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[0],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[1],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[2],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[3],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[4],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[5],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[6],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[7],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[8],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[9],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[10],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[11],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[12],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[13],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			Label{
				Text:      answer[14],
				Font:      Font{"Arial", 11, false, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
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
		(xScreen-SIZE_W)/2,
		(yScreen-SIZE_H)/2,
		SIZE_W,
		SIZE_H,
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
				Font:      Font{"Arial", 12, true, false, false, false},
				TextColor: walk.RGB(255, 255, 255),
			},
			TextEdit{
				ReadOnly: true,
				Text:     "Suelte el archivo aquí",
			},
			PushButton{
				Text: "Cancelar",
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(238, 50, 19),
			},
			Label{
				Text:      text,
				Font:      Font{"Arial", 11, false, false, false, false},
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
				Font:      Font{"Arial", 20, true, false, false, false},
				TextColor: walk.RGB(25, 167, 40),
			},
			Label{
				Text:      text,
				Font:      Font{"Arial", 11, false, false, false, false},
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
