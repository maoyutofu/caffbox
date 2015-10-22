package util

import (
	"github.com/tjz101/caffbox"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func DirExists(name string) (bool, error) {
	fileInfo, err := os.Stat(name)
	if err == nil {
		return fileInfo.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileExists(name string) (bool, error) {
	fileInfo, err := os.Stat(name)
	if err == nil {
		return !fileInfo.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func TimeToString(layout string) string {
	t := time.Now()
	return t.Format(layout)
}

func PathFromTime() string {
	return TimeToString("/2006/01/02")
}

func Mkdir(path string) error {
	dirFlag, _ := DirExists(path)
	if !dirFlag {
		return os.MkdirAll(path, 0665)
	}
	return nil
}

func GetFileSuffix(filename string) string {
	n := strings.LastIndex(filename, ".")
	if n == -1 {
		return ""
	}
	return filename[n:]
}

func RandomString() string {
	b := []byte("abcdefghijklmnopqrstuvwxyz")
	t := time.Now()
	rand.Seed(t.UnixNano())
	prefix := string(b[rand.Intn(len(b)-1)])
	unixstr := strconv.FormatInt(t.Unix(), 10)
	str := []string{prefix, unixstr}
	return strings.Join(str, "_")
}

func getFileIndex(localName, name string) (bool, int) {
	lastIndex := strings.LastIndex(name, ".")
	expr := "(_(\\d{1,}))?"
	if lastIndex != -1 {
		prefix := name[:lastIndex]
		suffix := name[lastIndex:]
		expr = prefix + expr + suffix
	} else {
		expr = name + expr
	}
	r := regexp.MustCompile(expr)
	if r.MatchString(localName) {
		index := r.FindStringSubmatch(localName)[2]
		i, _ := strconv.Atoi(index)
		return true, i
	}
	return false, 0
}

func isExist(localName, name string) bool {
	fla, _ := getFileIndex(localName, name)
	return fla
}

func MaxNum(arr []int) int {
	size := len(arr)
	if size == 0 {
		return 0
	}
	tmp := arr[0]
	for i := 0; i < size; i++ {
		if tmp < arr[i] {
			tmp = arr[i]
		}
	}
	return tmp
}

func IncreaseFilename(dir string, name string) (string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return name, err
	}
	index := []int{}
	for _, info := range infos {
		if !info.IsDir() {
			localName := info.Name()
			fla, i := getFileIndex(localName, name)
			if fla {
				index = append(index, i)
			}
		}
	}
	if len(index) > 0 {
		lastIndex := strings.LastIndex(name, ".")
		newname := name
		count := MaxNum(index) + 1
		if lastIndex != -1 {
			prefix := name[:lastIndex]
			suffix := name[lastIndex:]
			str := []string{prefix, suffix}
			newname = strings.Join(str, "_"+strconv.Itoa(count))
		} else {
			newname = name + "_" + strconv.Itoa(count)
		}
		return newname, nil
	}
	return name, nil
}

func GetPhysicalPath(path string) string {
	return filepath.Join(caffbox.RootPhysicalPath, path)
}

func GetAbsPath(path string) (string, error) {
	rel, err := filepath.Rel(caffbox.RootPhysicalPath, path)
	if err != nil {
		return "", err
	}
	return filepath.Join(caffbox.ROOT_DIR, rel), nil
}

func WritePid() error {
	pid := os.Getpid()
	f, err := os.OpenFile("./logs/caffbox.pid", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(pid))
	return err
}
