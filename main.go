package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/go-findfont"
	"github.com/alexflint/go-arg"
)

const maxWidthPercent = .9

var font *giu.FontInfo

func tryLoadFont() {
	fontName := os.Getenv("NOTIFY_TTF_FONT")
	if fontName == "" {
		fontName = "Inconsolata-Medium.ttf"
	}
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		panic(err)
	}
	sizeStr := os.Getenv("NOTIFY_SIZE_FONT")
	if sizeStr == "" {
		sizeStr = "32"
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 32
	}
	data, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	font = giu.AddFontFromBytes(fontName, data, float32(size))
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

func loop(start time.Time, delay time.Duration, message string, prompt bool, windowWidth, windowHeight float32) {
	message = message + "\n"
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
	for _, line := range strings.Split(message, "\n") {
		line = strings.Trim(line, " ")
		layout = append(layout, giu.Row(
			giu.Align(giu.AlignCenter).To(giu.Style().SetFont(font).To(giu.Label(line))),
		))
	}
	giu.SingleWindow().Layout(layout)
}

type Args struct {
	Message      string  `arg:"positional" help:"the message to display on screen"`
	Prompt       bool    `arg:"-p,--prompt" help:"prompt the user for a y/n response, and exit 0/1 accordingly"`
	DelaySeconds float32 `arg:"-d,--delay-seconds" help:"delay seconds before accepting user input for prompted y/n" default:"1"`
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
	tryLoadFont()
	wnd := giu.NewMasterWindow("notify", 400, 200, giu.MasterWindowFlagsMaximized)
	windowWidth, windowHeight := wnd.GetSize()
	start := time.Now()
	delay := time.Duration(int64(args.DelaySeconds*1000)) * time.Millisecond
	message := args.Message + "\n\n"
	go func() {
		for {
			message = args.Message + "\n\n"
			delaySeconds := delay.Seconds()
			elapsed := time.Since(start).Seconds()
			if elapsed > delaySeconds {
				elapsed = delaySeconds
			}
			remaining := delay.Seconds() - elapsed
			if remaining < 0 {
				remaining = 0
			}
			for i := float64(0); i < elapsed; i += .025 {
				message += " "
			}
			for i := float64(0); i < remaining; i += .025 {
				message += "="
			}
			giu.Update()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	wnd.Run(func() {
		loop(start, delay, message, args.Prompt, float32(windowWidth), float32(windowHeight))
	})
}
