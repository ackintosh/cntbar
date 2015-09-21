package main

import (
	"bufio"
	"os"
	"time"

	"github.com/ackintosh/cntbar/cntbar"
	"github.com/andrew-d/go-termutil"
	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

var (
	tickInterval = 1000 * time.Millisecond
	themeDefault = termui.ColorScheme{HasBorder: true}
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	if termutil.Isatty(os.Stdin.Fd()) {
		panic("...")
	}

	eventChan := make(chan termbox.Event)
	go func() {
		for {
			eventChan <- termbox.PollEvent()
		}
	}()

	dataChan := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			dataChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	summary := cntbar.NewSummary()
	tick := time.Tick(tickInterval)
	termui.UseTheme("helloworld")

	render(summary)

	for {
		select {
		case event := <-eventChan:
			if event.Type == termbox.EventKey && event.Ch == 'q' {
				return
			}
		case data := <-dataChan:
			summary.CountUp(data)
		case <-tick:
			render(summary)
		}
	}
}

func render(summary *cntbar.Summary) {
	data := summary.GetChart()
	if data != nil {
		termui.Render(data)
	}
}

func renderGauges(gauges []*termui.Gauge) {
	termbox.Clear(termbox.ColorDefault, termbox.Attribute(themeDefault.BodyBg))
	for _, g := range gauges {
		buf := g.Buffer()
		for _, v := range buf {
			termbox.SetCell(v.X, v.Y, v.Ch, termbox.Attribute(v.Fg), termbox.Attribute(v.Bg))
		}
	}
	termbox.Flush()
}
