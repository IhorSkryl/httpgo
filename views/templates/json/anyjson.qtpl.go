// This file is automatically generated by qtc from "anyjson.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:1
package json

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
// выводим массив массивов (основное назначение для таблиц БД)
//

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:5
import (
	"strconv"
)

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:12
func StreamWriteSliceJSON(qw422016 *qt422016.Writer, mapJSON MapMultiDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:12
	qw422016.N().S(`{`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:14
	for key, arrJSON := range mapJSON {
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		if key > 0 {
			//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
			qw422016.N().S(`,`)
			//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		}
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		qw422016.N().S(`"`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		qw422016.E().S(strconv.Itoa(key))
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		qw422016.N().S(`":`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:15
		StreamWriteAnyJSON(qw422016, arrJSON)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:16
	}
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:16
	qw422016.N().S(`}`)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
func WriteWriteSliceJSON(qq422016 qtio422016.Writer, mapJSON MapMultiDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	StreamWriteSliceJSON(qw422016, mapJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
func WriteSliceJSON(mapJSON MapMultiDimension) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	WriteWriteSliceJSON(qb422016, mapJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:18
}

// получаем объект произвольной формы и возвращаем JSON текстом

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:21
func StreamWriteAnyJSON(qw422016 *qt422016.Writer, arrJSON MultiDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:21
	qw422016.N().S(`{`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:22
	comma := ""

	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:23
	for key, value := range arrJSON {
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		qw422016.E().S(comma)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		qw422016.N().S(`"`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		qw422016.E().S(key)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		qw422016.N().S(`":`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		StreamWriteElement(qw422016, value)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:24
		comma = ","

		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:25
	}
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:25
	qw422016.N().S(`}`)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
func WriteWriteAnyJSON(qq422016 qtio422016.Writer, arrJSON MultiDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	StreamWriteAnyJSON(qw422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
func WriteAnyJSON(arrJSON MultiDimension) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	WriteWriteAnyJSON(qb422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:27
}

// пишем элемент массива в зависемости от типа

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:32
func StreamWriteElement(qw422016 *qt422016.Writer, value interface{}) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:33
	switch vv := value.(type) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:34
	case string:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:34
		qw422016.N().S(`"`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:34
		qw422016.E().Q(vv)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:34
		qw422016.N().S(`"`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:35
	case bool:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:35
		qw422016.E().V(vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:36
	case int, uint, int32, int64:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:36
		qw422016.N().D(vv.(int))
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:37
	case float64:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:37
		qw422016.N().F(vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:38
	case nil:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:38
		qw422016.N().S(`null`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:39
	case StringDimension:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:39
		StreamWriteStringDimension(qw422016, vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:40
	case SimpleDimension:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:40
		StreamWriteSimpleDimension(qw422016, vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:41
	case MultiDimension:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:41
		StreamWriteAnyJSON(qw422016, vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:42
	case MapMultiDimension:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:42
		StreamWriteSliceJSON(qw422016, vv)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:43
	default:
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:43
		qw422016.E().V(vv)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:44
	}
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
func WriteWriteElement(qq422016 qtio422016.Writer, value interface{}) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	StreamWriteElement(qw422016, value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
func WriteElement(value interface{}) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	WriteWriteElement(qb422016, value)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:45
}

// получаем массив произвольной формы и возвращаем JSON текстом

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:47
func StreamWriteArrJSON(qw422016 *qt422016.Writer, arrJSON []interface{}) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:47
	qw422016.N().S(`[`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:48
	comma := ""

	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
	for key, value := range arrJSON {
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
		qw422016.E().S(comma)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
		qw422016.N().D(key)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
		qw422016.N().S(`:`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
		StreamWriteElement(qw422016, value)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:49
		comma = ","

		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:50
	}
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:50
	qw422016.N().S(`]`)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
func WriteWriteArrJSON(qq422016 qtio422016.Writer, arrJSON []interface{}) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	StreamWriteArrJSON(qw422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
func WriteArrJSON(arrJSON []interface{}) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	WriteWriteArrJSON(qb422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:52
}

// получаем массив объектов произвольной формы и возвращаем JSON текстом

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:54
func StreamWriteSimpleDimension(qw422016 *qt422016.Writer, arrJSON SimpleDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:54
	qw422016.N().S(`[`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:55
	comma := ""

	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
	for key, value := range arrJSON {
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
		qw422016.E().S(comma)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
		qw422016.N().D(key)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
		qw422016.N().S(`:`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
		StreamWriteElement(qw422016, value)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:56
		comma = ","

		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:57
	}
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:57
	qw422016.N().S(`]`)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
func WriteWriteSimpleDimension(qq422016 qtio422016.Writer, arrJSON SimpleDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	StreamWriteSimpleDimension(qw422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
func WriteSimpleDimension(arrJSON SimpleDimension) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	WriteWriteSimpleDimension(qb422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:59
}

// получаем массив строк и возвращаем JSON текстом

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:61
func StreamWriteStringDimension(qw422016 *qt422016.Writer, arrJSON StringDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:61
	qw422016.N().S(`[`)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:62
	comma := ""

	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
	for key, value := range arrJSON {
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
		qw422016.E().S(comma)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
		qw422016.N().D(key)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
		qw422016.N().S(`:`)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
		StreamWriteElement(qw422016, value)
		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:63
		comma = ","

		//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:64
	}
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:64
	qw422016.N().S(`]`)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
func WriteWriteStringDimension(qq422016 qtio422016.Writer, arrJSON StringDimension) {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	StreamWriteStringDimension(qw422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
}

//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
func WriteStringDimension(arrJSON StringDimension) string {
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	WriteWriteStringDimension(qb422016, arrJSON)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	qs422016 := string(qb422016.B)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
	//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
	return qs422016
//line src/github.com/ruslanBik4/httpgo/views/templates/json/anyjson.qtpl:66
}
