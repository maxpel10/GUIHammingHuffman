package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

const (
	SIZE_W = 600
	SIZE_H = 400
)

func main() {
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:   "Práctico de máquina TI",
		AssignTo:&mw,
		MinSize: Size{600, 400},
		MaxSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Hamming",
						OnClicked: func() {
							hammingWindow(mw)
						},

					},
					PushButton{
						Text: "Huffman",
						OnClicked: huffmanWindow,
					},
					PushButton{
						Text: "Hamming/Huffman",
						OnClicked: hammingHuffmanWindow,
					},
					PushButton{
						Text: "Estadisticas de tamaño",
						OnClicked: statisticsWindows,
					},
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN);
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN);
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen - SIZE_W)/2,
		(yScreen - SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW);
	mw.Run()
}

func hammingWindow(window *walk.MainWindow){
	window.Hide()
	var mw *walk.MainWindow
	_ = MainWindow{
		Title:   "Práctico de máquina TI",
		AssignTo:&mw,
		MinSize: Size{600, 400},
		MaxSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Hamming 7",
						OnClicked: func() {
							hammingWindow(mw)
						},

					},
					PushButton{
						Text: "Hamming 32",
						OnClicked: huffmanWindow,
					},
					PushButton{
						Text: "Hamming 1024",
						OnClicked: hammingHuffmanWindow,
					},
					PushButton{
						Text: "Hamming32768",
						OnClicked: statisticsWindows,
					},
				},
			},
			PushButton{
				Text: "Volver",
				OnClicked: func(){
					mw.Dispose()
					window.Show()
				},
			},
		},
	}.Create()

	windowColor, _ := walk.NewSolidColorBrush(walk.RGB(58, 52, 51))
	mw.SetBackground(windowColor)

	win.SetWindowLong(mw.Handle(), win.GWL_STYLE, win.WS_BORDER) // removes default styling

	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN);
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN);
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen - SIZE_W)/2,
		(yScreen - SIZE_H)/2,
		SIZE_W,
		SIZE_H,
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW);

	mw.Run()

}

func huffmanWindow(){

}

func hammingHuffmanWindow(){

}

func statisticsWindows(){

}