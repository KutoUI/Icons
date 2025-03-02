package icons

import (
	"context"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

//go:embed icons/*.svg
var folder embed.FS

type IconOptions struct {
	Color       string
	StrokeWidth float32
	Size        int
}

func getIcon(name string) ([]byte, error) {
	svg, err := folder.ReadFile("icons/" + name + ".svg")
	if err != nil {
		return nil, err
	}
	return svg, nil
}

func Icon(name string) func(IconOptions) templ.Component {
	return func(props IconOptions) templ.Component {
		return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			svg, err := getIcon(name)
			if err != nil {
				return err
			}

			//set sensible defaults
			if props.Color == "" {
				props.Color = "currentColor"
			}
			if props.StrokeWidth == 0 {
				props.StrokeWidth = 2
			}
			if props.Size == 0 {
				props.Size = 24
			}

			//clamp stroke width
			if props.StrokeWidth < 0.5 {
				props.StrokeWidth = 0.5
			} else if props.StrokeWidth > 3 {
				props.StrokeWidth = 3
			}

			//clamp size
			if props.Size < 16 {
				props.Size = 16
			} else if props.Size > 48 {
				props.Size = 48
			}

			stringSvg := string(svg)
			start := strings.Index(stringSvg, ">") + 1
			end := strings.LastIndex(stringSvg, "</svg>")
			stringSvg = strings.TrimSpace(stringSvg[start:end])

			configuredSvg := fmt.Sprintf(
				`<svg xmlns="http://www.w3.org/2000/svg" width="%s" height="%s" viewBox="0 0 24 24" fill="none" stroke="%s" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">%s</svg>`,
				strconv.Itoa(props.Size),
				strconv.Itoa(props.Size),
				props.Color,
				strconv.FormatFloat(float64(props.StrokeWidth), 'G', -1, 32),
				stringSvg,
			)

			_, err = w.Write([]byte(configuredSvg))
			return
		})
	}
}

//go:embed lucide-version.txt
var IconVersion string
