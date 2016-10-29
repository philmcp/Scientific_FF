package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteFile(outputLoc string) {

	if _, err := os.Stat(outputLoc); os.IsExist(err) {
		err := os.Remove(outputLoc)

		if err != nil {
			log.Println(err)
		}
	}
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		log.Println(err)
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetLastName(str string) string {
	source := ParseEncoding(strings.ToLower(str))

	names := strings.Split(source, " ")
	name := names[len(names)-1]
	return name
}

func ParseEncoding(str string) string {
	str = strings.ToLower(str)

	str = strings.Replace(str, "è", "e", -1)
	str = strings.Replace(str, "á", "a", -1)
	str = strings.Replace(str, "é", "e", -1)
	str = strings.Replace(str, "ö", "o", -1)
	str = strings.Replace(str, "í", "i", -1)
	str = strings.Replace(str, "à", "a", -1)
	str = strings.Replace(str, "ó", "o", -1)
	str = strings.Replace(str, "ñ", "n", -1)

	str = strings.Replace(str, "ú", "u", -1)
	str = strings.Replace(str, "ü", "u", -1)
	return str
}
