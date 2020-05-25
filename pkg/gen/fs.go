package gen

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// Open opens the given file with the provided flag, and
// created the nested directories and file if not
// exists.
func Open(name string, flag int) (*os.File, error) {
	dirpath, err := Resolve(path.Dir(name))
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		return nil, err
	}

	filepath, err := Resolve(name)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Resolve(name string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, name)
	return path, nil
}

// Create creates a file and its directories if not exists, and err when
// exists.
func Create(name string) error {
	f, err := Open(name, os.O_RDONLY|os.O_CREATE|os.O_EXCL)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// Read reads a file and creates it if it does not exist.
func Read(name string) ([]byte, error) {
	path, err := Resolve(name)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Write writes the template with the data to a given file, and creates the
// file if it does not exists.
func Write(out string, tpl []byte, data interface{}) error {
	t := template.Must(template.New("").Parse(string(tpl)))

	// Open as write-only, create if not exists.
	w, err := Open(out, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
	if err != nil {
		return err
	}
	defer w.Close()

	// Write to a temporary buffer.
	var b bytes.Buffer
	if err := t.Execute(&b, data); err != nil {
		return err
	}

	// Formats the go-code before writing.
	b2, err := FormatSource(b.Bytes())
	if err != nil {
		return err
	}

	_, err = w.Write(b2)
	if err != nil {
		return err
	}
	return nil
}

func Overwrite(name string, content []byte) error {
	f, err := Open(name, os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}

func RemoveIfExists(name string) error {
	path, err := Resolve(name)
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}
