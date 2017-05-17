// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// this module has more structures from creating sql-query with relation to db schema

package qb

type QBFields struct {
	Name  string
	Alias string

}
type QBTables struct {
	Name string
	Alias string
	Join string
	Using string
	Fields map[string] *QBFields
}
type QueryBuilder struct {
	Tables [] *QBTables
	Where   string
	Args [] interface{}
	GroupBy string
	OrderBy string
}
// constructor
func  Create(where, groupBy, orderBy string) *QueryBuilder{

	qb := &QueryBuilder{Where: where, OrderBy: orderBy, GroupBy: groupBy}
	return qb
}
// addding arguments
func (qb *QueryBuilder) AddArgs(arg interface{}) *QueryBuilder{
	qb.Args = append(qb.Args, arg)

	return qb
}
// add Tables list, returns qb
func (qb *QueryBuilder) AddTables(names map[string] string) *QueryBuilder {
	for alias, name := range names {
		qb.AddTable(alias, name)
	}

	return qb
}
//add Table, returns object table
func (qb *QueryBuilder) AddTable(alias, name string) *QBTables {

	table := &QBTables{Name: name}
	table.Fields = make(map[string] *QBFields, 0)
	qb.Tables    = append(qb.Tables, table)

	return table
}
// add table with join
func (qb *QueryBuilder) JoinTable(alias, name, join, usingOrOn string) *QBTables {

	table := &QBTables{Name: name}
	table.Fields = make(map[string] *QBFields, 0)
	table.Join   = join
	table.Using  = usingOrOn
	qb.Tables    = append(qb.Tables, table)

	return table
}
// add fields to table from map
func (table *QBTables) AddFields(fields map[string] string) *QBTables {
	for alias, name := range fields {
		table.AddField(alias, name)
	}

	return table
}
// add field and returns table object
func (table *QBTables) AddField(alias, name string) *QBTables {

	field := &QBFields{Name: name}
	table.Fields[alias] = field

	return table
}

