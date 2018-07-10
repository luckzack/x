package file

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

func WriteString(filePath string, s string) (int, error) {
	return WriteBytes(filePath, []byte(s))
}

func ReplaceContent(filepath string, old, new string) error {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	newContent := strings.Replace(string(buf), old, new, -1)
	return ioutil.WriteFile(filepath, []byte(newContent), 0)
}
