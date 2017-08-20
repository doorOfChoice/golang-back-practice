package tools

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

//获取本机时区的时间
func GetLocalTime() time.Time {
	if l, err := time.LoadLocation("Local"); err == nil {
		return time.Now().In(l)
	}

	return time.Time{}
}

//获取本机时区的指定时间
func ParseLocalTime(layout, t string) (time.Time, error) {
	if l, err := time.LoadLocation("Local"); err == nil {
		if t, err := time.ParseInLocation(layout, t, l); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("转化失败")
}

func ParseSecondTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func FileDelete(filename string) bool {
	err := os.Remove(filename)

	return err == nil
}

func UniqIntArray(ints []int) []int {
	uniqMap := make(map[int]bool)
	var uniqArray []int
	for _, v := range ints {
		if _, have := uniqMap[v]; !have {
			uniqArray = append(uniqArray, v)
			uniqMap[v] = true
		}
	}

	return uniqArray
}

func UniqStringArray(strings []string) []string {
	uniqMap := make(map[string]bool)
	var uniqArray []string
	for _, v := range strings {
		if _, have := uniqMap[v]; !have {
			uniqArray = append(uniqArray, v)
			uniqMap[v] = true
		}
	}

	return uniqArray
}

func FilterString(reg string, strs ...string) bool {
	r, err := regexp.Compile(reg)

	if err != nil {
		return false
	}

	for _, v := range strs {
		if !r.MatchString(v) {
			return false
		}
	}

	return true
}
