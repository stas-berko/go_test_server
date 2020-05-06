package utils

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type VisitorsData struct {
	Visitors int `json:"visitors"`
}

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
	//os.Exit(1)
}

func (newData *VisitorsData) RewriteStorageData(jsonStorage *os.File) {
	r, _ := json.Marshal(newData)
	err := jsonStorage.Truncate(0); Check(err)
	_, seekErr := jsonStorage.Seek(0, 0); Check(seekErr)
	_, writeErr := jsonStorage.Write(r); Check(writeErr)
	_, seekErr = jsonStorage.Seek(0, 0); Check(seekErr)

}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}


func InitStorage(path string) *os.File {
	ensureDir(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666); Check(err)
	fileInfo, err := file.Stat(); Check(err)

	if fileInfo.Size() == 0 {
		initData := &VisitorsData{Visitors: 0}
		initData.RewriteStorageData(file)
	}
	return file
}

func GetFromStorage(file *os.File) *VisitorsData {
	data := make([]byte, 0, 64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	err := scanner.Err(); Check(err)
	resp := VisitorsData{}

	decodeErr := json.Unmarshal(data, &resp); Check(decodeErr)
	return &resp
}
