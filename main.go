package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/EmilyShepherd/ota-tool/pkg/payload"
	"github.com/EmilyShepherd/ota-tool/internal/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		list            bool
		partitions      string
		outputDirectory string
		inputDirectory  string
		concurrency     int
	)

	group := &cmd.PayloadCmd{}
	group.Register("list", &cmd.List{})
	group.Execute(os.Args[1:], nil)

	return

	flag.IntVar(&concurrency, "concurrency", 4, "Number of multiple workers to extract")
	flag.BoolVar(&list, "list", false, "Show list of partitions in payload.bin and exit")
	flag.StringVar(&outputDirectory, "output", "", "Set output directory for partitions")
	flag.StringVar(&inputDirectory, "input", "", "Set directory for existing partitions")
	flag.StringVar(&partitions, "partitions", "", "Only work with selected partitions (comma-separated)")
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}
	filename := flag.Arg(0)

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

	now := time.Now()

	var targetDirectory = outputDirectory
	if targetDirectory == "" {
		targetDirectory = fmt.Sprintf("extracted_%d%02d%02d_%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	}
	if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(targetDirectory, 0755); err != nil {
			log.Fatal("Failed to create target directory")
		}
	}
	var sourceDirectory = inputDirectory
	update.SetConcurrency(concurrency)

	if partitions != "" {
		if err := update.ExtractSelected(sourceDirectory, targetDirectory, strings.Split(partitions, ",")); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := update.ExtractAll(sourceDirectory, targetDirectory); err != nil {
			log.Fatal(err)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [inputfile]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}
