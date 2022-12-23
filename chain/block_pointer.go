package chain

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
)

type BlockPointer interface {
	Update(uint64) error
	Create() error
	Exists() bool
	Read() (uint64, error)
	Inc() error
}

type fileBlockPointer struct {
	min      uint64
	fileDir  string
	fileName string
}

func NewFileBlockPointer(path, name string, min uint64) BlockPointer {
	bp := new(fileBlockPointer)
	bp.fileDir = path
	bp.fileName = name
	bp.min = min
	return bp
}

func (fp *fileBlockPointer) Update(n uint64) error {
	p := strconv.FormatUint(n, 10)
	f, err := os.OpenFile(path.Join(fp.fileDir, fp.fileName), os.O_WRONLY, 0757)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	_, err = w.Write([]byte(p))
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func (fp *fileBlockPointer) Exists() bool {
	_, err := os.Open(path.Join(fp.fileDir, fp.fileName))
	return err == nil
}

func (fp *fileBlockPointer) Create() error {
	if _, err := os.Stat(fp.fileDir); err != nil {
		err = os.MkdirAll(fp.fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	fil, err := os.Create(path.Join(fp.fileDir, fp.fileName))
	if err != nil {
		return err
	}
	defer fil.Close()
	// new position file created, minimum position is returned
	_, err = fil.Write([]byte(strconv.FormatUint(fp.min, 10)))
	if err != nil {
		return err
	}
	return nil
}

func (fp *fileBlockPointer) Read() (uint64, error) {
	fil, err := os.Open(path.Join(fp.fileDir, fp.fileName))
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(fil)
	var block string
	for scanner.Scan() {
		block = strings.Trim(scanner.Text(), "\n")
	}
	blockI, err := strconv.ParseUint(block, 0, 64)
	if err != nil {
		return 0, err
	}
	return blockI, nil
}

func (fp *fileBlockPointer) Inc() error {
	r, err := fp.Read()
	if err != nil {
		return err
	}
	return fp.Update(r + 1)
}
