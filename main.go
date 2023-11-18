package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	c := NewChip8()
	c.Screen.SetSize(40, 40)
	for {
		switch ev := c.Screen.PollEvent().(type){
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape{
				c.Screen.Fini()
				os.Exit(0)
			}

		}

		go func() { for  i := 0; i < 40; i++ {
			for j:=0; j < 40; j++{
				c.Screen.SetContent(i, j, tcell.RuneDiamond, nil, tcell.Style{})
				c.Screen.Show()
				time.Sleep(1 * time.Second)
			}
		}
	}()
}
}
