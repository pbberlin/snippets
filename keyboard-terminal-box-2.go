package main

import "github.com/nsf/termbox-go"
import "fmt"

type key struct {
	x  int
	y  int
	ch rune
}

func hlpr1(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbPrintf(x, y int, fg, bg termbox.Attribute, format string,
	args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	hlpr1(x, y, fg, bg, s)
}

func paintCompleteKeyboard() {
	termbox.SetCell(0, 0, 0x250C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 0, 0x2510, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, 23, 0x2514, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(79, 23, 0x2518, termbox.ColorWhite, termbox.ColorBlack)

	for i := 1; i < 79; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 23, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 17, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 4, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}
	for i := 1; i < 23; i++ {
		termbox.SetCell(0, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(79, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := 5; i < 17; i++ {
		termbox.SetCell(1, i, 0x2588, termbox.ColorYellow, termbox.ColorYellow)
		termbox.SetCell(78, i, 0x2588, termbox.ColorYellow, termbox.ColorYellow)
	}

	tbPrintf(33, 1, termbox.ColorMagenta|termbox.AttrBold, termbox.ColorBlack, "Keyboard demo!")
	tbPrintf(21, 2, termbox.ColorMagenta, termbox.ColorBlack, "(press CTRL+X and then CTRL+Q to exit)")
	tbPrintf(15, 3, termbox.ColorMagenta, termbox.ColorBlack, "(press CTRL+X and then CTRL+C to change input mode)")

	inputmode := termbox.SetInputMode(termbox.InputCurrent)
	inputmode_str := ""
	switch {
	case inputmode&termbox.InputEsc != 0:
		inputmode_str = "termbox.InputEsc"
	case inputmode&termbox.InputAlt != 0:
		inputmode_str = "termbox.InputAlt"
	}

	if inputmode&termbox.InputMouse != 0 {
		inputmode_str += " | termbox.InputMouse"
	}
	tbPrintf(3, 18, termbox.ColorWhite, termbox.ColorBlack, "Input mode: %s", inputmode_str)
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
	tbPrintf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "Key: ")
	tbPrintf(8, 19, termbox.ColorYellow, termbox.ColorBlack, "decimal: %d", ev.Key)
	tbPrintf(8, 20, termbox.ColorGreen, termbox.ColorBlack, "hex:     0x%X", ev.Key)
	tbPrintf(8, 21, termbox.ColorCyan, termbox.ColorBlack, "octal:   0%o", ev.Key)
	tbPrintf(8, 22, termbox.ColorRed, termbox.ColorBlack, "string:  %s", funckeymap(ev.Key))

	tbPrintf(54, 19, termbox.ColorWhite, termbox.ColorBlack, "Char: ")
	tbPrintf(60, 19, termbox.ColorYellow, termbox.ColorBlack, "decimal: %d", ev.Ch)
	tbPrintf(60, 20, termbox.ColorGreen, termbox.ColorBlack, "hex:     0x%X", ev.Ch)
	tbPrintf(60, 21, termbox.ColorCyan, termbox.ColorBlack, "octal:   0%o", ev.Ch)
	tbPrintf(60, 22, termbox.ColorRed, termbox.ColorBlack, "string:  %s", string(ev.Ch))

	modifier := "none"
	if ev.Mod != 0 {
		modifier = "termbox.ModAlt"
	}
	tbPrintf(54, 18, termbox.ColorWhite, termbox.ColorBlack, "Modifier: %s", modifier)
}

func pretty_print_resize(ev *termbox.Event) {
	tbPrintf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "Resize event: %d x %d", ev.Width, ev.Height)
}

var counter = 0

func pretty_print_mouse(ev *termbox.Event) {
	tbPrintf(3, 19, termbox.ColorWhite, termbox.ColorBlack, "Mouse event: %d x %d", ev.MouseX, ev.MouseY)
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
	tbPrintf(43, 19, termbox.ColorWhite, termbox.ColorBlack, "Key: ")
	tbPrintf(48, 19, termbox.ColorYellow, termbox.ColorBlack, button, counter)
}

func dispatch_press(ev *termbox.Event) {
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	paintCompleteKeyboard()
	termbox.Flush()
	inputmode := 0
	ctrlxpressed := false
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlS && ctrlxpressed {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ && ctrlxpressed {
				break loop
			}
			if ev.Key == termbox.KeyCtrlC && ctrlxpressed {
				chmap := []termbox.InputMode{
					termbox.InputEsc | termbox.InputMouse,
					termbox.InputAlt | termbox.InputMouse,
					termbox.InputEsc,
					termbox.InputAlt,
				}
				inputmode++
				if inputmode >= len(chmap) {
					inputmode = 0
				}
				termbox.SetInputMode(chmap[inputmode])
			}
			if ev.Key == termbox.KeyCtrlX {
				ctrlxpressed = true
			} else {
				ctrlxpressed = false
			}

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			dispatch_press(&ev)
			pretty_print_press(&ev)
			termbox.Flush()
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			pretty_print_resize(&ev)
			termbox.Flush()
		case termbox.EventMouse:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			paintCompleteKeyboard()
			pretty_print_mouse(&ev)
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
