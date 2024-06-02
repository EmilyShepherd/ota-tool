package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/EmilyShepherd/ota-tool/pkg/payload"
	"github.com/EmilyShepherd/ota-tool/pkg/command"
)


type PayloadCmd struct {
	next command.Group[payload.Payload]
}

func (c *PayloadCmd) Register(name string, cmd command.Command[payload.Payload]) {
	c.next.Register(name, cmd)
}

func (c *PayloadCmd) Usage() string {
	return c.next.Usage()
}

func (c *PayloadCmd) Execute(f []string, _ any) {
	filename := f[0]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s\n", filename)
	}

	var update *payload.Payload
	var err error

	if strings.HasSuffix(filename, ".zip") {
		update, err = payload.NewPayloadFromZipFile(filename)
	} else {
		f, _ := os.Open(filename)
		update = payload.NewPayload(f)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	update.Init()

	c.next.Execute(f[1:], update)
}
