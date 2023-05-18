package log

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	formatRegex = regexp.MustCompile(`\*+([^\r\n*]+)\*+`)

	mu sync.Mutex // Mutex for synchronization
)

func printFormattedMessage(color *color.Color, message string, args ...any) {

	fmtMessage := fmt.Sprintf(message, args...)

	matches := formatRegex.FindAllStringSubmatch(fmtMessage, -1)

	splits := formatRegex.Split(fmtMessage, -1)
	for i := 0; i < len(splits); i++ {
		fmt.Print(splits[i])
		if i < len(splits)-1 {
			color.Print(matches[i][1])
		}
	}
}

var (
	tag      string
	tagColor *color.Color
	g        = color.New(color.FgGreen)
	c        = color.New(color.FgHiBlue)
	y        = color.New(color.FgYellow)
	r        = color.New(color.FgRed)
)

func Initialize(title string, titleColor color.Attribute) {
	tag = title
	tagColor = color.New(titleColor)
}

func getCurrentTime() string {
	return time.Now().Format("15:04")
}

func Ok(message string, args ...any) {
	mu.Lock()         // Lock the mutex
	defer mu.Unlock() // Unlock the mutex when function returns

	tagColor.Printf("[%s]", tag)
	printFormattedMessage(g, fmt.Sprintf(" [%s] *OKAY* %s\n", getCurrentTime(), message), args...)
}

func Info(message string, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	tagColor.Printf("[%s]", tag)
	printFormattedMessage(c, fmt.Sprintf(" [%s] *INFO* %s\n", getCurrentTime(), message), args...)
}

func Warn(message string, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	tagColor.Printf("[%s]", tag)
	printFormattedMessage(y, fmt.Sprintf(" [%s] *WARN* %s\n", getCurrentTime(), message), args...)
}

func Err(message string, args ...any) {
	mu.Lock()
	defer mu.Unlock()

	tagColor.Printf("[%s]", tag)
	printFormattedMessage(r, fmt.Sprintf(" [%s] *ERR** %s\n", getCurrentTime(), message), args...)
	os.Exit(1)
}

func Response(rw http.ResponseWriter, message string, args ...any) {
	fmt.Fprintf(rw, message, args...)
}
