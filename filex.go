package filex

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Filex struct {
	pathname string
}

func NewFile(pathname string) *Filex {
	return &Filex{pathname: pathname}
}

func NewFile1(parent, child string) *Filex {
	return NewFile(filepath.Join(parent, child))
}

func NewFile2(parent *Filex, child string) *Filex {
	return NewFile(filepath.Join(parent.pathname, child))
}

var PathSeparator = string(os.PathSeparator)

func (f *Filex) Walk(walkFn filepath.WalkFunc) error {
	return filepath.Walk(f.pathname, walkFn)
}

func (f *Filex) VolumeName() string {
	return filepath.VolumeName(f.pathname)
}

func (f *Filex) CanonicalPath() string {
	if PathSeparator == "/" {
		return strings.ReplaceAll(f.pathname, "\\", PathSeparator)
	} else {
		canonicalPath := strings.ReplaceAll(f.pathname, "/", PathSeparator)
		b, _ := regexp.MatchString(`^[a-zA-Z]:.*$`, canonicalPath)
		if b {
			return strings.ToUpper(canonicalPath[0:1]) + canonicalPath[1:]
		}
		return canonicalPath
	}
}

func (f *Filex) InvariantSeparatorsPath() string {
	return strings.ReplaceAll(f.pathname, "\\", "/")
}
func (f *Filex) ListFiles() []*Filex {
	var listFiles []*Filex
	info, err := ioutil.ReadDir(f.pathname)
	if err != nil {
		return listFiles
	}
	for _, fileInfo := range info {
		childFile := &Filex{pathname: filepath.Join(f.pathname, fileInfo.Name())}
		listFiles = append(listFiles, childFile)
	}
	return listFiles
}

func (f *Filex) List() []string {
	var list []string
	info, err := ioutil.ReadDir(f.pathname)
	if err != nil {
		return list
	}
	for _, fileInfo := range info {
		list = append(list, filepath.Join(f.pathname, fileInfo.Name()))
	}
	return list
}

func (f *Filex) Parent() string {
	return filepath.Dir(f.pathname)
}

func (f *Filex) ParentFile() *Filex {
	return NewFile(f.Parent())
}

func (f *Filex) IsExist() bool {
	_, err := os.Stat(f.pathname)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (f *Filex) IsDir() bool {
	info, err := os.Stat(f.pathname)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (f *Filex) IsFile() bool {
	info, err := os.Stat(f.pathname)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (f *Filex) Name() string {
	pathname := strings.ReplaceAll(f.pathname, "\\", "/")
	index := strings.LastIndex(pathname, "/")
	if index < 0 {
		index = -1
	}
	return pathname[index+1:]
}

func (f *Filex) Extension() string {
	return filepath.Ext(f.pathname)
}

func (f *Filex) NameWithoutExtension() string {
	return strings.TrimSuffix(f.Name(), f.Extension())
}

func (f *Filex) Delete() error {
	return os.Remove(f.pathname)
}

func (f *Filex) Create() (*os.File, error) {
	return os.Create(f.pathname)
}

func (f *Filex) Rename(newpath string) error {
	return os.Rename(f.pathname, newpath)
}

func (f *Filex) Open() (*os.File, error) {
	return os.Open(f.pathname)
}

func (f *Filex) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(f.pathname, flag, perm)
}

func (f *Filex) Length() int64 {
	info, err := os.Stat(f.pathname)
	if err != nil {
		return 0
	}
	return info.Size()
}

func (f *Filex) LastModified() time.Time {
	info, err := os.Stat(f.pathname)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

func (f *Filex) Mode() os.FileMode {
	info, err := os.Stat(f.pathname)
	if err != nil {
		return 0
	}
	return info.Mode()
}

func (f *Filex) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(f.pathname, perm)
}

func (f *Filex) Mkdir(perm os.FileMode) error {
	return os.Mkdir(f.pathname, perm)
}

func (f *Filex) Write(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.pathname, data, perm)
}

func (f *Filex) WriteString(data string, perm os.FileMode) error {
	return ioutil.WriteFile(f.pathname, []byte(data), perm)
}

func (f *Filex) ReadAll() ([]byte, error) {
	return ioutil.ReadFile(f.pathname)
}

func (f *Filex) ReadAllString() (string, error) {
	b, err := f.ReadAll()
	return string(b), err
}

func (f *Filex) AppendText(file *os.File, text string) (int, error) {
	return file.WriteAt([]byte(text), f.Length())
}

func (f *Filex) Append(file *os.File, data []byte) (int, error) {
	return file.WriteAt(data, f.Length())
}
