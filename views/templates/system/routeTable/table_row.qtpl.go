// This file is automatically generated by qtc from "table_row.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/system/routeTable/table_row.qtpl:1
package routeTable

//line views/templates/system/routeTable/table_row.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/system/routeTable/table_row.qtpl:1
import (
	"time"
)

//line views/templates/system/routeTable/table_row.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/system/routeTable/table_row.qtpl:5
func StreamTableRow(qw422016 *qt422016.Writer, columns []string, rows [][]interface{}) {
	//line views/templates/system/routeTable/table_row.qtpl:5
	qw422016.N().S(`
<style>
`)
	//line views/templates/system/routeTable/table_row.qtpl:7
	StreamTableCSS(qw422016)
	//line views/templates/system/routeTable/table_row.qtpl:7
	qw422016.N().S(`}
</style>
<div class="usr-table">
    <div class="usr-table-header">
        <div class="usr-table__t-head  usr-table-row">
        `)
	//line views/templates/system/routeTable/table_row.qtpl:12
	for i, col := range columns {
		//line views/templates/system/routeTable/table_row.qtpl:12
		qw422016.N().S(`
            <div class="usr-table-col" a=`)
		//line views/templates/system/routeTable/table_row.qtpl:13
		qw422016.N().D(i)
		//line views/templates/system/routeTable/table_row.qtpl:13
		qw422016.N().S(`>
                `)
		//line views/templates/system/routeTable/table_row.qtpl:14
		qw422016.E().S(string(col))
		//line views/templates/system/routeTable/table_row.qtpl:14
		qw422016.N().S(`
            </div>
        `)
		//line views/templates/system/routeTable/table_row.qtpl:16
	}
	//line views/templates/system/routeTable/table_row.qtpl:16
	qw422016.N().S(`
        </div>

        <div class="usr-table__filter  usr-table-row">
            <div class="usr-table-col"></div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top  active" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col"></div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>

            <div class="usr-table-col">
                <div class="filt-arrow">
                    <a class="filt-arrow__link arrow-top" href="#"></a>
                    <a class="filt-arrow__link arrow-bottom" href="#"></a>
                </div>
            </div>
        </div>
    </div>

    <div class="usr-table-content">
        <div class="usr-table-content-scroll">
            <div class="usr-table-row-cont">

               `)
	//line views/templates/system/routeTable/table_row.qtpl:79
	for _, row := range rows {
		//line views/templates/system/routeTable/table_row.qtpl:79
		qw422016.N().S(`

                    <div  class="usr-table-row">

                    `)
		//line views/templates/system/routeTable/table_row.qtpl:83
		for i := range columns {
			//line views/templates/system/routeTable/table_row.qtpl:83
			qw422016.N().S(`
                        <div class="usr-table-col">
                            `)
			//line views/templates/system/routeTable/table_row.qtpl:85
			switch r := row[i].(type) {
			//line views/templates/system/routeTable/table_row.qtpl:86
			case time.Time:
				//line views/templates/system/routeTable/table_row.qtpl:86
				qw422016.N().S(`  `)
				//line views/templates/system/routeTable/table_row.qtpl:86
				qw422016.E().S(r.Format("2006-01-02"))
				//line views/templates/system/routeTable/table_row.qtpl:86
				qw422016.N().S(`
                            `)
			//line views/templates/system/routeTable/table_row.qtpl:87
			case float32:
				//line views/templates/system/routeTable/table_row.qtpl:87
				qw422016.N().S(`   `)
				//line views/templates/system/routeTable/table_row.qtpl:87
				qw422016.N().V(r)
				//line views/templates/system/routeTable/table_row.qtpl:87
				qw422016.N().S(`
                            `)
			//line views/templates/system/routeTable/table_row.qtpl:88
			case float64:
				//line views/templates/system/routeTable/table_row.qtpl:88
				qw422016.N().S(`   `)
				//line views/templates/system/routeTable/table_row.qtpl:88
				qw422016.N().FPrec(r, 2)
				//line views/templates/system/routeTable/table_row.qtpl:88
				qw422016.N().S(`
                            `)
			//line views/templates/system/routeTable/table_row.qtpl:89
			default:
				//line views/templates/system/routeTable/table_row.qtpl:89
				qw422016.N().S(`        `)
				//line views/templates/system/routeTable/table_row.qtpl:89
				qw422016.N().V(row[i])
				//line views/templates/system/routeTable/table_row.qtpl:89
				qw422016.N().S(`
                            `)
				//line views/templates/system/routeTable/table_row.qtpl:90
			}
			//line views/templates/system/routeTable/table_row.qtpl:90
			qw422016.N().S(`
                        </div>
                    `)
			//line views/templates/system/routeTable/table_row.qtpl:92
		}
		//line views/templates/system/routeTable/table_row.qtpl:92
		qw422016.N().S(`

                  </div>

                `)
		//line views/templates/system/routeTable/table_row.qtpl:96
	}
	//line views/templates/system/routeTable/table_row.qtpl:96
	qw422016.N().S(`

            </div>
        </div>
    </div>
</div>

`)
//line views/templates/system/routeTable/table_row.qtpl:103
}

//line views/templates/system/routeTable/table_row.qtpl:103
func WriteTableRow(qq422016 qtio422016.Writer, columns []string, rows [][]interface{}) {
	//line views/templates/system/routeTable/table_row.qtpl:103
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/system/routeTable/table_row.qtpl:103
	StreamTableRow(qw422016, columns, rows)
	//line views/templates/system/routeTable/table_row.qtpl:103
	qt422016.ReleaseWriter(qw422016)
//line views/templates/system/routeTable/table_row.qtpl:103
}

//line views/templates/system/routeTable/table_row.qtpl:103
func TableRow(columns []string, rows [][]interface{}) string {
	//line views/templates/system/routeTable/table_row.qtpl:103
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/system/routeTable/table_row.qtpl:103
	WriteTableRow(qb422016, columns, rows)
	//line views/templates/system/routeTable/table_row.qtpl:103
	qs422016 := string(qb422016.B)
	//line views/templates/system/routeTable/table_row.qtpl:103
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/system/routeTable/table_row.qtpl:103
	return qs422016
//line views/templates/system/routeTable/table_row.qtpl:103
}
