package actions

import (
	"io"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/urfave/cli"
)

func Stop(ctx *cli.Context) error {
	path := ctx.String("pid-file")

	file, err := os.OpenFile(path, os.O_RDONLY, 0700)

	if err != nil {
		log.Fatalf("Got error while opening pid file: %s.\n", err)
	}

	log.Printf("Pid file '%s' opened successfully. Reading...\n", path)

	_pid, readErr := io.ReadAll(file)

	if readErr != nil {
		log.Fatalf("Got error while reading pid file: %s.\n", readErr)
	}

	log.Print("Read successfully. Parsing...\n")

	pid, atoiErr := strconv.Atoi(string(_pid))

	if atoiErr != nil {
		log.Fatalf("Got error while parsing pid file: %s.\n", atoiErr)
	}

	log.Printf("Parsed successfully, pid determined as %d, killing...\n", pid)

	killErr := syscall.Kill(pid, syscall.SIGINT)

	if killErr != nil {
		log.Fatalf("Got error while killing daemon: %s.\n", killErr)
	}

	log.Print("Killed successfully, deleting pid file...\n")

	rmErr := os.Remove(path)

	if err != nil {
		log.Fatalf("Got error while removing pid file: %s.\n", rmErr)
	}

	log.Print("Deleted successfully. Exiting...\n")

	return nil
}
