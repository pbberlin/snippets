package main

import "github.com/nsf/termbox-go"
import "fmt"

type key struct {
	x  int
	y  int
	ch rune
}

func helperPrintAtF(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func PrintAtF(x, y int, fg, bg termbox.Attribute, format string,
	args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	helperPrintAtF(x, y, fg, bg, s)
}

func paintCompleteKeyboard() {

	// edges
	termbox.SetCell(0, 0, 0x250C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 0, 0x2510, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, 23, 0x2514, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 23, 0x2518, termbox.ColorWhite, termbox.ColorBlack)

	// horizontal lines
	for i := 1; i < 79; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 23, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 17, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 4, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}

	// vertical lines
	for i := 1; i < 23; i++ {
		termbox.SetCell(0, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(79, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
	}

	// fat yellow vertical lines
	// for i := 5; i < 17; i++ {
	// 	termbox.SetCell(1, i, 0x2588, termbox.ColorYellow, termbox.ColorYellow)
	// 	termbox.SetCell(78, i, 0x2588, termbox.ColorYellow, termbox.ColorYellow)
	// }

	PrintAtF(21, 1, termbox.ColorMagenta|termbox.AttrBold, termbox.ColorBlack,
		"CTRL+C and then CTRL+Q to exit")
	PrintAtF(15, 2, termbox.ColorMagenta, termbox.ColorBlack,
		"CTRL+C and then CTRL+M to change input mode")

}

var fcmap = []string{
	"CTRL+2, CTRL+~",
	"CTRL+A",
	"CTRL+B",
	"CTRL+C",
	"CTRL+D",
	"CTRL+E",
	"CTRL+F",
	"CTRL+G",
	"CTRL+H, BACKSPACE",
	"CTRL+I, TAB",
	"CTRL+J",
	"CTRL+K",
	"CTRL+L",
	"CTRL+M, ENTER",
	"CTRL+N",
	"CTRL+O",
	"CTRL+P",
	"CTRL+Q",
	"CTRL+R",
	"CTRL+S",
	"CTRL+T",
	"CTRL+U",
	"CTRL+V",
	"CTRL+W",
	"CTRL+X",
	"CTRL+Y",
	"CTRL+Z",
	"CTRL+3, ESC, CTRL+[",
	"CTRL+4, CTRL+\\",
	"CTRL+5, CTRL+]",
	"CTRL+6",
	"CTRL+7, CTRL+/, CTRL+_",
	"SPACE",
}

var fkmap = []string{
	"F1",
	"F2",
	"F3",
	"F4",
	"F5",
	"F6",
	"F7",
	"F8",
	"F9",
	"F10",
	"F11",
	"F12",
	"INSERT",
	"DELETE",
	"HOME",
	"END",
	"PGUP",
	"PGDN",
	"ARROW UP",
	"ARROW DOWN",
	"ARROW LEFT",
	"ARROW RIGHT",
}

func funckeymap(k termbox.Key) string {
	if k == termbox.KeyCtrl8 {
		return "CTRL+8, BACKSPACE 2" /* 0x7F */
	} else if k >= termbox.KeyArrowRight && k <= 0xFFFF {
		return fkmap[0xFFFF-k]
	} else if k <= termbox.KeySpace {
		return fcmap[k]
	}
	return "UNKNOWN"
}

func pretty_print_press(ev *termbox.Event) {
	PrintAtF(3, 19, termbox.ColorWhite, termbox.ColorBlack, "Key: ")
	PrintAtF(8, 19, termbox.ColorYellow, termbox.ColorBlack, "decimal: %d", ev.Key)
	PrintAtF(8, 20, termbox.ColorGreen, termbox.ColorBlack, "hex:     0x%X", ev.Key)
	PrintAtF(8, 21, termbox.ColorCyan, termbox.ColorBlack, "octal:   0%o", ev.Key)
	PrintAtF(8, 22, termbox.ColorRed, termbox.ColorBlack, "string:  %s", funckeymap(ev.Key))

	PrintAtF(54, 19, termbox.ColorWhite, termbox.ColorBlack, "Char: ")
	PrintAtF(60, 19, termbox.ColorYellow, termbox.ColorBlack, "decimal: %d", ev.Ch)
	PrintAtF(60, 20, termbox.ColorGreen, termbox.ColorBlack, "hex:     0x%X", ev.Ch)
	PrintAtF(60, 21, termbox.ColorCyan, termbox.ColorBlack, "octal:   0%o", ev.Ch)
	PrintAtF(60, 22, termbox.ColorRed, termbox.ColorBlack, "string:  %s", string(ev.Ch))

	modifier := "none"
	if ev.Mod != 0 {
		modifier = "termbox.ModAlt"
	}
	PrintAtF(54, 18, termbox.ColorWhite, termbox.ColorBlack, "Modifier: %s", modifier)

	im := termbox.SetInputMode(termbox.InputCurrent)
	PrintAtF(3, 18, termbox.ColorWhite, termbox.ColorBlack, "Input mode: %v", im)

}

func printResizeEvent(ev *termbox.Event) {
	PrintAtF(3, 19, termbox.ColorWhite, termbox.ColorBlack,
		"Resize event: %d x %d", ev.Width, ev.Height)
}

var counter = 0

func printMouseEvent(ev *termbox.Event) {
	PrintAtF(3, 19, termbox.ColorWhite, termbox.ColorBlack,
		"Mouse event: %d x %d", ev.MouseX, ev.MouseY)
	button := ""
	switch ev.Key {
	case termbox.MouseLeft:
		button = "MouseLeft: %d"
	case termbox.MouseMiddle:
		button = "MouseMiddle: %d"
	case termbox.MouseRight:
		button = "MouseRight: %d"
	}
	counter++
	PrintAtF(43, 19, termbox.ColorWhite, termbox.ColorBlack, "Key: ")
	PrintAtF(48, 19, termbox.ColorYellow, termbox.ColorBlack, button, counter)
}

func dispatch_press(ev *termbox.Event) {
}

func termboxInitAndInputLoop() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlue)
	paintCompleteKeyboard()
	termbox.Flush()

	inpMode := 0
	dblCtrlInit := false
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if dblCtrlInit && ev.Key == termbox.KeyCtrlQ {
				break loop
			}
			if dblCtrlInit && ev.Key == termbox.KeyCtrlS {
				termbox.Sync()
			}
			if dblCtrlInit && ev.Key == termbox.KeyCtrlM {
				chmap := []termbox.InputMode{
					termbox.InputEsc | termbox.InputMouse,
					termbox.InputAlt | termbox.InputMouse,
					termbox.InputEsc,
					termbox.InputAlt,
				}
				inpMode++
				if inpMode >= len(chmap) {
					inpMode = 0
				}
				termbox.SetInputMode(chmap[inpMode])
			}

			if ev.Key == termbox.KeyCtrlC {
				dblCtrlInit = true
			} else {
				dblCtrlInit = false
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			dispatch_press(&ev)
			pretty_print_press(&ev)
			termbox.Flush()
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			printResizeEvent(&ev)
			termbox.Flush()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			printMouseEvent(&ev)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
	cTerm <- true

}

var cTerm = make(chan bool)

func main() {
	go termboxInitAndInputLoop()
	<-cTerm
	fmt.Printf("exiting...\n")
}
