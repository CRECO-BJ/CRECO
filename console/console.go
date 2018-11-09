package console

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterh/liner"
)

const historyCommandFile = ".creco_cmd_history"

var command = []string{""}

// Console ...
type Console struct {
	*liner.State
	abort chan struct{}
	exit  chan struct{}
}

// NewConsole ...
func NewConsole() *Console {
	var con = &Console{}
	con.State = liner.NewLiner()
	con.exit = make(chan struct{})
	con.abort = make(chan struct{})

	con.SetCtrlCAborts(true)
	con.SetCompleter(func(line string) (c []string) {
		for _, n := range command {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(filepath.Join(os.TempDir(), historyCommandFile)); err == nil {
		con.ReadHistory(f)
		f.Close()
	} else {
		log.Print("Error logging history command.\n")
	}

	go con.consoleLoop()

	return con
}

// CloseConsole ...
func (c *Console) CloseConsole() {
	// close linter
	c.Close()
	// write history
	if f, err := os.Create(filepath.Join(os.TempDir(), historyCommandFile)); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		c.WriteHistory(f)
		f.Close()
	}
	// stop console loop
	c.abort <- struct{}{}
	<-c.exit
}

func (c *Console) consoleLoop() {
	for {
		select {
		case <-c.abort:
			c.exit <- struct{}{}
			return
		default:
			if co, err := c.Prompt(":>"); err == nil {
				log.Print("Got: ", co)
				c.AppendHistory(co)
			} else if err == liner.ErrPromptAborted {
				log.Print("Aborted")
			} else {
				log.Print("Error reading line: ", err)
			}
		}
	}
}
