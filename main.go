package main

import (
    "github.com/wxnacy/goterminal"
    "github.com/nsf/termbox-go"
    "os/exec"
    "strings"
    "os"
)

func LogFile(str ...string) {
    file, _ := os.OpenFile("wsh.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    file.WriteString(strings.Join(str, " ") + "\n")
}

func switchCh(t *terminal.Terminal, e termbox.Event) {
    if e.Ch <= 0 {
        return
    }
    switch e.Ch {
        case 'j': {
            t.MoveCursor(0, 1)
        }
        case 'k': {
            t.MoveCursor(0, -1)
        }
    }
}

func GetStatusResult() string {
    gst := exec.Command("/bin/bash", "-c", "git status -s")
    bytes, err := gst.Output()
    if err != nil {
        panic(err)
    }

    content := string(bytes)
    return content
}

func InitCells(s string) [][]terminal.Cell {

    statusList := strings.Split(s, "\n")

    initCells := make([][]terminal.Cell, 0)
    for _, d := range statusList {
        if d == "\n" {
            continue
        }
        if strings.HasPrefix(d, " M") {
            cells := terminal.StringToCellsWithColor(
                d, terminal.ColorRed, terminal.ColorDefault,
            )
            initCells = append(initCells, cells[0])
        }  else if strings.HasPrefix(d, "A") {
            cells := terminal.StringToCellsWithColor(
                d, terminal.ColorYellow, terminal.ColorDefault,
            )
            initCells = append(initCells, cells[0])
        } else {
            cells := terminal.StringToCellsWithColor(
                strings.Replace(d, "??", " U", 1),
                terminal.ColorRed,
                terminal.ColorDefault,
            )
            initCells = append(initCells, cells[0])
        }
    }
    return initCells
}


func main() {
    t, err := terminal.New()
    if err != nil {
        panic(err)
    }

    statusList := GetStatusResult()
    for {
        t.SetCells(InitCells(statusList))
        t.SetCursorLineColor(terminal.ColorWhite, terminal.ColorYellow)
        t.Rendering()
        e := t.PollEvent()
        t.ListenKeyBorad(e)

        switchCh(t, e)
    }
}
