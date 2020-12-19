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
	Pathname string
}

func NewFile(pathname string) *Filex {
	return &Filex{Pathname: pathname}
}

func NewFile1(parent, child string) *Filex {
	return NewFile(filepath.Join(parent, child))
}

func NewFile2(parent *Filex, child string) *Filex {
	return NewFile(filepath.Join(parent.Pathname, child))
}

var PathSeparator = string(os.PathSeparator)

func (f *Filex) Walk(walkFn filepath.WalkFunc) error {
	return filepath.Walk(f.Pathname, walkFn)
}

func (f *Filex) VolumeName() string {
	return filepath.VolumeName(f.Pathname)
}

func (f *Filex) CanonicalPath() string {
	if PathSeparator == "/" {
		return strings.ReplaceAll(f.Pathname, "\\", PathSeparator)
	} else {
		canonicalPath := strings.ReplaceAll(f.Pathname, "/", PathSeparator)
		b, _ := regexp.MatchString(`^[a-zA-Z]:.*$`, canonicalPath)
		if b {
			return strings.ToUpper(canonicalPath[0:1]) + canonicalPath[1:]
		}
		return canonicalPath
	}
}

func (f *Filex) InvariantSeparatorsPath() string {
	return strings.ReplaceAll(f.Pathname, "\\", "/")
}
func (f *Filex) ListFiles() []*Filex {
	var listFiles []*Filex
	info, err := ioutil.ReadDir(f.Pathname)
	if err != nil {
		return listFiles
	}
	for _, fileInfo := range info {
		childFile := &Filex{Pathname: filepath.Join(f.Pathname, fileInfo.Name())}
		listFiles = append(listFiles, childFile)
	}
	return listFiles
}

func (f *Filex) List() []string {
	var list []string
	info, err := ioutil.ReadDir(f.Pathname)
	if err != nil {
		return list
	}
	for _, fileInfo := range info {
		list = append(list, filepath.Join(f.Pathname, fileInfo.Name()))
	}
	return list
}

func (f *Filex) Parent() string {
	return filepath.Dir(f.Pathname)
}

func (f *Filex) ParentFile() *Filex {
	return NewFile(f.Parent())
}

func (f *Filex) IsExist() bool {
	_, err := os.Stat(f.Pathname)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (f *Filex) IsDir() bool {
	info, err := os.Stat(f.Pathname)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (f *Filex) IsFile() bool {
	info, err := os.Stat(f.Pathname)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (f *Filex) Name() string {
	pathname := strings.ReplaceAll(f.Pathname, "\\", "/")
	index := strings.LastIndex(pathname, "/")
	if index < 0 {
		index = -1
	}
	return pathname[index+1:]
}

func (f *Filex) Extension() string {
	return filepath.Ext(f.Pathname)
}

func (f *Filex) NameWithoutExtension() string {
	return strings.TrimSuffix(f.Name(), f.Extension())
}

func (f *Filex) Delete() error {
	return os.Remove(f.Pathname)
}

func (f *Filex) Create() (*os.File, error) {
	return os.Create(f.Pathname)
}

func (f *Filex) Rename(newpath string) error {
	return os.Rename(f.Pathname, newpath)
}

func (f *Filex) Open() (*os.File, error) {
	return os.Open(f.Pathname)
}

func (f *Filex) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(f.Pathname, flag, perm)
}

func (f *Filex) Length() int64 {
	info, err := os.Stat(f.Pathname)
	if err != nil {
		return 0
	}
	return info.Size()
}

func (f *Filex) LastModified() time.Time {
	info, err := os.Stat(f.Pathname)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}

func (f *Filex) Mode() os.FileMode {
	info, err := os.Stat(f.Pathname)
	if err != nil {
		return 0
	}
	return info.Mode()
}

func (f *Filex) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(f.Pathname, perm)
}

func (f *Filex) Mkdir(perm os.FileMode) error {
	return os.Mkdir(f.Pathname, perm)
}

func (f *Filex) Write(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.Pathname, data, perm)
}

func (f *Filex) WriteString(data string, perm os.FileMode) error {
	return ioutil.WriteFile(f.Pathname, []byte(data), perm)
}

func (f *Filex) ReadAll() ([]byte, error) {
	return ioutil.ReadFile(f.Pathname)
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
