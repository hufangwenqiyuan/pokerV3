package readFile

import (
	"io/ioutil"
)

func ReadFile(path string) []byte {
	//file, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return data
}
