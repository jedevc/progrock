package ui

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/vito/midterm"
)

type Vterm struct {
	Offset int
	Height int
	Width  int

	Prefix string

	vt *midterm.Terminal

	viewBuf *bytes.Buffer
}

func NewVterm() *Vterm {
	return &Vterm{
		vt:      midterm.NewAutoResizingTerminal(),
		viewBuf: new(bytes.Buffer),
	}
}

func (term *Vterm) Write(p []byte) (int, error) {
	atBottom := term.Offset+term.Height >= term.vt.UsedHeight()
	if term.Height == 0 {
		atBottom = true
	}

	n, err := term.vt.Write(p)
	if err != nil {
		return n, err
	}

	if atBottom {
		term.Offset = max(0, term.vt.UsedHeight()-term.Height)
	}

	return n, nil
}

func (term *Vterm) UsedHeight() int {
	return term.vt.UsedHeight()
}

func (term *Vterm) SetHeight(height int) {
	atBottom := term.Offset+term.Height >= term.vt.UsedHeight()

	term.Height = height

	if atBottom {
		term.Offset = max(0, term.vt.UsedHeight()-term.Height)
	}
}

func (term *Vterm) SetWidth(width int) {
	term.Width = width
	term.vt.AutoResizeX = false // stop auto-resizing vterm width
	prefixWidth := lipgloss.Width(term.Prefix)
	if term.Width > prefixWidth {
		term.vt.ResizeX(width - prefixWidth)
	}
}

func (term *Vterm) SetPrefix(prefix string) {
	term.Prefix = prefix
	prefixWidth := lipgloss.Width(prefix)
	if term.Width > prefixWidth && !term.vt.AutoResizeX {
		term.vt.ResizeX(term.Width - prefixWidth)
	}
}

func (term *Vterm) Init() tea.Cmd {
	return nil
}

func (term *Vterm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) { // nolint:gocritic
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, Keys.Up):
			term.Offset = max(0, term.Offset-1)
		case key.Matches(msg, Keys.Down):
			term.Offset = min(term.vt.UsedHeight()-term.Height, term.Offset+1)
		case key.Matches(msg, Keys.PageUp):
			term.Offset = max(0, term.Offset-term.Height)
		case key.Matches(msg, Keys.PageDown):
			term.Offset = min(term.vt.UsedHeight()-term.Height, term.Offset+term.Height)
		case key.Matches(msg, Keys.Home):
			term.Offset = 0
		case key.Matches(msg, Keys.End):
			term.Offset = term.vt.UsedHeight() - term.Height
		}
	}
	return term, nil
}

func (term *Vterm) ScrollPercent() float64 {
	return min(1, float64(term.Offset+term.Height)/float64(term.vt.UsedHeight()))
}

const reset = termenv.CSI + termenv.ResetSeq + "m"

// View returns the output for the current region of the terminal, with ANSI
// formatting.
func (term *Vterm) View() string {
	return string(term.Bytes(term.Offset, term.Height))
}

// Bytes returns the output for the given region of the terminal, with
// ANSI formatting.
func (term *Vterm) Bytes(offset, height int) []byte {
	used := term.vt.UsedHeight()
	if used == 0 {
		return nil
	}

	buf := term.viewBuf
	buf.Reset()

	var lines int
	for row := range term.vt.Content {
		if row < offset {
			continue
		}
		if row+1 > (offset + height) {
			break
		}

		buf.WriteString(term.Prefix)
		term.vt.RenderLine(buf, row)
		lines++

		if row > used {
			break
		}
	}

	return buf.Bytes()
}

// LastLine returns the last line of visible text, with ANSI formatting, but
// without any trailing whitespace.
func (term *Vterm) LastLine() string {
	used := term.vt.UsedHeight()
	if used == 0 {
		return ""
	}

	var lastLine string
	for row := used - 1; row >= 0; row-- {
		var lastFormat midterm.Format

		buf := new(strings.Builder)
		for col, r := range term.vt.Content[row] {
			f := term.vt.Format[row][col]

			if f != lastFormat {
				lastFormat = f
				buf.Write([]byte(f.Render()))
			}

			buf.Write([]byte(string(r)))
		}

		if strings.TrimSpace(buf.String()) == "" {
			continue
		}

		lastLine = strings.TrimRightFunc(buf.String(), unicode.IsSpace)

		break
	}

	return lastLine + reset
}

// Print prints the full log output without any formatting.
func (term *Vterm) Print(w io.Writer) error {
	used := term.vt.UsedHeight()

	for row, l := range term.vt.Content {
		_, err := fmt.Fprintln(w, strings.TrimRight(string(l), " "))
		if err != nil {
			return err
		}

		if row > used {
			break
		}
	}

	return nil
}
