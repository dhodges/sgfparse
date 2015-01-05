package parse

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func sgfFileList(dirname string) ([]string, error) {
	var fileList []string
	dirname = strings.TrimSpace(dirname)
	if dirname[len(dirname)-1] != '/' {
		dirname = dirname + "/"
	}
	fileInfoList, err := ioutil.ReadDir(dirname)
	if err != nil {
		return fileList, err
	}
	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			fileList = append(fileList, dirname+fileInfo.Name())
		}
	}
	return fileList, err
}

func TestParsingSGF(t *testing.T) {
	dirname := "/Users/david/Google Drive/SGF/AWAGC-2014"
	fileList, err := sgfFileList(dirname)
	if err != nil {
		t.Error(fmt.Sprintf("error reading sgf file list: %s", err.Error()))
		return
	}
Loop:
	for _, fname := range fileList {
		buf, err := ioutil.ReadFile(fname)
		if err != nil {
			t.Error(fmt.Sprintf("Error reading file: %q, %q", fname, err.Error()))
			return
		}
		sgf := new(SGFGame)
		sgf.Parse(string(buf))
		if len(sgf.errors) > 0 {
			fmt.Printf("problems parsing file: %q\n", fname)
			fmt.Println(sgf.errors[0].Error())
			break Loop
		}
	}
}