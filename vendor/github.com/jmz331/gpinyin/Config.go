package gpinyin

import (
	"errors"
	"regexp"
	"strings"
)

// Load adds or updates entries in an existing map with string keys
// and string values using a configuration file.
//
// The filename paramter indicates the configuration file to load ...
// the dest parameter is the map that will be updated.
//
// The configuration file entries should be constructed in key=value
// syntax.  A # symbol at the beginning of a line indicates a comment.
// Blank lines are ignored.
func loadResource(filename string, dest map[string]string, reverse bool) error {
	var str string
	if filename == data_pinyin {
		str = pinyin_db
	} else if filename == data_Chinese_tas {
		str = chinese_db
	} else if filename == data_multi_pinyin {
		str = multi_pinyin_db
	}
	if len(str) == 0 {
		return errors.New("Resource load failed.")
	}
	if !strings.HasSuffix(str, "\n") {
		return errors.New("Config file does not end with a newline character.")
	}
	re := regexp.MustCompile(".*=.*")
	s2 := re.FindAllString(str, -1)

	for _, tempStr := range s2 {
		arr := strings.Split(tempStr, "=")
		if reverse {
			dest[arr[1]] = arr[0]
		} else {
			dest[arr[0]] = arr[1]
		}
	}

	return nil
}
