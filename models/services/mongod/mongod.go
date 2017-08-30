// Copyright 2017 Author: Yurii Kravchuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Реализует работу с базой данных mongodb
package mongod

import (
	"github.com/ruslanBik4/httpgo/models/server"
	"github.com/ruslanBik4/httpgo/models/services"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongod *mdService = &mdService{name: "mongod"}
)

type mdService struct {
	name    string
	connect *mongo.Session
	status  string
}

//Запускает связь с mongodb на дефолтный порт. (localhost:27017)
//TODO: перенести порт для связи в config для его настройки
func (mongod *mdService) Init() error {

	session, err := mongo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mongo.Monotonic, true)

	mongod.connect = session
	mongod.status = "ready"
	return nil
}

func (mongod *mdService) Connect(in <-chan interface{}) (out chan interface{}, err error) {
	out = make(chan interface{})

	go func() {
		out <- "open"
		for {
			select {
			case v := <-in:
				if v.(string) == "close" {
					mongod.Close(out)
				} else {
					out <- v
				}
			}
		}
	}()
	return out, nil
}

func (mongod *mdService) Close(out chan<- interface{}) error {
	close(out)
	return nil
}

//получаем текущий статус сервиса
func (mongod *mdService) Status() string {
	return mongod.status
}

func (mongod *mdService) Get(args ...interface{}) (interface{}, error) {

	if len(args) < 2 {
		return nil, ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}

	connection_status := Status("mongod")

	if connection_status != "ready" {
		return nil, ErrBrokenConnection{Name: mongod.name, Param: args}
	}

	var collection string
	switch col := args[0].(type) {
	case string:
		collection = col
	default:
		return nil, ErrServiceNotCorrectParamType{Name: mongod.name, Param: col, Number: 1}
	}

	connetion := mongod.connect
	cConnect := connetion.DB(server.GetMongodConfig().MongoDBName()).C(collection)

	switch option := args[1].(type) {
	case string:
		switch option {
		case "Find":
			return findRecord(cConnect, args)

		default:
			return nil, ErrServiceNotCorrectParamType{Name: mongod.name, Param: args, Number: 2}
		}
	}

	return nil, ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
}

func (mongod *mdService) Send(args ...interface{}) error {

	if len(args) < 2 {
		return ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}

	connection_status := Status("mongod")

	if connection_status != "ready" {
		return ErrBrokenConnection{Name: mongod.name, Param: args}
	}

	var collection string
	switch col := args[0].(type) {
	case string:
		collection = col
	default:
		return ErrServiceNotCorrectParamType{Name: mongod.name, Param: col, Number: 1}
	}

	connetion := mongod.connect
	cConnect := connetion.DB(server.GetMongodConfig().MongoDBName()).C(collection)

	switch option := args[1].(type) {
	case string:
		switch option {
		case "Insert":
			return insertRecord(cConnect, args)

		case "Update":
			return updateRecord(cConnect, args)

		case "Remove":
			return removeRecord(cConnect, args)

		default:
			return ErrServiceNotCorrectParamType{Name: mongod.name, Param: args, Number: 2}
		}
	}

	return ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
}

//поиск записей в mongodb
func findRecord(cConnect *mongo.Collection, args []interface{}) (interface{}, error) {

	if len(args) < 4 {
		return nil, ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}
	switch oType := args[2].(type) {
	case string:
		switch oType {
		case "All":
			switch args[3].(type) {
			case bson.M:
				responce := make([]interface{}, 0)
				err := cConnect.Find(args[3]).All(&responce)

				if err != nil {
					return nil, err
				}

				return responce, nil

			default:
				return nil, ErrServiceNotCorrectParamType{Name: mongod.name, Param: args, Number: 4}
			}
		case "One":
			switch args[3].(type) {
			case bson.M:
				var responce interface{}
				err := cConnect.Find(args[3]).One(&responce)

				if err != nil {
					return nil, err
				}

				return responce, nil

			default:
				return nil, ErrServiceNotCorrectParamType{Name: mongod.name, Param: args, Number: 4}
			}
		default:
			return nil, ErrServiceNotCorrectParamType{Name: mongod.name, Param: args, Number: 3}
		}
	}

	return nil, nil
}

//создание новой записи в mongodb
func insertRecord(cConnect *mongo.Collection, args []interface{}) error {

	if len(args) < 3 {
		return ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}
	err := cConnect.Insert(args[2])
	if err != nil {
		return err
	}

	return nil
}

//обновление записи в mongodb
func updateRecord(cConnect *mongo.Collection, args []interface{}) error {

	if len(args) < 4 {
		return ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}
	_, err := cConnect.UpdateAll(args[2], args[3])
	if err != nil {
		return err
	}

	return nil
}

//удаление записи в mongodb
func removeRecord(cConnect *mongo.Collection, args []interface{}) error {

	if len(args) < 3 {
		return ErrServiceNotEnoughParameter{Name: mongod.name, Param: args}
	}
	err := cConnect.Remove(args[2])

	if err != nil {
		return err
	}

	return nil
}

func GetMongoCollectionConnect(collection string) *mongo.Collection {
	return mongod.connect.DB(server.GetMongodConfig().MongoDBName()).C(collection)
}

func init() {
	services.AddService(mongod.name, mongod)
}
