package gen

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrInvalidPath = errors.New("path is invalid")

type Volume string

func (v Volume) Split() (string, string, error) {
	volumes := filepath.SplitList(string(v))
	if len(volumes) != 2 {
		return "", "", ErrInvalidPath
	}
	src, dst := volumes[0], volumes[1]
	src, dst = os.ExpandEnv(src), os.ExpandEnv(dst)
	return src, dst, nil
}
