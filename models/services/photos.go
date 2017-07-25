// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package services

import (
	"fmt"
	"github.com/ruslanBik4/httpgo/models/logs"
	"github.com/ruslanBik4/httpgo/models/server"
	"io"
	"log"
	"os"
	"path/filepath"
)

type photosService struct {
	name     string
	path     string
	fileName string
	status   string
}

var (
	photos *photosService = &photosService{name: "photos"}
)

func (photos *photosService) Init() error {
	photos.status = "starting"

	ServerConfig := server.GetServerConfig()

	photos.path = filepath.Join(ServerConfig.WWWPath(), "files/photos")
	photos.fileName = "test.txt"

	photos.status = "ready"
	return nil
}

// выполняет операцию по записи/чтении файла
// это зависит от первого параметра - "save" или "read"
// третий параметр - имя файла
func (photos *photosService) Send(args ...interface{}) error {

	var oper string
	if len(args) < 2 {
		return ErrServiceNotEnoughParameter{Name: photos.name, Param: args}
	}
	switch message := args[0].(type) {
	case string:
		oper = message
		logs.DebugLog(oper)
	default:
		return ErrServiceNotCorrectParamType{Name: photos.name, Param: message, Number: 1}
	}

	if oper == "save" {
		switch message := args[1].(type) {
		case string:
			photos.fileName = message
		default:
			return ErrServiceNotCorrectParamType{Name: photos.name, Param: message, Number: 2}
		}
		if len(args) < 3 {
			return ErrServiceNotEnoughParameter{Name: photos.name, Param: args}
		}
		return photos.saveFile(args[2].(io.Reader))
	} else if oper == "read" {
		log.Println(oper)
		return ErrServiceNotCorrectOperation{Name: photos.name, OperName: oper}

	} else {
		return ErrServiceNotCorrectOperation{Name: photos.name, OperName: oper}

	}

	return nil

}

//возвращает интерфейс чтения файла картинки
//согласно параметрам:
//1 - имя таблицы, view или сервиса
//2 - id записи
//3 - порядковый номер файла файла
// (-1 означает, что нужно вернуть массив со списком файлов)
func (photos *photosService) Get(args ...interface{}) (responce interface{}, err error) {

	var catalog, id string
	var num int
	if len(args) < 3 {
		return nil, ErrServiceNotEnoughParameter{Name: photos.name, Param: args}
	}
	switch message := args[0].(type) {
	case string:
		catalog = message
	default:
		return nil, ErrServiceNotCorrectParamType{Name: photos.name, Param: message, Number: 1}
	}

	switch message := args[1].(type) {
	case string:
		id = message
	default:
		return nil, ErrServiceNotCorrectParamType{Name: photos.name, Param: message, Number: 2}
	}

	switch message := args[2].(type) {
	case int:
		num = message
	case int32, int64:
		num = message.(int)
	default:
		log.Println(message)
		return nil, ErrServiceNotCorrectParamType{Name: photos.name, Param: message, Number: 3}
	}

	if num == -1 {

		return photos.listLinks(catalog, id)
	} else {

		return photos.readFile(catalog, id, num)
	}
}
func (photos *photosService) Connect(in <-chan interface{}) (out chan interface{}, err error) {

	return nil, nil
}
func (photos *photosService) Close(out chan<- interface{}) error {

	return nil
}
func (photos *photosService) Status() string {

	return photos.status
}
func (photos *photosService) saveFile(inFile io.Reader) error {

	fullName := filepath.Join(photos.path, photos.fileName)
	outFile, err := os.Create(fullName)
	if err != nil {
		if os.IsNotExist(err) {
			dir := filepath.Dir(fullName)
			//TODO: add privilegias from config file!
			err = os.MkdirAll(dir, os.ModeDir | os.ModePerm)
			if err == nil {
				outFile, err = os.Create(fullName)
			}
		}
	}
	// записываем данные файла из потока
	if err == nil {
		defer outFile.Close()
		_, err = io.Copy(outFile, inFile)
	}

	if err != nil {
		logs.ErrorLog(err, "Error saving file: " + fullName)
		return err
	}

	logs.DebugLog("Succesfull save file " + fullName)
	return nil

}
func (photos *photosService) readFile(catalog, id string, num int) (io.Reader, error) {

	files, err := photos.listFiles(catalog, id)
	if err != nil {
		logs.ErrorLog(err.(error))
		return nil, err
	}

	if num > len(files)-1 {
		return nil, ErrServiceWrongIndex{Name: "Number not in files array range", Index: num}
	}
	outFile, err := os.Open(files[num])
	if err != nil {
		logs.ErrorLog(err.(error))
		return nil, err
	}

	return outFile, nil

}
func (photos *photosService) listFiles(catalog, id string) (files []string, err error) {

	fullPath := filepath.Join(photos.path, catalog, id)
	return filepath.Glob(fullPath + "/*.*")

}
func (photos *photosService) listLinks(catalog, id string) (list []string, err error) {
	list, err = photos.listFiles(catalog, id)
	if err != nil {
		return nil, err
	}

	for num, _ := range list {
		list[num] = fmt.Sprintf("/api/v1/photos/?table=%s&id=%s&num=%d", catalog, id, num)
	}

	return list, nil
}
func init() {
	AddService(photos.name, photos)
}
