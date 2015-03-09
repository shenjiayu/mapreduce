package src

import (
	"bufio"
	"io"
	"strings"
)

type FileScanner struct {
	Reader *bufio.Reader
}

func NewScanner(fd io.Reader) *FileScanner {
	return &FileScanner{bufio.NewReader(fd)}
}

func (this *FileScanner) Next() []byte {
	if line, _, err := this.Reader.ReadLine(); err != nil {
		return nil
	} else {
		return line
	}
}

func (this *FileScanner) Find(user_id string) (string, bool) {
	for {
		if line := this.Next(); line != nil {
			values := strings.Split(string(line), ":")
			id := values[0]
			if user_id == id {
				return values[1], true
			}
		} else {
			return "", false
		}
	}
}
