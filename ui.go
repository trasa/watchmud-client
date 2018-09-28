package main

import (
	"github.com/jroimartin/gocui"
	"io"
	"log"
	"os"
)

var activeClientUI *ClientUI

type ClientUI struct {
	g         *gocui.Gui
	client    *Client
	isTesting bool
}

func NewClientUI(client *Client) *ClientUI {
	activeClientUI = &ClientUI{
		client: client,
	}
	return activeClientUI
}

func (c *ClientUI) initUi() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()
	c.g = g

	g.Cursor = true
	g.SetManagerFunc(layout)

	if err := c.keybindings(); err != nil {
		panic(err)
	}

	// the main loop runs here:
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (c *ClientUI) outputView() io.Writer {
	if c.isTesting {
		return os.Stdout
	}
	v, err := c.g.View("output")
	if err != nil {
		panic(err)
	}
	return v
}

// not part of ClientUI because it just sets the state in g
// and has to have the parameters of *gocui.Gui anyway
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	outputHeightPercentage := 0.9
	bottomOfOutput := int(float64(maxY) * outputHeightPercentage)

	// tries to create the view, if it already exists throws an error
	if _, err := g.SetView("output", 0, 0, maxX-1, bottomOfOutput); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	// tries to create the view, if it already exists throws an error
	if inputView, err := g.SetView("input", 0, bottomOfOutput+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		inputView.Editable = true
		inputView.Wrap = true
		inputView.SetCursor(0, 0)
		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}
	return nil
}

func (c *ClientUI) keybindings() error {
	if err := c.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, c.quit); err != nil {
		return err
	}

	if err := c.g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, c.getLine); err != nil {
		return err
	}
	return nil
}

// args are required by function definition
func (c *ClientUI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// args are required by function definition
func (c *ClientUI) getLine(g *gocui.Gui, v *gocui.View) error {
	buf := v.Buffer()
	UIPrintf("> " + buf)
	c.client.processInput(buf)
	v.Clear()
	v.SetCursor(0, 0)
	return nil
}
