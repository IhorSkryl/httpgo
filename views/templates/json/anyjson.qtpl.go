// This file is automatically generated by qtc from "anyjson.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/json/anyjson.qtpl:1
package json

//line views/templates/json/anyjson.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// All the text outside function templates is treated as comments,
// i.e. it is just ignored by quicktemplate compiler (`qtc`). It is for humans.
// выводим массив массивов (основное назначение для таблиц БД)

//line views/templates/json/anyjson.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/json/anyjson.qtpl:5
func StreamWriteSliceJSON(qw422016 *qt422016.Writer, mapJSON MapMultiDimension) {
	//line views/templates/json/anyjson.qtpl:6
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:7
	comma := ""

	//line views/templates/json/anyjson.qtpl:8
	for key, arrJSON := range mapJSON {
		//line views/templates/json/anyjson.qtpl:8
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:8
		qw422016.N().D(key)
		//line views/templates/json/anyjson.qtpl:8
		qw422016.N().S(`:`)
		//line views/templates/json/anyjson.qtpl:8
		StreamWriteAnyJSON(qw422016, arrJSON)
		//line views/templates/json/anyjson.qtpl:8
		comma = ","

		//line views/templates/json/anyjson.qtpl:9
	}
	//line views/templates/json/anyjson.qtpl:9
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:11
}

//line views/templates/json/anyjson.qtpl:11
func WriteWriteSliceJSON(qq422016 qtio422016.Writer, mapJSON MapMultiDimension) {
	//line views/templates/json/anyjson.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:11
	StreamWriteSliceJSON(qw422016, mapJSON)
	//line views/templates/json/anyjson.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:11
}

//line views/templates/json/anyjson.qtpl:11
func WriteSliceJSON(mapJSON MapMultiDimension) string {
	//line views/templates/json/anyjson.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:11
	WriteWriteSliceJSON(qb422016, mapJSON)
	//line views/templates/json/anyjson.qtpl:11
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:11
	return qs422016
//line views/templates/json/anyjson.qtpl:11
}

// получаем объект произвольной формы и возвращаем JSON текстом

//line views/templates/json/anyjson.qtpl:14
func StreamWriteAnyJSON(qw422016 *qt422016.Writer, arrJSON MultiDimension) {
	//line views/templates/json/anyjson.qtpl:15
	qw422016.N().S(`{`)
	//line views/templates/json/anyjson.qtpl:16
	comma := ""

	//line views/templates/json/anyjson.qtpl:17
	for key, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:17
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:17
		qw422016.N().S(`"`)
		//line views/templates/json/anyjson.qtpl:17
		qw422016.E().S(key)
		//line views/templates/json/anyjson.qtpl:17
		qw422016.N().S(`":`)
		//line views/templates/json/anyjson.qtpl:17
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:17
		comma = ","

		//line views/templates/json/anyjson.qtpl:18
	}
	//line views/templates/json/anyjson.qtpl:18
	qw422016.N().S(`}`)
//line views/templates/json/anyjson.qtpl:21
}

//line views/templates/json/anyjson.qtpl:21
func WriteWriteAnyJSON(qq422016 qtio422016.Writer, arrJSON MultiDimension) {
	//line views/templates/json/anyjson.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:21
	StreamWriteAnyJSON(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:21
}

//line views/templates/json/anyjson.qtpl:21
func WriteAnyJSON(arrJSON MultiDimension) string {
	//line views/templates/json/anyjson.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:21
	WriteWriteAnyJSON(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:21
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:21
	return qs422016
//line views/templates/json/anyjson.qtpl:21
}

// пишем элемент массива в зависемости от типа

//line views/templates/json/anyjson.qtpl:23
func StreamWriteElement(qw422016 *qt422016.Writer, value interface{}) {
	//line views/templates/json/anyjson.qtpl:25
	switch vv := value.(type) {
	//line views/templates/json/anyjson.qtpl:26
	case string:
		//line views/templates/json/anyjson.qtpl:26
		qw422016.N().S(`"`)
		//line views/templates/json/anyjson.qtpl:26
		qw422016.E().Q(vv)
		//line views/templates/json/anyjson.qtpl:26
		qw422016.N().S(`"`)
	//line views/templates/json/anyjson.qtpl:27
	case bool:
		//line views/templates/json/anyjson.qtpl:27
		qw422016.E().V(vv)
	//line views/templates/json/anyjson.qtpl:28
	case int, uint, int32, int64:
		//line views/templates/json/anyjson.qtpl:28
		qw422016.N().D(vv.(int))
	//line views/templates/json/anyjson.qtpl:29
	case float64:
		//line views/templates/json/anyjson.qtpl:29
		qw422016.N().F(vv)
	//line views/templates/json/anyjson.qtpl:30
	case nil:
		//line views/templates/json/anyjson.qtpl:30
		qw422016.N().S(`NULL`)
	//line views/templates/json/anyjson.qtpl:31
	case StringDimension:
		//line views/templates/json/anyjson.qtpl:31
		StreamWriteStringDimension(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:32
	case SimpleDimension:
		//line views/templates/json/anyjson.qtpl:32
		StreamWriteSimpleDimension(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:33
	case MultiDimension:
		//line views/templates/json/anyjson.qtpl:33
		qw422016.E().S(WriteAnyJSON(vv))
	//line views/templates/json/anyjson.qtpl:34
	default:
		//line views/templates/json/anyjson.qtpl:34
		qw422016.E().V(vv)
		//line views/templates/json/anyjson.qtpl:35
	}
//line views/templates/json/anyjson.qtpl:37
}

//line views/templates/json/anyjson.qtpl:37
func WriteWriteElement(qq422016 qtio422016.Writer, value interface{}) {
	//line views/templates/json/anyjson.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:37
	StreamWriteElement(qw422016, value)
	//line views/templates/json/anyjson.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:37
}

//line views/templates/json/anyjson.qtpl:37
func WriteElement(value interface{}) string {
	//line views/templates/json/anyjson.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:37
	WriteWriteElement(qb422016, value)
	//line views/templates/json/anyjson.qtpl:37
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:37
	return qs422016
//line views/templates/json/anyjson.qtpl:37
}

//line views/templates/json/anyjson.qtpl:39
func StreamWriteArrJSON(qw422016 *qt422016.Writer, arrJSON []interface{}) {
	//line views/templates/json/anyjson.qtpl:40
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:41
	comma := ""

	//line views/templates/json/anyjson.qtpl:42
	for key, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:42
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:42
		qw422016.N().D(key)
		//line views/templates/json/anyjson.qtpl:42
		qw422016.N().S(`:`)
		//line views/templates/json/anyjson.qtpl:42
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:42
		comma = ","

		//line views/templates/json/anyjson.qtpl:43
	}
	//line views/templates/json/anyjson.qtpl:43
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:46
}

//line views/templates/json/anyjson.qtpl:46
func WriteWriteArrJSON(qq422016 qtio422016.Writer, arrJSON []interface{}) {
	//line views/templates/json/anyjson.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:46
	StreamWriteArrJSON(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:46
}

//line views/templates/json/anyjson.qtpl:46
func WriteArrJSON(arrJSON []interface{}) string {
	//line views/templates/json/anyjson.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:46
	WriteWriteArrJSON(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:46
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:46
	return qs422016
//line views/templates/json/anyjson.qtpl:46
}

//line views/templates/json/anyjson.qtpl:47
func StreamWriteSimpleDimension(qw422016 *qt422016.Writer, arrJSON SimpleDimension) {
	//line views/templates/json/anyjson.qtpl:48
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:49
	comma := ""

	//line views/templates/json/anyjson.qtpl:50
	for key, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:50
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:50
		qw422016.N().D(key)
		//line views/templates/json/anyjson.qtpl:50
		qw422016.N().S(`:`)
		//line views/templates/json/anyjson.qtpl:50
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:50
		comma = ","

		//line views/templates/json/anyjson.qtpl:51
	}
	//line views/templates/json/anyjson.qtpl:51
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:54
}

//line views/templates/json/anyjson.qtpl:54
func WriteWriteSimpleDimension(qq422016 qtio422016.Writer, arrJSON SimpleDimension) {
	//line views/templates/json/anyjson.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:54
	StreamWriteSimpleDimension(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:54
}

//line views/templates/json/anyjson.qtpl:54
func WriteSimpleDimension(arrJSON SimpleDimension) string {
	//line views/templates/json/anyjson.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:54
	WriteWriteSimpleDimension(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:54
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:54
	return qs422016
//line views/templates/json/anyjson.qtpl:54
}

//line views/templates/json/anyjson.qtpl:55
func StreamWriteStringDimension(qw422016 *qt422016.Writer, arrJSON StringDimension) {
	//line views/templates/json/anyjson.qtpl:56
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:57
	comma := ""

	//line views/templates/json/anyjson.qtpl:58
	for key, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:58
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:58
		qw422016.N().D(key)
		//line views/templates/json/anyjson.qtpl:58
		qw422016.N().S(`:`)
		//line views/templates/json/anyjson.qtpl:58
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:58
		comma = ","

		//line views/templates/json/anyjson.qtpl:59
	}
	//line views/templates/json/anyjson.qtpl:59
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:62
}

//line views/templates/json/anyjson.qtpl:62
func WriteWriteStringDimension(qq422016 qtio422016.Writer, arrJSON StringDimension) {
	//line views/templates/json/anyjson.qtpl:62
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:62
	StreamWriteStringDimension(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:62
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:62
}

//line views/templates/json/anyjson.qtpl:62
func WriteStringDimension(arrJSON StringDimension) string {
	//line views/templates/json/anyjson.qtpl:62
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:62
	WriteWriteStringDimension(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:62
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:62
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:62
	return qs422016
//line views/templates/json/anyjson.qtpl:62
}
