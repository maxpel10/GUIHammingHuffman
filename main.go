package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"os"
	"strings"
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
							statisticsWindows(mw)
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
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
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
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
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
	var urlString string
	_ = MainWindow{
		Title:    "Práctico de máquina TI",
		AssignTo: &mw,
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
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
							url.SetText(urlString)
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
						Text: "Proteger",
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

func deHammingWindow(window *walk.MainWindow) {

}

func introduceErrorsWindow(window *walk.MainWindow) {

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
		MinSize:  Size{600, 400},
		MaxSize:  Size{600, 400},
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
				MaxSize: Size{600, 20},
				Children: []Widget{
					TextEdit{
						AssignTo: &url,
					},
					PushButton{
						Text: "Dropear archivo",
						OnClicked: func() {
							urlString = dropFile(mw)
							url.SetText(urlString)
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
				MaxSize: Size{600, 50},
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
	/*var url *walk.TextEdit
	var urlString string*/
}

func hammingHuffmanWindow(window *walk.MainWindow) {
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
							url.SetText(urlString)
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

func deHammingHuffmanWindow(window *walk.MainWindow) {
	/*var url *walk.TextEdit
	var urlString string*/
}

func statisticsWindows(window *walk.MainWindow) {

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
