// Code generated by qtc from "table_test.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/tables/table_test.qtpl:1
package tables

//line views/templates/tables/table_test.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/tables/table_test.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/tables/table_test.qtpl:1
func StreamTableTesting(qw422016 *qt422016.Writer) {
//line views/templates/tables/table_test.qtpl:1
	qw422016.N().S(`
<html>
<script type="text/javascript" src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
<link rel="import" href="/components/paper-datatable/paper-datatable-card.html">
<link rel="import" href="/components/paper-datatable/paper-datatable.html">
<iron-ajax url="/bookings/check.json" last-response="{{data}}" auto></iron-ajax>
<template is="dom-bind" id="app">
<paper-datatable data="[[data]]" >
        <template is="dom-repeat" items="[[data.items]]">
	<paper-datatable-column header="[[item[0]]]" property="[[item[0][1]]]" type="Number" sortable></paper-datatable-column>
        </template>
</paper-datatable>
</template>
<script>
		$(function(){
			Polymer({
                is:"dom-bind",
                    data:{
                        type: Array,
                        notify: true
                    },
             });
            })
</script>
    <script  src="/webcomponents.js"></script>
`)
//line views/templates/tables/table_test.qtpl:26
}

//line views/templates/tables/table_test.qtpl:26
func WriteTableTesting(qq422016 qtio422016.Writer) {
//line views/templates/tables/table_test.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/templates/tables/table_test.qtpl:26
	StreamTableTesting(qw422016)
//line views/templates/tables/table_test.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line views/templates/tables/table_test.qtpl:26
}

//line views/templates/tables/table_test.qtpl:26
func TableTesting() string {
//line views/templates/tables/table_test.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/templates/tables/table_test.qtpl:26
	WriteTableTesting(qb422016)
//line views/templates/tables/table_test.qtpl:26
	qs422016 := string(qb422016.B)
//line views/templates/tables/table_test.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/templates/tables/table_test.qtpl:26
	return qs422016
//line views/templates/tables/table_test.qtpl:26
}
