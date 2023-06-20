package loader

import (
	"fmt"
	"io/ioutil"
	"os"
)

func openJson(path string) (*os.File, error) {
	//Opens a json file
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s\n", path)
	// defer the closing of our jsonFile so that we can parse it later on

	return jsonFile, err
}
func Load(path string) ([]byte, error) {
	file, err := openJson(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return byteValue, err
}
