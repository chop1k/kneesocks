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
	"socks/pkg/protocol/v5"
	"socks/test/stand/config"
	"syscall"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Picture struct {
	config config.Config
	_case  config.Case
	t      *testing.T
	parser v5.Parser
}

func NewPicture(
	config config.Config,
	_case config.Case,
	t *testing.T,
	parser v5.Parser,
) (Picture, error) {
	return Picture{
		config: config,
		_case:  _case,
		t:      t,
		parser: parser,
	}, nil
}

func (p Picture) CompareUsingTcp(picture byte, conn net.Conn) {
	p.compareUsingTcp(p.getPath(picture), picture, conn)
}

func (p Picture) CompareUsingUdp(picture byte, packet net.PacketConn) {
	p.compareUsingUdp(p.getPath(picture), picture, packet)
}

func (p Picture) getPath(picture byte) string {
	if picture == 1 {
		return p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "big-picture"))
	} else if picture == 2 {
		return p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "middle-picture"))
	} else if picture == 3 {
		return p.generateFilePath(fmt.Sprintf("%s-%s-%s", p._case.Protocol, p._case.Command, "small-picture"))
	} else {
		require.Fail(p.t, "Unknown picture %d. ", picture)

		return ""
	}
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

func (p Picture) compareUsingTcp(path string, picture byte, conn net.Conn) {
	writers, file, h := p.createWriter(path)

	p.receivePictureUsingTcp(conn, writers)

	p.compareHash(path, file, picture, h)
}

func (p Picture) compareUsingUdp(path string, picture byte, conn net.PacketConn) {
	writers, file, h := p.createWriter(path)

	p.receivePictureUsingUdp(conn, writers)

	p.compareHash(path, file, picture, h)
}

func (p Picture) createWriter(path string) (io.Writer, *os.File, hash.Hash) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)

	require.NoError(p.t, err)

	h := sha256.New()

	return io.MultiWriter(file, h), file, h
}

func (p Picture) receivePictureUsingTcp(reader io.Reader, writer io.Writer) {
	for {
		buffer := make([]byte, 512)

		i, err := reader.Read(buffer)

		if err != nil {
			break
		}

		_, err = writer.Write(buffer[:i])

		require.NoError(p.t, err)
	}
}

func (p Picture) receivePictureUsingUdp(packet net.PacketConn, writer io.Writer) {
	buffer := make([]byte, 60000)

	i, _, err := packet.ReadFrom(buffer)

	require.NoError(p.t, err)

	chunk, parseErr := p.parser.ParseUdpRequest(buffer[:i])

	require.NoError(p.t, parseErr)

	_, err = writer.Write(chunk.Data)

	require.NoError(p.t, err)
}

func (p Picture) compareHash(path string, file *os.File, picture byte, h hash.Hash) {
	if picture == 1 {
		require.Equal(p.t, p.config.Picture.BigPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 2 {
		require.Equal(p.t, p.config.Picture.MiddlePictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	} else if picture == 3 {
		require.Equal(p.t, p.config.Picture.SmallPictureHash, fmt.Sprintf("%x", h.Sum(nil)))
	}

	p.cleanUp(path, file, h)
}

func (p Picture) cleanUp(path string, file *os.File, sha hash.Hash) {
	sha.Reset()

	require.NoError(p.t, file.Close())

	require.NoError(p.t, os.Remove(path))
}
