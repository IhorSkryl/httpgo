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
//

//line views/templates/json/anyjson.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/json/anyjson.qtpl:8
func StreamWriteSliceJSON(qw422016 *qt422016.Writer, mapJSON MapMultiDimension) {
	//line views/templates/json/anyjson.qtpl:8
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:10
	for key, arrJSON := range mapJSON {
		//line views/templates/json/anyjson.qtpl:11
		if key > 0 {
			//line views/templates/json/anyjson.qtpl:11
			qw422016.N().S(`,`)
			//line views/templates/json/anyjson.qtpl:11
		}
		//line views/templates/json/anyjson.qtpl:11
		StreamWriteAnyJSON(qw422016, arrJSON)
		//line views/templates/json/anyjson.qtpl:12
	}
	//line views/templates/json/anyjson.qtpl:12
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:14
}

//line views/templates/json/anyjson.qtpl:14
func WriteWriteSliceJSON(qq422016 qtio422016.Writer, mapJSON MapMultiDimension) {
	//line views/templates/json/anyjson.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:14
	StreamWriteSliceJSON(qw422016, mapJSON)
	//line views/templates/json/anyjson.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:14
}

//line views/templates/json/anyjson.qtpl:14
func WriteSliceJSON(mapJSON MapMultiDimension) string {
	//line views/templates/json/anyjson.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:14
	WriteWriteSliceJSON(qb422016, mapJSON)
	//line views/templates/json/anyjson.qtpl:14
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:14
	return qs422016
//line views/templates/json/anyjson.qtpl:14
}

// получаем объект произвольной формы и возвращаем JSON текстом

//line views/templates/json/anyjson.qtpl:17
func StreamWriteAnyJSON(qw422016 *qt422016.Writer, arrJSON MultiDimension) {
	//line views/templates/json/anyjson.qtpl:17
	qw422016.N().S(`{`)
	//line views/templates/json/anyjson.qtpl:18
	comma := ""

	//line views/templates/json/anyjson.qtpl:19
	for key, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:20
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:20
		qw422016.N().S(`"`)
		//line views/templates/json/anyjson.qtpl:20
		qw422016.E().S(key)
		//line views/templates/json/anyjson.qtpl:20
		qw422016.N().S(`":`)
		//line views/templates/json/anyjson.qtpl:20
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:20
		comma = ","

		//line views/templates/json/anyjson.qtpl:21
	}
	//line views/templates/json/anyjson.qtpl:21
	qw422016.N().S(`}`)
//line views/templates/json/anyjson.qtpl:23
}

//line views/templates/json/anyjson.qtpl:23
func WriteWriteAnyJSON(qq422016 qtio422016.Writer, arrJSON MultiDimension) {
	//line views/templates/json/anyjson.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:23
	StreamWriteAnyJSON(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:23
}

//line views/templates/json/anyjson.qtpl:23
func WriteAnyJSON(arrJSON MultiDimension) string {
	//line views/templates/json/anyjson.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:23
	WriteWriteAnyJSON(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:23
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:23
	return qs422016
//line views/templates/json/anyjson.qtpl:23
}

// пишем элемент массива в зависемости от типа

//line views/templates/json/anyjson.qtpl:26
func StreamWriteElement(qw422016 *qt422016.Writer, value interface{}) {
	//line views/templates/json/anyjson.qtpl:27
	switch vv := value.(type) {
	//line views/templates/json/anyjson.qtpl:28
	case string:
		//line views/templates/json/anyjson.qtpl:28
		qw422016.N().S(`"`)
		//line views/templates/json/anyjson.qtpl:28
		qw422016.N().J(vv)
		//line views/templates/json/anyjson.qtpl:28
		qw422016.N().S(`"`)
	//line views/templates/json/anyjson.qtpl:29
	case bool:
		//line views/templates/json/anyjson.qtpl:29
		qw422016.E().V(vv)
	//line views/templates/json/anyjson.qtpl:30
	case int, uint, int32, int64:
		//line views/templates/json/anyjson.qtpl:30
		qw422016.N().D(vv.(int))
	//line views/templates/json/anyjson.qtpl:31
	case float64:
		//line views/templates/json/anyjson.qtpl:31
		qw422016.N().F(vv)
	//line views/templates/json/anyjson.qtpl:32
	case nil:
		//line views/templates/json/anyjson.qtpl:32
		qw422016.N().S(`null`)
	//line views/templates/json/anyjson.qtpl:33
	case StringDimension:
		//line views/templates/json/anyjson.qtpl:34
		StreamWriteStringDimension(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:35
	case SimpleDimension:
		//line views/templates/json/anyjson.qtpl:36
		StreamWriteSimpleDimension(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:37
	case MultiDimension:
		//line views/templates/json/anyjson.qtpl:38
		StreamWriteAnyJSON(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:39
	case MapMultiDimension:
		//line views/templates/json/anyjson.qtpl:40
		StreamWriteSliceJSON(qw422016, vv)
	//line views/templates/json/anyjson.qtpl:41
	default:
		//line views/templates/json/anyjson.qtpl:42
		if vSimple, ok := isSimpleDimension(vv); ok {
			//line views/templates/json/anyjson.qtpl:43
			StreamWriteSimpleDimension(qw422016, vSimple)
			//line views/templates/json/anyjson.qtpl:44
		} else if arrJSON, ok := isMultiDimension(vv); ok {
			//line views/templates/json/anyjson.qtpl:45
			StreamWriteAnyJSON(qw422016, arrJSON)
			//line views/templates/json/anyjson.qtpl:46
		} else if mapArrJSON, ok := isMapMultiDimension(vv); ok {
			//line views/templates/json/anyjson.qtpl:47
			StreamWriteSliceJSON(qw422016, mapArrJSON)
			//line views/templates/json/anyjson.qtpl:48
		} else {
			//line views/templates/json/anyjson.qtpl:48
			qw422016.N().S(`"`)
			//line views/templates/json/anyjson.qtpl:49
			qw422016.E().V(vv)
			//line views/templates/json/anyjson.qtpl:49
			qw422016.N().S(`"`)
			//line views/templates/json/anyjson.qtpl:50
		}
		//line views/templates/json/anyjson.qtpl:51
	}
//line views/templates/json/anyjson.qtpl:52
}

//line views/templates/json/anyjson.qtpl:52
func WriteWriteElement(qq422016 qtio422016.Writer, value interface{}) {
	//line views/templates/json/anyjson.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:52
	StreamWriteElement(qw422016, value)
	//line views/templates/json/anyjson.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:52
}

//line views/templates/json/anyjson.qtpl:52
func WriteElement(value interface{}) string {
	//line views/templates/json/anyjson.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:52
	WriteWriteElement(qb422016, value)
	//line views/templates/json/anyjson.qtpl:52
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:52
	return qs422016
//line views/templates/json/anyjson.qtpl:52
}

// получаем массив произвольной формы и возвращаем JSON текстом

//line views/templates/json/anyjson.qtpl:54
func StreamWriteArrJSON(qw422016 *qt422016.Writer, arrJSON []interface{}) {
	//line views/templates/json/anyjson.qtpl:54
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:55
	comma := ""

	//line views/templates/json/anyjson.qtpl:56
	for _, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:56
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:56
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:56
		comma = ","

		//line views/templates/json/anyjson.qtpl:57
	}
	//line views/templates/json/anyjson.qtpl:57
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:59
}

//line views/templates/json/anyjson.qtpl:59
func WriteWriteArrJSON(qq422016 qtio422016.Writer, arrJSON []interface{}) {
	//line views/templates/json/anyjson.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:59
	StreamWriteArrJSON(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:59
}

//line views/templates/json/anyjson.qtpl:59
func WriteArrJSON(arrJSON []interface{}) string {
	//line views/templates/json/anyjson.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:59
	WriteWriteArrJSON(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:59
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:59
	return qs422016
//line views/templates/json/anyjson.qtpl:59
}

// получаем массив объектов произвольной формы и возвращаем JSON текстом

//line views/templates/json/anyjson.qtpl:61
func StreamWriteSimpleDimension(qw422016 *qt422016.Writer, arrJSON SimpleDimension) {
	//line views/templates/json/anyjson.qtpl:61
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:62
	comma := ""

	//line views/templates/json/anyjson.qtpl:63
	for _, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:63
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:63
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:63
		comma = ","

		//line views/templates/json/anyjson.qtpl:64
	}
	//line views/templates/json/anyjson.qtpl:64
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:66
}

//line views/templates/json/anyjson.qtpl:66
func WriteWriteSimpleDimension(qq422016 qtio422016.Writer, arrJSON SimpleDimension) {
	//line views/templates/json/anyjson.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:66
	StreamWriteSimpleDimension(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:66
}

//line views/templates/json/anyjson.qtpl:66
func WriteSimpleDimension(arrJSON SimpleDimension) string {
	//line views/templates/json/anyjson.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:66
	WriteWriteSimpleDimension(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:66
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:66
	return qs422016
//line views/templates/json/anyjson.qtpl:66
}

// получаем массив строк и возвращаем JSON текстом

//line views/templates/json/anyjson.qtpl:68
func StreamWriteStringDimension(qw422016 *qt422016.Writer, arrJSON StringDimension) {
	//line views/templates/json/anyjson.qtpl:68
	qw422016.N().S(`[`)
	//line views/templates/json/anyjson.qtpl:69
	comma := ""

	//line views/templates/json/anyjson.qtpl:70
	for _, value := range arrJSON {
		//line views/templates/json/anyjson.qtpl:70
		qw422016.E().S(comma)
		//line views/templates/json/anyjson.qtpl:70
		StreamWriteElement(qw422016, value)
		//line views/templates/json/anyjson.qtpl:70
		comma = ","

		//line views/templates/json/anyjson.qtpl:71
	}
	//line views/templates/json/anyjson.qtpl:71
	qw422016.N().S(`]`)
//line views/templates/json/anyjson.qtpl:73
}

//line views/templates/json/anyjson.qtpl:73
func WriteWriteStringDimension(qq422016 qtio422016.Writer, arrJSON StringDimension) {
	//line views/templates/json/anyjson.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/json/anyjson.qtpl:73
	StreamWriteStringDimension(qw422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line views/templates/json/anyjson.qtpl:73
}

//line views/templates/json/anyjson.qtpl:73
func WriteStringDimension(arrJSON StringDimension) string {
	//line views/templates/json/anyjson.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/json/anyjson.qtpl:73
	WriteWriteStringDimension(qb422016, arrJSON)
	//line views/templates/json/anyjson.qtpl:73
	qs422016 := string(qb422016.B)
	//line views/templates/json/anyjson.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/json/anyjson.qtpl:73
	return qs422016
//line views/templates/json/anyjson.qtpl:73
}
