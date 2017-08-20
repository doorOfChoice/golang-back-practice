package tools

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func SplitFilename(filename string) (string, string, error) {
	n := strings.LastIndex(filename, ".")
	if n == -1 {
		return "", "", fmt.Errorf("不是有效的文件")
	}
	prefix := filename[:n]
	suffix := filename[n+1:]

	return prefix, suffix, nil
}

func Sha256Filename(filename string, mixes ...string) (string, error) {
	pre, suf, err := SplitFilename(filename)
	if err != nil {
		return "", err
	}

	if len(mixes) != 0 {
		for _, v := range mixes {
			pre += v
		}
	}
	shaPre := Sha256([]byte(pre))

	return fmt.Sprintf("%s.%s", shaPre, suf), nil
}

func Sha256FilenameByData(h *multipart.FileHeader, dir string) (string, error) {
	_, suf, err := SplitFilename(h.Filename)
	if err != nil {
		return "", err
	}
	file, _ := Sha256FileData(h)

	return dir + file + "." + suf, nil
}
