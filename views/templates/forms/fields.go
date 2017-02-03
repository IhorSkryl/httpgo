package forms

import (
	"strings"
	"github.com/ruslanBik4/httpgo/models/db"
	"regexp"
	"fmt"
	"log"
	"database/sql"
)
var (
	enumValidator = regexp.MustCompile(`(?:'([^,]+)',?)`)

)
type FieldStructure struct {
	COLUMN_NAME   	string
	DATA_TYPE 	string
	COLUMN_DEFAULT 	string
	IS_NULLABLE 	string
	CHARACTER_SET_NAME string
	COLUMN_COMMENT 	string
	COLUMN_TYPE 	string
	CHARACTER_MAXIMUM_LENGTH int
	Value 		string
	IsHidden 	bool
	CSSClass  	string
	TableName 	string
	Events 		map[string] string
	Figure 		string
	Html		string
}
type FieldsTable struct {
	Name string
	ID   int
	IsDadata bool
	Rows [] FieldStructure
	Hiddens map[string] string
	SaveFormEvents 	map[string] string
}

func getFields(tableName string) (fields FieldsTable, err error) {
	var ns db.FieldsTable

	ns.GetColumnsProp(tableName)


	fields.PutDataFrom( ns )
	fields.Name = tableName

	return fields, nil
}

// Scan implements the Scanner interface.
func (field *FieldStructure) Scan(value interface{}) error {
	var temp sql.NullString

	if err := temp.Scan(value); err != nil {
		field.Value = ""
		return err
	}
	field.Value = temp.String

	return nil
}

func (ns *FieldsTable) FindField(name string) *FieldStructure {
	for idx, field := range ns.Rows {
		if field.COLUMN_NAME == name {
			return &ns.Rows[idx]
		}
	}

	return nil
}
func (field *FieldStructure) whereFromSet(ns *FieldsTable) (result string) {
	fields := enumValidator.FindAllStringSubmatch(field.COLUMN_TYPE, -1)
	comma  := " WHERE "
	for _, title := range fields {
		enumVal := title[len(title) - 1]
		if i := strings.Index(enumVal, ":"); i > 0 {
			param := ""
			if paramField := ns.FindField(enumVal[i+1:]); paramField != nil {
				param = paramField.Value
				enumVal = enumVal[:i] + fmt.Sprintf("%s", param)
			} else {
				continue
			}
		}
		result += comma + enumVal
		comma = " OR "
	}

	return result
}
//получаем связанную таблицу с полями

func (field *FieldStructure) getTableFrom(ns *FieldsTable) {
	key := field.COLUMN_NAME
	tableProps := key[ len("tableid_") : ]

	fields, err := getFields(tableProps)
	if err != nil {
		field.Html += "<td>Error during readin table schema!" + tableProps + "</td>"
		return
	}

	where := field.whereFromSet(ns)
	if where > "" {
		where += " AND (id_%s=?)"
	} else {
		where = " WHERE (id_%s=?)"
	}
	sqlCommand := fmt.Sprintf( `SELECT * FROM %s p ` + where, tableProps, ns.Name )

	rows, err := db.DoSelect( sqlCommand, ns.ID )
	if err != nil {
		log.Println(sqlCommand, err)
		return
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if (err != nil) {
		log.Println(err)
	}

	field.Html = "<thead> <tr>"

	var row [] interface {}
	var newRow string

	rowField := make([] *sql.NullString, len(columns))
	for idx, fieldName := range columns {

		if (fieldName == "id") || (fieldName == "id_" + ns.Name ) {
			//newRow += "<td></td>"
		} else {
			field.Html += "<td>" + fieldName + "</td>"
			newRow += getTD(tableProps, fieldName, "", "id_" + ns.Name, 0, fields.FindField(fieldName) )
		}

		rowField[idx] = new(sql.NullString)
		row = append( row, rowField[idx] )
	}
	field.Html += "</tr></thead><tbody>"

	idx := 0

	for rows.Next() {
		if err := rows.Scan(row...); err != nil {
			log.Println(err)
			continue
		}
		idx++
		field.Html += "<tr>"
		for i, value := range rowField {
			field.Html += getTD(tableProps, columns[i], value.String, "id_" + ns.Name, idx, fields.FindField(columns[i]) )
		}

		field.Html += "</tr>"

	}
	field.Html += "<tr>" + newRow + "</tr>" + "</tbody>"
}
const CELL_TABLE  = `<td><input type="%s" name="%s:%s" value="%s"/></td>`
const CELL_SELECT  = `<td><select name="%s:%s" class="controls">%s</select></td>`
func getTD(tableProps, fieldName, value, parentField string, idx int, fieldStruct *FieldStructure) (html string){
	inputName := fieldName + fmt.Sprintf("[%d]", idx)

	if (fieldName == "id") || (fieldName == parentField) {
		if value > "" {
			html += fmt.Sprintf(CELL_TABLE, "hidden", tableProps, inputName, value)
		}
	} else if strings.HasPrefix(fieldName, "id_") {
		fieldStruct.getOptions(fieldName[3:], value)
		html += fmt.Sprintf(CELL_SELECT, tableProps, inputName, fieldStruct.Html )
	} else {
		html += fmt.Sprintf(CELL_TABLE, "text", tableProps, inputName, value)
	}

	return html
}

func (field *FieldStructure) getMultiSelect(ns *FieldsTable){

	key := field.COLUMN_NAME
	tableProps := strings.TrimLeft(key, "setid_")
	tableValue := ns.Name + "_" + tableProps + "_has"

	titleField := db.GetParentFieldName(tableProps)
	if titleField == "" {
		return
	}

	where := field.whereFromSet(ns)
	sqlCommand := fmt.Sprintf( `SELECT p.id, %s, id_%s
	FROM %s p LEFT JOIN %s v ON (p.id=v.id_%[3]s AND id_%[2]s=?) %[5]s`,
		titleField, ns.Name,
		tableProps, tableValue, where)


	rows, err := db.DoSelect( sqlCommand, ns.ID )
	if err != nil {
		log.Println(sqlCommand, err)
		return
	}
	defer rows.Close()
	idx := 0

	field.Html = ""
	for rows.Next() {
		var id string
		var title, checked string
		var idRooms sql.NullInt64

		if err := rows.Scan(&id, &title, &idRooms); err != nil {
				log.Println(err)
				continue
		}
		if idRooms.Valid {
			checked = "checked"
		}
		idx++


		field.Html += "<li role='presentation'>" + renderCheckBox(key + "[]", id, title, idx, checked, "", "") + "</li>"
	}

}
func (field *FieldStructure) getOptions(tableName, val string) {

	name := db.GetParentFieldName(tableName)
	if name == "" {
		return
	}
	rows, err := db.DoSelect("select id, " + name + " from " + tableName )
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	idx := 0
	//valueID, _ := strconv.Atoi(val)

	field.Html = ""

	for rows.Next() {

		var id, title, selected string

		if err := rows.Scan(&id, &title); err != nil {
			log.Println(err)
			continue
		}
		if val == id {
			selected = "selected"
		}
		idx++

		field.Html += renderOption(id, title, selected)
	}
}

func (field *FieldStructure) renderSet(key, val, required, events, dataJson string) (result string) {
	fields := enumValidator.FindAllStringSubmatch(field.COLUMN_TYPE, -1)

	for idx, title := range fields {
		enumVal := title[len(title)-1]
		checked := ""
		if strings.Contains(val, enumVal) {
			checked = "checked"
		}
		result += renderCheckBox(key + "[]", enumVal, enumVal, idx, checked, events, dataJson)
	}

	return result
}
func (field *FieldStructure) renderEnum(key, val, required, events, dataJson string) (result string) {


	fields := enumValidator.FindAllStringSubmatch(field.COLUMN_TYPE, -1)
	isRenderSelect := len(fields) > 2

	for idx, title := range fields {
		enumVal := title[len(title)-1]
		checked, selected  := "", ""
		if val == enumVal {
			checked, selected = "checked", "selected"
		}
		if isRenderSelect {
			result += renderOption(enumVal, enumVal, selected)
		} else {
			result += renderRadioBox(key, enumVal, enumVal, idx, checked, events, dataJson)
		}
	}
	if isRenderSelect {
		return renderSelect(key, result, required, events, dataJson )
	}

	return result
}

func cutPartFromTitle(title, pattern, defaultStr string) (titleFull, titlePart string)  {
	titleFull = title
	if title == "" {
		return "", ""
	}
	posPattern := strings.Index(titleFull, pattern)
	if posPattern > 0 {
		titlePart = titleFull[posPattern + len(pattern):]
		titleFull = titleFull[:posPattern]
	} else {
		titlePart = defaultStr
	}

	return titleFull, titlePart
}
func (field *FieldStructure) GetColumnTitles() (titleFull, titleLabel, placeholder, pattern, dataJson string)  {
	titleFull = field.COLUMN_COMMENT
	if titleFull == "" {
		titleLabel = field.COLUMN_NAME
		return titleFull, titleLabel, placeholder, pattern, dataJson
	} else if strings.Index(titleFull, ".") > 0 {
		titleLabel = titleFull[:strings.Index(titleFull, ".")]
	} else {
		titleLabel = titleFull
	}
	titleFull, pattern = cutPartFromTitle(titleFull, "//", "")
	titleFull, dataJson = cutPartFromTitle(titleFull, "{", "")
	titleFull, placeholder = cutPartFromTitle(titleFull, "#", titleFull)

	return titleFull, titleLabel, placeholder, pattern, dataJson
}
// заполняет структуру для формы данными, взятыми из структуры БД
func (fields *FieldsTable) PutDataFrom(ns db.FieldsTable) {

	for _, field := range ns.Rows {
		fieldStrc := &FieldStructure{
			COLUMN_NAME: field.COLUMN_NAME,
			DATA_TYPE  : field.DATA_TYPE,
			IS_NULLABLE: field.IS_NULLABLE,
			COLUMN_TYPE: field.COLUMN_TYPE,
			Events     : make(map[string] string, 0),
		}
		if field.CHARACTER_SET_NAME.Valid {
			fieldStrc.CHARACTER_SET_NAME = field.CHARACTER_SET_NAME.String
		}
		if field.COLUMN_COMMENT.Valid {
			fieldStrc.COLUMN_COMMENT = field.COLUMN_COMMENT.String
		}
		if field.CHARACTER_MAXIMUM_LENGTH.Valid {
			fieldStrc.CHARACTER_MAXIMUM_LENGTH = int(field.CHARACTER_MAXIMUM_LENGTH.Int64)
		}
		if field.COLUMN_DEFAULT.Valid {
			fieldStrc.COLUMN_DEFAULT = field.COLUMN_DEFAULT.String
		}
		fieldStrc.IsHidden = false
		fields.Rows = append(fields.Rows,*fieldStrc)
	}
}

