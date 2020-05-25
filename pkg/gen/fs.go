package gen

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// Open opens the given file with the provided flag, and
// created the nested directories and file if not
// exists.
func Open(fname string, flag int) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dirpath := filepath.Join(dir, path.Dir(fname))
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		return nil, err
	}

	filepath := filepath.Join(dir, fname)
	file, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
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
	r, err = os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
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

	_, err := w.Write(b2)
	if err != nil {
		return err
	}
	return nil
}
