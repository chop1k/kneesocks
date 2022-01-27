package picture

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/require"
	"hash"
	"io"
	"io/fs"
	"math/rand"
	"net"
	"os"
	"socks/test/stand/config"
	"syscall"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Picture struct {
	config config.Config
	_case  config.Case
	t      *testing.T
}

func NewPicture(config config.Config, _case config.Case, t *testing.T) (Picture, error) {
	return Picture{config: config, _case: _case, t: t}, nil
}

func (p Picture) Compare(picture byte, conn net.Conn) {
	var path string

	if picture == 1 {
		path = p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "big-picture"))
	} else if picture == 2 {
		path = p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "middle-picture"))
	} else if picture == 3 {
		path = p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "small-picture"))
	} else {
		require.Fail(p.t, "Unknown picture %d. ", picture)
	}

	p.compare(path, picture, conn)
}

func (p Picture) generateFilePath(name string) string {
	i := 0

	for {
		i += 1

		path := fmt.Sprintf(p.config.Misc.TempFileNamePattern, p.config.Misc.TempDirPath, name, p.randomString())

		if p.checkFileExists(path) {
			return path
		}

		if i > 20 {
			break
		}
	}

	require.Fail(p.t, "Loop detected.")

	return ""
}

func (p Picture) checkFileExists(path string) bool {
	_, err := os.Stat(path)

	fsErr, ok := err.(*fs.PathError)

	if !ok {
		return false
	}

	errno, ko := fsErr.Err.(syscall.Errno)

	if !ko {
		return false
	}

	return errno == 2
}

func (p Picture) randomString() string {
	s := make([]rune, p.config.Misc.RandomSuffixLength)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func (p Picture) compare(path string, picture byte, conn net.Conn) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)

	require.NoError(p.t, err)

	h := sha256.New()

	writers := io.MultiWriter(file, h)

	for {
		buffer := make([]byte, 512)

		i, err := conn.Read(buffer)

		if err != nil {
			break
		}

		_, err = writers.Write(buffer[:i])

		require.NoError(p.t, err)
	}

	if picture == 1 {
		require.Equal(p.t, p.config.Picture.BigPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 2 {
		require.Equal(p.t, p.config.Picture.MiddlePictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 3 {
		require.Equal(p.t, p.config.Picture.SmallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	}

	p.cleanUp(path, file, h, conn)
}

func (p Picture) cleanUp(path string, file *os.File, sha hash.Hash, conn net.Conn) {
	sha.Reset()

	require.NoError(p.t, file.Close())

	require.NoError(p.t, os.Remove(path))

	_ = conn.Close()
}
