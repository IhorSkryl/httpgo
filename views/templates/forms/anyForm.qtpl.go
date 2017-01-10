// This file is automatically generated by qtc from "anyForm.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/forms/anyForm.qtpl:1
package forms

//line views/templates/forms/anyForm.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
//

//line views/templates/forms/anyForm.qtpl:4
import (
	"fmt"
	"github.com/ruslanBik4/httpgo/models/db"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//line views/templates/forms/anyForm.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/forms/anyForm.qtpl:13
type FieldStructure struct {
	COLUMN_NAME              string
	DATA_TYPE                string
	COLUMN_DEFAULT           string
	IS_NULLABLE              string
	CHARACTER_SET_NAME       string
	COLUMN_COMMENT           string
	COLUMN_TYPE              string
	CHARACTER_MAXIMUM_LENGTH int
	Value                    string
	IsHidden                 bool
	CSSClass                 string
	TableName                string
	Events                   map[string]string
}
type FieldsTable struct {
	Name     string
	IsDadata bool
	Rows     []FieldStructure
	Hiddens  map[string]string
}

//line views/templates/forms/anyForm.qtpl:35
func (field *FieldStructure) streamrenderParentSelect(qw422016 *qt422016.Writer, nameTable, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:35
	qw422016.N().S(`
    <label for="`)
	//line views/templates/forms/anyForm.qtpl:36
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:36
	qw422016.N().S(`">`)
	//line views/templates/forms/anyForm.qtpl:36
	qw422016.E().S(titleLabel)
	//line views/templates/forms/anyForm.qtpl:36
	qw422016.N().S(`</label>
    <select id="`)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.N().S(`" class="form-control" `)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.E().S(required)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.N().S(` `)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.E().S(val)
	//line views/templates/forms/anyForm.qtpl:37
	qw422016.N().S(`>
        `)
	//line views/templates/forms/anyForm.qtpl:38
	if field.IS_NULLABLE == "YES" && val == "" {
		//line views/templates/forms/anyForm.qtpl:38
		qw422016.N().S(`
        <option value="" selected>Значение можно не указывать</option>
        `)
		//line views/templates/forms/anyForm.qtpl:40
	} else {
		//line views/templates/forms/anyForm.qtpl:40
		qw422016.N().S(`
            <option disabled >Выберите значение из списка</option>
        `)
		//line views/templates/forms/anyForm.qtpl:42
	}
	//line views/templates/forms/anyForm.qtpl:42
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:44
	if nameTable == "" {
		nameTable = field.TableName
	}

	//line views/templates/forms/anyForm.qtpl:47
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:48
	streamgetOptions(qw422016, "id_"+nameTable, field.Value)
	//line views/templates/forms/anyForm.qtpl:48
	qw422016.N().S(`
    </select>
`)
//line views/templates/forms/anyForm.qtpl:50
}

//line views/templates/forms/anyForm.qtpl:50
func (field *FieldStructure) writerenderParentSelect(qq422016 qtio422016.Writer, nameTable, key, val, titleLabel, required string) {
	//line views/templates/forms/anyForm.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:50
	field.streamrenderParentSelect(qw422016, nameTable, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:50
}

//line views/templates/forms/anyForm.qtpl:50
func (field *FieldStructure) renderParentSelect(nameTable, key, val, titleLabel, required string) string {
	//line views/templates/forms/anyForm.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:50
	field.writerenderParentSelect(qb422016, nameTable, key, val, titleLabel, required)
	//line views/templates/forms/anyForm.qtpl:50
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:50
	return qs422016
//line views/templates/forms/anyForm.qtpl:50
}

//line views/templates/forms/anyForm.qtpl:51
func streamgetOptions(qw422016 *qt422016.Writer, key string, value string) {
	//line views/templates/forms/anyForm.qtpl:51
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:53
	var listNs db.FieldsTable

	tableName := key[3:]
	err := listNs.GetColumnsProp(tableName)

	//line views/templates/forms/anyForm.qtpl:57
	qw422016.N().S(`
        `)
	//line views/templates/forms/anyForm.qtpl:58
	if err == nil {
		//line views/templates/forms/anyForm.qtpl:58
		qw422016.N().S(`
        `)
		//line views/templates/forms/anyForm.qtpl:61
		name := ""
		for _, list := range listNs.Rows {
			switch list.COLUMN_NAME {
			case "name":
				name = "name"
				break
			case "title":
				name = "title"
				break
			case "fullname":
				name = "fullname"
				break
			}
		}

		//line views/templates/forms/anyForm.qtpl:75
		qw422016.N().S(`
            `)
		//line views/templates/forms/anyForm.qtpl:76
		if name > "" {
			//line views/templates/forms/anyForm.qtpl:76
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:78
			rows := db.DoQuery("select id, " + name + " from " + tableName)
			defer rows.Close()
			selected := ""
			valueID, _ := strconv.Atoi(value)

			//line views/templates/forms/anyForm.qtpl:82
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:83
			for rows.Next() {
				//line views/templates/forms/anyForm.qtpl:83
				qw422016.N().S(`
                `)
				//line views/templates/forms/anyForm.qtpl:85
				var id int
				var name string

				if err := rows.Scan(&id, &name); err != nil {
					log.Println(err)
					continue
				}
				if valueID == id {
					selected = "selected"
				} else {
					selected = ""
				}

				//line views/templates/forms/anyForm.qtpl:97
				qw422016.N().S(`
                    <option value='`)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.N().D(id)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.N().S(`' `)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.E().S(selected)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.N().S(`>`)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.E().S(name)
				//line views/templates/forms/anyForm.qtpl:98
				qw422016.N().S(`</option>
            `)
				//line views/templates/forms/anyForm.qtpl:99
			}
			//line views/templates/forms/anyForm.qtpl:99
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:100
		}
		//line views/templates/forms/anyForm.qtpl:100
		qw422016.N().S(`
    `)
		//line views/templates/forms/anyForm.qtpl:101
	} else {
		//line views/templates/forms/anyForm.qtpl:101
		qw422016.N().S(`
        <option value='-1' >No found</option>
    `)
		//line views/templates/forms/anyForm.qtpl:103
	}
	//line views/templates/forms/anyForm.qtpl:103
	qw422016.N().S(`
`)
//line views/templates/forms/anyForm.qtpl:104
}

//line views/templates/forms/anyForm.qtpl:104
func writegetOptions(qq422016 qtio422016.Writer, key string, value string) {
	//line views/templates/forms/anyForm.qtpl:104
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:104
	streamgetOptions(qw422016, key, value)
	//line views/templates/forms/anyForm.qtpl:104
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:104
}

//line views/templates/forms/anyForm.qtpl:104
func getOptions(key string, value string) string {
	//line views/templates/forms/anyForm.qtpl:104
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:104
	writegetOptions(qb422016, key, value)
	//line views/templates/forms/anyForm.qtpl:104
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:104
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:104
	return qs422016
//line views/templates/forms/anyForm.qtpl:104
}

//line views/templates/forms/anyForm.qtpl:106
func cutPartFromTitle(pattern, defaultStr string) (titleFull, titlePart string) {

	posPattern := strings.Index(titleFull, pattern)
	if posPattern > 0 {
		titlePart = titleFull[posPattern+len(pattern):]
		titleFull = titleFull[:posPattern]
	} else {
		titlePart = defaultStr
	}

	return titleFull, titlePart
}
func (field *FieldStructure) GetColumnTitles() (titleFull, titleLabel, placeholder, pattern, dataJson string) {
	titleFull = field.COLUMN_COMMENT
	if titleFull == "" {
		titleLabel = field.COLUMN_NAME
	} else if strings.Index(titleFull, ".") > 0 {
		titleLabel = titleFull[:strings.Index(titleFull, ".")]
	} else {
		titleLabel = titleFull
	}
	titleFull, pattern = cutPartFromTitle("//", "")
	titleFull, placeholder = cutPartFromTitle("#", titleFull)
	titleFull, dataJson = cutPartFromTitle("{", "")

	return titleFull, titleLabel, placeholder, pattern, dataJson
}

//line views/templates/forms/anyForm.qtpl:134
func (ns *FieldsTable) StreamShowAnyForm(qw422016 *qt422016.Writer, Action, Title string) {
	//line views/templates/forms/anyForm.qtpl:134
	qw422016.N().S(`
<form class="row-fluid" target="content" action="`)
	//line views/templates/forms/anyForm.qtpl:135
	qw422016.E().S(Action)
	//line views/templates/forms/anyForm.qtpl:135
	qw422016.N().S(`" method="post" class="form-horizontal" onsubmit="return saveForm(this, afterSaveAnyForm);" >
    <h2 class="form-signin-heading">`)
	//line views/templates/forms/anyForm.qtpl:136
	qw422016.E().S(Title)
	//line views/templates/forms/anyForm.qtpl:136
	qw422016.N().S(`</h2>
    `)
	//line views/templates/forms/anyForm.qtpl:137
	if ns.Name > "" {
		//line views/templates/forms/anyForm.qtpl:137
		qw422016.N().S(`
        <input hidden name="table" value="`)
		//line views/templates/forms/anyForm.qtpl:138
		qw422016.E().S(ns.Name)
		//line views/templates/forms/anyForm.qtpl:138
		qw422016.N().S(`" >
    `)
		//line views/templates/forms/anyForm.qtpl:139
	}
	//line views/templates/forms/anyForm.qtpl:139
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:140
	for idx, field := range ns.Rows {
		//line views/templates/forms/anyForm.qtpl:140
		qw422016.N().S(`

        `)
		//line views/templates/forms/anyForm.qtpl:143
		key := field.COLUMN_NAME
		titleFull, titleLabel, placeholder, pattern, dataJson := field.GetColumnTitles()
		val := field.Value
		tableName := ""
		required := ""

		if field.IS_NULLABLE == "NO" {
			required = "required"
		}

		if val > "" {
			val = "value=" + val + ""
		} else if field.COLUMN_DEFAULT > "" {
			val = "value=" + field.COLUMN_DEFAULT + ""
		}
		if field.TableName > "" {
			tableName = field.TableName + ":"
		}
		events := ""
		for name, funcName := range field.Events {
			events += fmt.Sprintf("%s=\"return %s;\"", name, funcName)
		}

		//line views/templates/forms/anyForm.qtpl:165
		qw422016.N().S(`

        `)
		//line views/templates/forms/anyForm.qtpl:167
		if (val > "") && ((key == "id") || field.IsHidden) {
			//line views/templates/forms/anyForm.qtpl:167
			qw422016.N().S(`
            <input hidden name="`)
			//line views/templates/forms/anyForm.qtpl:168
			qw422016.E().S(key)
			//line views/templates/forms/anyForm.qtpl:168
			qw422016.N().S(`" `)
			//line views/templates/forms/anyForm.qtpl:168
			qw422016.E().S(val)
			//line views/templates/forms/anyForm.qtpl:168
			qw422016.N().S(`>
            `)
			//line views/templates/forms/anyForm.qtpl:169
			continue
			//line views/templates/forms/anyForm.qtpl:170
		} else if key == "id" {
			//line views/templates/forms/anyForm.qtpl:170
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:171
			continue
			//line views/templates/forms/anyForm.qtpl:172
		}
		//line views/templates/forms/anyForm.qtpl:172
		qw422016.N().S(`

        <div id="divField`)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.N().D(idx)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.N().S(`" class="form-group field-`)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.E().S(key)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.N().S(` `)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.E().S(required)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.N().S(` `)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.E().S(field.CSSClass)
		//line views/templates/forms/anyForm.qtpl:174
		qw422016.N().S(`"
                    `)
		//line views/templates/forms/anyForm.qtpl:175
		if field.IsHidden {
			//line views/templates/forms/anyForm.qtpl:175
			qw422016.N().S(`style="display:none"`)
			//line views/templates/forms/anyForm.qtpl:175
		}
		//line views/templates/forms/anyForm.qtpl:175
		qw422016.N().S(`
             data-toggle="tooltip" title="`)
		//line views/templates/forms/anyForm.qtpl:176
		qw422016.N().S(titleFull)
		//line views/templates/forms/anyForm.qtpl:176
		qw422016.N().S(`"
        >
        `)
		//line views/templates/forms/anyForm.qtpl:178
		if key == "parent_id" {
			//line views/templates/forms/anyForm.qtpl:178
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:179
			field.streamrenderParentSelect(qw422016, ns.Name, key, val, titleLabel, required)
			//line views/templates/forms/anyForm.qtpl:179
			qw422016.N().S(`
        `)
			//line views/templates/forms/anyForm.qtpl:180
		} else if strings.HasPrefix(key, "id_") {
			//line views/templates/forms/anyForm.qtpl:180
			qw422016.N().S(`
            <label for="`)
			//line views/templates/forms/anyForm.qtpl:181
			qw422016.E().S(key)
			//line views/templates/forms/anyForm.qtpl:181
			qw422016.N().S(`">`)
			//line views/templates/forms/anyForm.qtpl:181
			qw422016.E().S(titleLabel)
			//line views/templates/forms/anyForm.qtpl:181
			qw422016.N().S(`</label>
            <select id="`)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.E().S(key)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.N().S(`" name="`)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.E().S(tableName + key)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.N().S(`" class="form-control" `)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.E().S(required)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.N().S(` `)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.E().S(val)
			//line views/templates/forms/anyForm.qtpl:182
			qw422016.N().S(`>
                `)
			//line views/templates/forms/anyForm.qtpl:183
			if field.IS_NULLABLE == "YES" && val == "" {
				//line views/templates/forms/anyForm.qtpl:183
				qw422016.N().S(`
                    <option value="" selected>Значение можно не указывать</option>
                `)
				//line views/templates/forms/anyForm.qtpl:185
			} else {
				//line views/templates/forms/anyForm.qtpl:185
				qw422016.N().S(`
                    <option disabled >Выберите значение из списка</option>
                `)
				//line views/templates/forms/anyForm.qtpl:187
			}
			//line views/templates/forms/anyForm.qtpl:187
			qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:188
			streamgetOptions(qw422016, key, field.Value)
			//line views/templates/forms/anyForm.qtpl:188
			qw422016.N().S(`
            </select>
        `)
			//line views/templates/forms/anyForm.qtpl:190
		} else {
			//line views/templates/forms/anyForm.qtpl:190
			qw422016.N().S(`

                `)
			//line views/templates/forms/anyForm.qtpl:192
			switch field.DATA_TYPE {
			//line views/templates/forms/anyForm.qtpl:193
			case "tinyint":
				//line views/templates/forms/anyForm.qtpl:193
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:195
				checked := ""
				if field.Value > "" {
					checked = "checked"
				}

				//line views/templates/forms/anyForm.qtpl:199
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:200
				streamrenderCheckBox(qw422016, tableName+key, titleLabel, 1, checked)
				//line views/templates/forms/anyForm.qtpl:200
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:201
			case "enum":
				//line views/templates/forms/anyForm.qtpl:201
				qw422016.N().S(`
                    <label for="`)
				//line views/templates/forms/anyForm.qtpl:202
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:202
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:202
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:202
				qw422016.N().S(`:</label>
                    `)
				//line views/templates/forms/anyForm.qtpl:203
				field.streamrenderEnum(qw422016, tableName+key, required)
				//line views/templates/forms/anyForm.qtpl:203
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:204
			case "set":
				//line views/templates/forms/anyForm.qtpl:204
				qw422016.N().S(`
                    <label for="`)
				//line views/templates/forms/anyForm.qtpl:205
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:205
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:205
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:205
				qw422016.N().S(`:</label>
                    `)
				//line views/templates/forms/anyForm.qtpl:206
				field.streamrenderSet(qw422016, tableName+key, required)
				//line views/templates/forms/anyForm.qtpl:206
				qw422016.N().S(`
                `)
			//line views/templates/forms/anyForm.qtpl:207
			default:
				//line views/templates/forms/anyForm.qtpl:207
				qw422016.N().S(`
                    <label for="`)
				//line views/templates/forms/anyForm.qtpl:208
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:208
				qw422016.N().S(`">`)
				//line views/templates/forms/anyForm.qtpl:208
				qw422016.E().S(titleLabel)
				//line views/templates/forms/anyForm.qtpl:208
				qw422016.N().S(`</label>
                    <input id="`)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.E().S(key)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.N().S(`" name="`)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.E().S(tableName + key)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.N().S(`" class="form-control" placeholder="`)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.E().S(placeholder)
				//line views/templates/forms/anyForm.qtpl:209
				qw422016.N().S(`"
                        `)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.E().S(required)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.N().S(` `)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.E().S(val)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.N().S(` `)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.N().S(events)
				//line views/templates/forms/anyForm.qtpl:210
				qw422016.N().S(`
                        `)
				//line views/templates/forms/anyForm.qtpl:211
				if dataJson > "" {
					//line views/templates/forms/anyForm.qtpl:211
					qw422016.N().S(`data-names="{`)
					//line views/templates/forms/anyForm.qtpl:211
					qw422016.N().S(dataJson)
					//line views/templates/forms/anyForm.qtpl:211
					qw422016.N().S(`" `)
					//line views/templates/forms/anyForm.qtpl:211
				}
				//line views/templates/forms/anyForm.qtpl:211
				qw422016.N().S(`

                        `)
				//line views/templates/forms/anyForm.qtpl:213
				if field.DATA_TYPE == "int" || field.DATA_TYPE == "double" {
					//line views/templates/forms/anyForm.qtpl:213
					qw422016.N().S(`
                            type="number"
                               `)
					//line views/templates/forms/anyForm.qtpl:215
					if strings.Contains(field.COLUMN_TYPE, "unsigned") {
						//line views/templates/forms/anyForm.qtpl:215
						qw422016.N().S(`min="0"`)
						//line views/templates/forms/anyForm.qtpl:215
					}
					//line views/templates/forms/anyForm.qtpl:215
					qw422016.N().S(`
                        `)
					//line views/templates/forms/anyForm.qtpl:216
				} else if strings.Contains(key, "email") {
					//line views/templates/forms/anyForm.qtpl:216
					qw422016.N().S(`
                            type="email"
                        `)
					//line views/templates/forms/anyForm.qtpl:218
				} else {
					//line views/templates/forms/anyForm.qtpl:218
					qw422016.N().S(`
                            type="text"
                               `)
					//line views/templates/forms/anyForm.qtpl:220
					if field.CHARACTER_MAXIMUM_LENGTH > 0 {
						//line views/templates/forms/anyForm.qtpl:220
						qw422016.N().S(`maxlength="`)
						//line views/templates/forms/anyForm.qtpl:220
						qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
						//line views/templates/forms/anyForm.qtpl:220
						qw422016.N().S(`"`)
						//line views/templates/forms/anyForm.qtpl:220
					}
					//line views/templates/forms/anyForm.qtpl:220
					qw422016.N().S(`
                               `)
					//line views/templates/forms/anyForm.qtpl:221
					if pattern > "" {
						//line views/templates/forms/anyForm.qtpl:221
						qw422016.N().S(`pattern="`)
						//line views/templates/forms/anyForm.qtpl:221
						qw422016.N().S(pattern)
						//line views/templates/forms/anyForm.qtpl:221
						qw422016.N().S(`" onkeyup="return validatePattern(this);"`)
						//line views/templates/forms/anyForm.qtpl:221
					}
					//line views/templates/forms/anyForm.qtpl:221
					qw422016.N().S(`
                        `)
					//line views/templates/forms/anyForm.qtpl:222
				}
				//line views/templates/forms/anyForm.qtpl:222
				qw422016.N().S(`
                    />
                `)
				//line views/templates/forms/anyForm.qtpl:224
			}
			//line views/templates/forms/anyForm.qtpl:224
			qw422016.N().S(`

        `)
			//line views/templates/forms/anyForm.qtpl:226
		}
		//line views/templates/forms/anyForm.qtpl:226
		qw422016.N().S(`
        </div>
    `)
		//line views/templates/forms/anyForm.qtpl:228
	}
	//line views/templates/forms/anyForm.qtpl:228
	qw422016.N().S(`
    <button class="btn btn-large btn-primary" type="submit">Сохранить</button>
</form>
`)
//line views/templates/forms/anyForm.qtpl:231
}

//line views/templates/forms/anyForm.qtpl:231
func (ns *FieldsTable) WriteShowAnyForm(qq422016 qtio422016.Writer, Action, Title string) {
	//line views/templates/forms/anyForm.qtpl:231
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:231
	ns.StreamShowAnyForm(qw422016, Action, Title)
	//line views/templates/forms/anyForm.qtpl:231
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:231
}

//line views/templates/forms/anyForm.qtpl:231
func (ns *FieldsTable) ShowAnyForm(Action, Title string) string {
	//line views/templates/forms/anyForm.qtpl:231
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:231
	ns.WriteShowAnyForm(qb422016, Action, Title)
	//line views/templates/forms/anyForm.qtpl:231
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:231
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:231
	return qs422016
//line views/templates/forms/anyForm.qtpl:231
}

//line views/templates/forms/anyForm.qtpl:232
func (field *FieldStructure) streamrenderSet(qw422016 *qt422016.Writer, key, required string) {
	//line views/templates/forms/anyForm.qtpl:232
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:234
	enumValidator := regexp.MustCompile(`(?:'([^,]+)',?)`)
	fields := enumValidator.FindAllStringSubmatch(field.COLUMN_TYPE, -1)
	checked := ""

	//line views/templates/forms/anyForm.qtpl:237
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:238
	for idx, title := range fields {
		//line views/templates/forms/anyForm.qtpl:238
		qw422016.N().S(`
        `)
		//line views/templates/forms/anyForm.qtpl:240
		enumVal := title[len(title)-1]
		if (field.Value > "") && (strings.Index(field.Value, enumVal) > -1) || (field.Value == "") && (enumVal == field.COLUMN_DEFAULT) {
			checked = "checked"
		} else {
			checked = ""
		}

		//line views/templates/forms/anyForm.qtpl:246
		qw422016.N().S(`
        `)
		//line views/templates/forms/anyForm.qtpl:247
		streamrenderCheckBox(qw422016, key+"[]", enumVal, idx, checked)
		//line views/templates/forms/anyForm.qtpl:247
		qw422016.N().S(`
    `)
		//line views/templates/forms/anyForm.qtpl:248
	}
	//line views/templates/forms/anyForm.qtpl:248
	qw422016.N().S(`
`)
//line views/templates/forms/anyForm.qtpl:249
}

//line views/templates/forms/anyForm.qtpl:249
func (field *FieldStructure) writerenderSet(qq422016 qtio422016.Writer, key, required string) {
	//line views/templates/forms/anyForm.qtpl:249
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:249
	field.streamrenderSet(qw422016, key, required)
	//line views/templates/forms/anyForm.qtpl:249
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:249
}

//line views/templates/forms/anyForm.qtpl:249
func (field *FieldStructure) renderSet(key, required string) string {
	//line views/templates/forms/anyForm.qtpl:249
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:249
	field.writerenderSet(qb422016, key, required)
	//line views/templates/forms/anyForm.qtpl:249
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:249
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:249
	return qs422016
//line views/templates/forms/anyForm.qtpl:249
}

//line views/templates/forms/anyForm.qtpl:250
func streamrenderCheckBox(qw422016 *qt422016.Writer, key, title string, idx int, checked string) {
	//line views/templates/forms/anyForm.qtpl:250
	qw422016.N().S(`
    <label for="`)
	//line views/templates/forms/anyForm.qtpl:251
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:251
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:251
	qw422016.N().S(`">
        <input type="checkbox" id="`)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.N().S(`" value="`)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.N().S(`" `)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.E().S(checked)
	//line views/templates/forms/anyForm.qtpl:252
	qw422016.N().S(`/>
        `)
	//line views/templates/forms/anyForm.qtpl:253
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:253
	qw422016.N().S(`
    </label>
`)
//line views/templates/forms/anyForm.qtpl:255
}

//line views/templates/forms/anyForm.qtpl:255
func writerenderCheckBox(qq422016 qtio422016.Writer, key, title string, idx int, checked string) {
	//line views/templates/forms/anyForm.qtpl:255
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:255
	streamrenderCheckBox(qw422016, key, title, idx, checked)
	//line views/templates/forms/anyForm.qtpl:255
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:255
}

//line views/templates/forms/anyForm.qtpl:255
func renderCheckBox(key, title string, idx int, checked string) string {
	//line views/templates/forms/anyForm.qtpl:255
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:255
	writerenderCheckBox(qb422016, key, title, idx, checked)
	//line views/templates/forms/anyForm.qtpl:255
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:255
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:255
	return qs422016
//line views/templates/forms/anyForm.qtpl:255
}

//line views/templates/forms/anyForm.qtpl:256
func (field *FieldStructure) streamrenderEnum(qw422016 *qt422016.Writer, key, required string) {
	//line views/templates/forms/anyForm.qtpl:256
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:258
	enumValidator := regexp.MustCompile(`'([^']+)'`)
	fields := enumValidator.FindAllStringSubmatch(field.COLUMN_TYPE, -1)
	renderSelect := len(fields) > 3

	//line views/templates/forms/anyForm.qtpl:261
	qw422016.N().S(`

    `)
	//line views/templates/forms/anyForm.qtpl:263
	if renderSelect {
		//line views/templates/forms/anyForm.qtpl:263
		qw422016.N().S(`
        <select id="`)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.E().S(key)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.N().S(`" name="`)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.E().S(key)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.N().S(`" class="form-control" `)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.E().S(required)
		//line views/templates/forms/anyForm.qtpl:264
		qw422016.N().S(` >
    `)
		//line views/templates/forms/anyForm.qtpl:265
	}
	//line views/templates/forms/anyForm.qtpl:265
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:266
	for idx, title := range fields {
		//line views/templates/forms/anyForm.qtpl:266
		qw422016.N().S(`
        `)
		//line views/templates/forms/anyForm.qtpl:268
		enumVal := title[len(title)-1]
		checked, selected := "", ""
		if (field.Value > "") && (field.Value == enumVal) || (field.Value == "") && (enumVal == field.COLUMN_DEFAULT) {
			checked, selected = "checked", "selected"
		}

		//line views/templates/forms/anyForm.qtpl:273
		qw422016.N().S(`
        `)
		//line views/templates/forms/anyForm.qtpl:274
		if renderSelect {
			//line views/templates/forms/anyForm.qtpl:274
			qw422016.N().S(`
            <option value='`)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.E().S(enumVal)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.N().S(`' `)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.E().S(selected)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.N().S(`>`)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.E().S(enumVal)
			//line views/templates/forms/anyForm.qtpl:275
			qw422016.N().S(`</option>
        `)
			//line views/templates/forms/anyForm.qtpl:276
		} else {
			//line views/templates/forms/anyForm.qtpl:276
			qw422016.N().S(`
            `)
			//line views/templates/forms/anyForm.qtpl:277
			streamrenderRadioBox(qw422016, key, enumVal, idx, checked)
			//line views/templates/forms/anyForm.qtpl:277
			qw422016.N().S(`
        `)
			//line views/templates/forms/anyForm.qtpl:278
		}
		//line views/templates/forms/anyForm.qtpl:278
		qw422016.N().S(`
    `)
		//line views/templates/forms/anyForm.qtpl:279
	}
	//line views/templates/forms/anyForm.qtpl:279
	qw422016.N().S(`
    `)
	//line views/templates/forms/anyForm.qtpl:280
	if renderSelect {
		//line views/templates/forms/anyForm.qtpl:280
		qw422016.N().S(`
        </select>
    `)
		//line views/templates/forms/anyForm.qtpl:282
	}
	//line views/templates/forms/anyForm.qtpl:282
	qw422016.N().S(`
`)
//line views/templates/forms/anyForm.qtpl:283
}

//line views/templates/forms/anyForm.qtpl:283
func (field *FieldStructure) writerenderEnum(qq422016 qtio422016.Writer, key, required string) {
	//line views/templates/forms/anyForm.qtpl:283
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:283
	field.streamrenderEnum(qw422016, key, required)
	//line views/templates/forms/anyForm.qtpl:283
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:283
}

//line views/templates/forms/anyForm.qtpl:283
func (field *FieldStructure) renderEnum(key, required string) string {
	//line views/templates/forms/anyForm.qtpl:283
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:283
	field.writerenderEnum(qb422016, key, required)
	//line views/templates/forms/anyForm.qtpl:283
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:283
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:283
	return qs422016
//line views/templates/forms/anyForm.qtpl:283
}

//line views/templates/forms/anyForm.qtpl:284
func streamrenderRadioBox(qw422016 *qt422016.Writer, key, title string, idx int, checked string) {
	//line views/templates/forms/anyForm.qtpl:284
	qw422016.N().S(`
    <label for="`)
	//line views/templates/forms/anyForm.qtpl:285
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:285
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:285
	qw422016.N().S(`">
        <input type="radio" id="`)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.N().D(idx)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.N().S(`" name="`)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.E().S(key)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.N().S(`" value="`)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.N().S(`" `)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.E().S(checked)
	//line views/templates/forms/anyForm.qtpl:286
	qw422016.N().S(`/>
        `)
	//line views/templates/forms/anyForm.qtpl:287
	qw422016.E().S(title)
	//line views/templates/forms/anyForm.qtpl:287
	qw422016.N().S(`
    </label>
`)
//line views/templates/forms/anyForm.qtpl:289
}

//line views/templates/forms/anyForm.qtpl:289
func writerenderRadioBox(qq422016 qtio422016.Writer, key, title string, idx int, checked string) {
	//line views/templates/forms/anyForm.qtpl:289
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:289
	streamrenderRadioBox(qw422016, key, title, idx, checked)
	//line views/templates/forms/anyForm.qtpl:289
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:289
}

//line views/templates/forms/anyForm.qtpl:289
func renderRadioBox(key, title string, idx int, checked string) string {
	//line views/templates/forms/anyForm.qtpl:289
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:289
	writerenderRadioBox(qb422016, key, title, idx, checked)
	//line views/templates/forms/anyForm.qtpl:289
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:289
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:289
	return qs422016
//line views/templates/forms/anyForm.qtpl:289
}

//line views/templates/forms/anyForm.qtpl:290
func (fields *FieldsTable) StreamPutDataFrom(qw422016 *qt422016.Writer, ns db.FieldsTable) {
	//line views/templates/forms/anyForm.qtpl:290
	qw422016.N().S(`
`)
	//line views/templates/forms/anyForm.qtpl:292
	for _, field := range ns.Rows {
		fieldStrc := &FieldStructure{
			COLUMN_NAME: field.COLUMN_NAME,
			DATA_TYPE:   field.DATA_TYPE,
			IS_NULLABLE: field.IS_NULLABLE,
			COLUMN_TYPE: field.COLUMN_TYPE,
			Events:      make(map[string]string, 0),
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
		fields.Rows = append(fields.Rows, *fieldStrc)
	}

	//line views/templates/forms/anyForm.qtpl:315
	qw422016.N().S(`
`)
//line views/templates/forms/anyForm.qtpl:316
}

//line views/templates/forms/anyForm.qtpl:316
func (fields *FieldsTable) WritePutDataFrom(qq422016 qtio422016.Writer, ns db.FieldsTable) {
	//line views/templates/forms/anyForm.qtpl:316
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/forms/anyForm.qtpl:316
	fields.StreamPutDataFrom(qw422016, ns)
	//line views/templates/forms/anyForm.qtpl:316
	qt422016.ReleaseWriter(qw422016)
//line views/templates/forms/anyForm.qtpl:316
}

//line views/templates/forms/anyForm.qtpl:316
func (fields *FieldsTable) PutDataFrom(ns db.FieldsTable) string {
	//line views/templates/forms/anyForm.qtpl:316
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/forms/anyForm.qtpl:316
	fields.WritePutDataFrom(qb422016, ns)
	//line views/templates/forms/anyForm.qtpl:316
	qs422016 := string(qb422016.B)
	//line views/templates/forms/anyForm.qtpl:316
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/forms/anyForm.qtpl:316
	return qs422016
//line views/templates/forms/anyForm.qtpl:316
}
