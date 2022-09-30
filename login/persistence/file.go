package persistence

import "strings"

func NewFileName(name string) fileName {
	return fileName{name}
}

type fileName struct {
	fileName string
}

func (n *fileName) String() string {
	return n.fileName
}

func (n *fileName) Initialize() {
	if strings.TrimSpace(n.fileName) == "" {
		n.fileName = defaultFilename
	}
}
