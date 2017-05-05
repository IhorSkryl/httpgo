// Copyright 2017
// 	Author's: Mykhailo Sizov sizov.mykhailo@gmail.com
// All rights reserved.
// version 1.0
// Базовый функционал для работы с транзакциями.
// Важно : откат изменений происходит только в рамках данной транзакции!
package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"encoding/json"
	"bytes"
	"io"
	"log"
	"strings"
)

type txConnect struct {
	tx *sql.Tx
}

func StartTransaction() (*txConnect, error) {
	tx, err := dbConn.Begin()
	if err != nil {
		return nil, err
	}
	return &txConnect{tx:tx}, nil
}
func (conn *txConnect) CommitTransaction() {
	conn.tx.Commit()
}
func (conn *txConnect) RollbackTransaction() {
	conn.tx.Rollback()
}
func (conn *txConnect) prepareQuery(sql string) (*sql.Stmt, error) {
	return conn.tx.Prepare(sql)
}
func (conn *txConnect) DoInsert(sql string, args ...interface{}) (int, error) {

	resultSQL, err := conn.tx.Exec(sql, args ...)

	if err != nil {
		log.Println(sql)
		return -1, err
	} else {
		lastInsertId, err := resultSQL.LastInsertId()
		return int(lastInsertId), err
	}
}

func (conn *txConnect) DoUpdate(sql string, args ...interface{}) (int, error) {
	resultSQL, err := conn.tx.Exec(sql, args ...)

	if err != nil {
		return -1, err
	} else {
		RowsAffected, err := resultSQL.RowsAffected()
		return int(RowsAffected), err
	}
}
func (conn *txConnect) DoSelect(sql string, args ...interface{}) (*sql.Rows, error) {
	if SQLvalidator.MatchString(strings.ToLower(sql)) {
		return conn.tx.Query(sql, args ...)
	} else {
		return nil, mysql.ErrMalformPkt
	}
}
func (conn *txConnect) DoQuery(sql string, args ...interface{}) *sql.Rows {
	var result bytes.Buffer

	w := io.Writer(&result)
	Encode := json.NewEncoder(w)

	if strings.HasPrefix(sql, "insert") {
		resultSQL, err := conn.tx.Exec(sql, args ...)

		if err != nil {
			Encode.Encode(err)
		} else {
			lastInsertId, _ := resultSQL.LastInsertId()
			Encode.Encode(lastInsertId)
		}

		log.Print(result.String())

		return nil
	}

	if strings.HasPrefix(sql, "update") {
		resultSQL, err := conn.tx.Exec(sql, args ...)

		if err != nil {
			Encode.Encode(err)
		} else {
			RowsAffected, _ := resultSQL.RowsAffected()
			Encode.Encode(RowsAffected)
		}

		log.Print(result.String())

		return nil
	}

	rows, err := conn.tx.Query(sql, args ...)

	if err != nil {
		log.Print(result.String())
		return nil
	}

	return rows
}