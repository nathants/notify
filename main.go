package main

import (
	"os"
	"strings"
	"time"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/go-findfont"
	"github.com/AllenDang/imgui-go"
	"github.com/alexflint/go-arg"
)

const maxWidthPercent = .9

func tryLoadFont() {
	font := os.Getenv("NOTIFY_TTF_FONT")
	if font == "" {
		font = "Inconsolata-Medium.ttf"
	}
	fonts := giu.Context.IO().Fonts()
	fontPath, err := findfont.Find(font)
	if err == nil {
		fonts.AddFontFromFileTTFV(fontPath, 24, imgui.DefaultFontConfig, fonts.GlyphRangesDefault())
	}
}

func keypress(start time.Time, delay time.Duration, prompt bool) {
	if prompt {
		if time.Since(start) < delay {
			//
		} else if giu.IsKeyPressed(giu.KeyN) {
			os.Exit(1)
		} else if giu.IsKeyPressed(giu.KeyY) {
			os.Exit(0)
		}
	} else if giu.IsKeyPressed(giu.KeyQ) || giu.IsKeyPressed(giu.KeyEnter) {
		os.Exit(0)
	}
}

func width(s string) float32 {
	width, _ := giu.CalcTextSize(s)
	return width
}

func height(s string) float32 {
	_, height := giu.CalcTextSize(s)
	return height
}

func wrap(s string, windowWidth float32) string {
	wrapped := ""
	parts := strings.Split(s, " ")
	line := ""
	for {
		if len(parts) == 0 {
			break
		}
		if width(line+parts[0]) > windowWidth*maxWidthPercent {
			wrapped += "\n" + line
			line = ""
		}
		line += " " + parts[0]
		parts = parts[1:]
	}
	if line != "" {
		wrapped += "\n" + line
	}
	return wrapped
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func loop(start time.Time, delay time.Duration, message string, prompt bool, windowWidth, windowHeight float32, center bool) {
	if width(message) > windowWidth*maxWidthPercent {
		message = wrap(message, windowWidth)
	}
	heightOffset := float32(0)
	if height(message)*2 < windowWidth {
		heightOffset = (windowHeight - height(message)) / 2
	}
	layout := giu.Layout{
		giu.Custom(func() { keypress(start, delay, prompt) }),
		giu.Dummy(0, heightOffset),
	}
	if center {
		for _, line := range strings.Split(message, "\n") {
			line = strings.Trim(line, " ")
			layout = append(layout, giu.Row(
				giu.Dummy((windowWidth-width(line))/2, 0),
				giu.Label(line)),
			)
		}
	} else {
		maxWidth := float32(0)
		for _, line := range strings.Split(message, "\n") {
			maxWidth = max(maxWidth, width(line))
		}
		for _, line := range strings.Split(message, "\n") {
			line = strings.Trim(line, " ")
			layout = append(layout, giu.Row(
				giu.Dummy((windowWidth-maxWidth)/2, 0),
				giu.Label(line)),
			)
		}
	}
	giu.SingleWindow("notify").Layout(layout)
}

type Args struct {
	Message      string  `arg:"positional" help:"the message to display on screen"`
	Prompt       bool    `arg:"-p,--prompt" help:"prompt the user for a y/n response, and exit 0/1 accordingly"`
	DelaySeconds float32 `arg:"-d,--delay-seconds" help:"delay seconds before accepting user input for prompted y/n"`
	Center       bool    `arg:"-c,--center" help:"horizontally center each line"`
}

func (Args) Description() string {
	return "\nnotify the user of a message with a fullscreen popup. hit Q or ENTER to exit.\n"
}

func main() {
	var args Args
	arg.MustParse(&args)
	args.Message = strings.Replace(args.Message, "\\n", "\n", -1)
	if args.Prompt {
		args.Message += "\n\nproceed? y/n"
	}
	wnd := giu.NewMasterWindow("notify", 400, 200, giu.MasterWindowFlagsMaximized, tryLoadFont)
	windowWidth, windowHeight := wnd.GetSize()
	start := time.Now()
	delay := time.Duration(int64(args.DelaySeconds*1000)) * time.Millisecond
	wnd.Run(func() {
		loop(start, delay, args.Message, args.Prompt, float32(windowWidth), float32(windowHeight), args.Center)
	})
}
