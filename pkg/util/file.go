package util

import (
	"os"
	"path/filepath"
)

func ReadAllFile(path string, recursive bool) ([]*os.File, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_RDONLY, stat.Mode())
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		if recursive {
			return ReadDirRecursively(f)
		}
		return ReadDir(f)
	}
	return []*os.File{f}, nil
}

func ReadDir(dir *os.File) ([]*os.File, error) {
	Dirs, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var res []*os.File
	for _, d := range Dirs {
		if d.IsDir() {
			continue
		}
		open, err := os.Open(filepath.Join(dir.Name(), d.Name()))
		if err != nil {
			return nil, err
		}
		res = append(res, open)
	}
	return res, nil
}
func ReadDirRecursively(dir *os.File) ([]*os.File, error) {
	Dirs, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var res []*os.File
	for _, d := range Dirs {
		if d.IsDir() {
			file, err := os.OpenFile(filepath.Join(dir.Name(), d.Name()), os.O_RDONLY, d.Mode())
			if err != nil {
				return nil, err
			}
			recursively, err := ReadDirRecursively(file)
			if err != nil {
				return nil, err
			}
			res = append(res, recursively...)
		} else {
			open, err := os.Open(filepath.Join(dir.Name(), d.Name()))
			if err != nil {
				return nil, err
			}
			res = append(res, open)
		}
	}
	return res, nil
}
