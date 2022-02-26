package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/urfave/cli"
)

func Start(ctx *cli.Context) error {
	path := ctx.String("pid-file")

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0700)

	if err != nil {
		log.Fatalf("Got error while opening pid file: %s.\n", err)
	}

	log.Printf("Successfully opened pid file '%s'. Checking whether file contains pid...\n", path)

	return check(file, ctx)
}

func check(file *os.File, ctx *cli.Context) error {
	_pid, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("Got error while reading pid file: %s.\n", err)
	}

	if len(_pid) <= 0 {
		log.Print("File is empty. Starting daemon...\n")

		return fork(file, ctx)
	}

	log.Print("File is not empty. Parsing content...\n")

	pid, atoiErr := strconv.Atoi(string(_pid))

	if atoiErr != nil {
		log.Print("Could not parse file content. Truncating file...\n")

		err := file.Truncate(0)

		if err != nil {
			log.Fatalf("Got error while truncating pid file: %s.\n", err)
		}

		log.Print("Truncated successfully. Starting new daemon...\n")

		return fork(file, ctx)
	}

	log.Printf("Kneesocks already running, pid: %d.", pid)

	return nil
}

func fork(file *os.File, ctx *cli.Context) error {
	binary := ctx.String("binary")

	chroot := ctx.String("chroot")

	env := []string{
		fmt.Sprintf("config_path=%s", ctx.String("config")),
	}

	pid, err := syscall.ForkExec(binary, nil, &syscall.ProcAttr{
		Dir: chroot,
		Env: env,
	})

	if err != nil {
		log.Fatalf("Got error while starting daemon: %s.\n", err)
	}

	log.Printf("Started successfully with parameters: pid=%d, binary='%s', chroot='%s', env=%+v.\n", pid, binary, chroot, env)

	return savePid(pid, file)
}

func savePid(pid int, file *os.File) error {
	log.Print("Saving pid to the pid file...")

	_, err := file.Write([]byte(strconv.Itoa(pid)))

	if err != nil {
		log.Fatalf("Got error while writing to pid file: %s.\n", err)
	}

	log.Print("Saved successfully. Exiting...\n")

	return nil
}
