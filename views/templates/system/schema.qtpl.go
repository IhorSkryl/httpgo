// This file is automatically generated by qtc from "schema.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/system/schema.qtpl:1
package system

//line views/templates/system/schema.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/system/schema.qtpl:1
import (
	"github.com/ruslanBik4/httpgo/models/db/schema"
)

// Показ схемы определенной таблицы.

//line views/templates/system/schema.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/system/schema.qtpl:6
func StreamShowSchema(qw422016 *qt422016.Writer, fields *schema.FieldsTable) {
	//line views/templates/system/schema.qtpl:6
	qw422016.N().S(`
    <div>
    `)
	//line views/templates/system/schema.qtpl:8
	for idx, field := range fields.Rows {
		//line views/templates/system/schema.qtpl:8
		qw422016.N().S(`
        <div id="divField`)
		//line views/templates/system/schema.qtpl:9
		qw422016.N().D(idx)
		//line views/templates/system/schema.qtpl:9
		qw422016.N().S(`" class="input-wrap `)
		//line views/templates/system/schema.qtpl:9
		qw422016.E().S(field.CSSClass)
		//line views/templates/system/schema.qtpl:9
		qw422016.N().S(`" >
        <div> COLUMN_NAME   <b>	`)
		//line views/templates/system/schema.qtpl:10
		qw422016.E().S(field.COLUMN_NAME)
		//line views/templates/system/schema.qtpl:10
		qw422016.N().S(` </b> </div>
        <div> DATA_TYPE 	<b> `)
		//line views/templates/system/schema.qtpl:11
		qw422016.E().S(field.DATA_TYPE)
		//line views/templates/system/schema.qtpl:11
		qw422016.N().S(` </b></div>
        <div> COLUMN_DEFAULT 	<b> `)
		//line views/templates/system/schema.qtpl:12
		qw422016.E().S(field.COLUMN_DEFAULT)
		//line views/templates/system/schema.qtpl:12
		qw422016.N().S(` </b></div>
        <div> IS_NULLABLE 	<b> `)
		//line views/templates/system/schema.qtpl:13
		qw422016.E().S(field.IS_NULLABLE)
		//line views/templates/system/schema.qtpl:13
		qw422016.N().S(` </b></div>
        <div> CHARACTER_SET_NAME <b> `)
		//line views/templates/system/schema.qtpl:14
		qw422016.E().S(field.CHARACTER_SET_NAME)
		//line views/templates/system/schema.qtpl:14
		qw422016.N().S(` </b></div>
        <div> COLUMN_COMMENT 	<b> `)
		//line views/templates/system/schema.qtpl:15
		qw422016.E().S(field.COLUMN_COMMENT)
		//line views/templates/system/schema.qtpl:15
		qw422016.N().S(` </b></div>
        <div> COLUMN_TYPE 	<b> `)
		//line views/templates/system/schema.qtpl:16
		qw422016.E().S(field.COLUMN_TYPE)
		//line views/templates/system/schema.qtpl:16
		qw422016.N().S(` </b></div>
        <div> CHARACTER_MAXIMUM_LENGTH <b> `)
		//line views/templates/system/schema.qtpl:17
		qw422016.N().D(field.CHARACTER_MAXIMUM_LENGTH)
		//line views/templates/system/schema.qtpl:17
		qw422016.N().S(` </b></div>
        <div> Value 		<b> `)
		//line views/templates/system/schema.qtpl:18
		qw422016.E().S(field.Value)
		//line views/templates/system/schema.qtpl:18
		qw422016.N().S(` </b></div>
        <div> IsHidden 	<b> `)
		//line views/templates/system/schema.qtpl:19
		qw422016.E().V(field.IsHidden)
		//line views/templates/system/schema.qtpl:19
		qw422016.N().S(` </b></div>
        <div> InputType	<b> `)
		//line views/templates/system/schema.qtpl:20
		qw422016.E().S(field.InputType)
		//line views/templates/system/schema.qtpl:20
		qw422016.N().S(` </b></div>
        <div> CSSClass  	<b> `)
		//line views/templates/system/schema.qtpl:21
		qw422016.E().S(field.CSSClass)
		//line views/templates/system/schema.qtpl:21
		qw422016.N().S(` </b></div>
        <div> CSSStyle    <b> `)
		//line views/templates/system/schema.qtpl:22
		qw422016.E().S(field.CSSStyle)
		//line views/templates/system/schema.qtpl:22
		qw422016.N().S(` </b></div>
        <div> TableName 	<b> `)
		//line views/templates/system/schema.qtpl:23
		qw422016.E().S(field.TableName)
		//line views/templates/system/schema.qtpl:23
		qw422016.N().S(` </b></div>
        <div> Where 		<b> `)
		//line views/templates/system/schema.qtpl:24
		qw422016.E().S(field.Where)
		//line views/templates/system/schema.qtpl:24
		qw422016.N().S(` </b></div>
        <div> Figure 		<b> `)
		//line views/templates/system/schema.qtpl:25
		qw422016.E().S(field.Figure)
		//line views/templates/system/schema.qtpl:25
		qw422016.N().S(` </b></div>
        <div> Placeholder	<b> `)
		//line views/templates/system/schema.qtpl:26
		qw422016.E().S(field.Placeholder)
		//line views/templates/system/schema.qtpl:26
		qw422016.N().S(` </b></div>
        <div> Pattern		<b> `)
		//line views/templates/system/schema.qtpl:27
		qw422016.E().S(field.Pattern)
		//line views/templates/system/schema.qtpl:27
		qw422016.N().S(` </b></div>
        <div> MinDate		<b> `)
		//line views/templates/system/schema.qtpl:28
		qw422016.E().S(field.MinDate)
		//line views/templates/system/schema.qtpl:28
		qw422016.N().S(` </b></div>
        <div> MaxDate		<b> `)
		//line views/templates/system/schema.qtpl:29
		qw422016.E().S(field.MaxDate)
		//line views/templates/system/schema.qtpl:29
		qw422016.N().S(` </b></div>
        <div> BeforeHtml	<b> `)
		//line views/templates/system/schema.qtpl:30
		qw422016.E().S(field.BeforeHtml)
		//line views/templates/system/schema.qtpl:30
		qw422016.N().S(` </b></div>
        <div> Html		<b> `)
		//line views/templates/system/schema.qtpl:31
		qw422016.E().S(field.Html)
		//line views/templates/system/schema.qtpl:31
		qw422016.N().S(` </b></div>
        <div> AfterHtml	<b> `)
		//line views/templates/system/schema.qtpl:32
		qw422016.E().S(field.AfterHtml)
		//line views/templates/system/schema.qtpl:32
		qw422016.N().S(` </b></div>
        <div> ForeignFields	<b> `)
		//line views/templates/system/schema.qtpl:33
		qw422016.E().S(field.ForeignFields)
		//line views/templates/system/schema.qtpl:33
		qw422016.N().S(` </b></div>
        <div> LinkTD		<b> `)
		//line views/templates/system/schema.qtpl:34
		qw422016.E().S(field.LinkTD)
		//line views/templates/system/schema.qtpl:34
		qw422016.N().S(` </b></div>
        `)
		//line views/templates/system/schema.qtpl:35
		for name, value := range field.Events {
			//line views/templates/system/schema.qtpl:35
			qw422016.N().S(`
            <div>   `)
			//line views/templates/system/schema.qtpl:36
			qw422016.E().S(name)
			//line views/templates/system/schema.qtpl:36
			qw422016.N().S(` <b>`)
			//line views/templates/system/schema.qtpl:36
			qw422016.E().S(value)
			//line views/templates/system/schema.qtpl:36
			qw422016.N().S(` </b>  </div>
        `)
			//line views/templates/system/schema.qtpl:37
		}
		//line views/templates/system/schema.qtpl:37
		qw422016.N().S(`
        `)
		//line views/templates/system/schema.qtpl:38
		for name, value := range field.DataJSOM {
			//line views/templates/system/schema.qtpl:38
			qw422016.N().S(`
         <div>  `)
			//line views/templates/system/schema.qtpl:39
			qw422016.E().S(name)
			//line views/templates/system/schema.qtpl:39
			qw422016.N().S(` <b>`)
			//line views/templates/system/schema.qtpl:39
			qw422016.E().S(value.(string))
			//line views/templates/system/schema.qtpl:39
			qw422016.N().S(` </b> </div>
        `)
			//line views/templates/system/schema.qtpl:40
		}
		//line views/templates/system/schema.qtpl:40
		qw422016.N().S(`
        `)
		//line views/templates/system/schema.qtpl:41
		for i, value := range field.EnumValues {
			//line views/templates/system/schema.qtpl:41
			qw422016.N().S(`
       <div>  `)
			//line views/templates/system/schema.qtpl:42
			qw422016.N().D(i)
			//line views/templates/system/schema.qtpl:42
			qw422016.N().S(` <b>`)
			//line views/templates/system/schema.qtpl:42
			qw422016.E().S(value)
			//line views/templates/system/schema.qtpl:42
			qw422016.N().S(` </b> </div>
        `)
			//line views/templates/system/schema.qtpl:43
		}
		//line views/templates/system/schema.qtpl:43
		qw422016.N().S(`
    `)
		//line views/templates/system/schema.qtpl:44
	}
	//line views/templates/system/schema.qtpl:44
	qw422016.N().S(`
    </div>
`)
//line views/templates/system/schema.qtpl:46
}

//line views/templates/system/schema.qtpl:46
func WriteShowSchema(qq422016 qtio422016.Writer, fields *schema.FieldsTable) {
	//line views/templates/system/schema.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/system/schema.qtpl:46
	StreamShowSchema(qw422016, fields)
	//line views/templates/system/schema.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line views/templates/system/schema.qtpl:46
}

//line views/templates/system/schema.qtpl:46
func ShowSchema(fields *schema.FieldsTable) string {
	//line views/templates/system/schema.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/system/schema.qtpl:46
	WriteShowSchema(qb422016, fields)
	//line views/templates/system/schema.qtpl:46
	qs422016 := string(qb422016.B)
	//line views/templates/system/schema.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/system/schema.qtpl:46
	return qs422016
//line views/templates/system/schema.qtpl:46
}
