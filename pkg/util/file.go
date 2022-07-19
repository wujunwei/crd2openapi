package util

import "os"

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

//ReadDir todo 配置读取文件夹深度
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
		open, err := os.Open(d.Name())
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
			file, err := os.OpenFile(d.Name(), os.O_RDONLY, d.Mode())
			if err != nil {
				return nil, err
			}
			recursively, err := ReadDirRecursively(file)
			if err != nil {
				return nil, err
			}
			res = append(res, recursively...)
		} else {
			open, err := os.Open(d.Name())
			if err != nil {
				return nil, err
			}
			res = append(res, open)
		}
	}
	return res, nil
}
