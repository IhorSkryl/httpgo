// This file is automatically generated by qtc from "table_row.css.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line views/templates/system/routeTable/table_row.css.qtpl:1
package routeTable

//line views/templates/system/routeTable/table_row.css.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/templates/system/routeTable/table_row.css.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/templates/system/routeTable/table_row.css.qtpl:1
func StreamTableCSS(qw422016 *qt422016.Writer) {
	//line views/templates/system/routeTable/table_row.css.qtpl:1
	qw422016.N().S(`
.usr-table-row{
  display: flex;
  transition: .3s;
  text-decoration: none !important;
  color: #232b2f;
  &:nth-child(1){
    color: #232bff;
  }

}

.usr-table-col{
  flex:1;
  padding: 1px;
  align-content:  center;
  overflow: hidden;
  &:nth-child(1){
    padding: 3px;
  }
}

.usr-table__t-head{
  .usr-table-col{
    padding: 20px;
    color: #8d8e8e;
    font-size: 12px;
  }
}

.usr-table-content{
  a.usr-table-row{
    &:hover{
      background-color: rgba(120, 147, 182, 0.1);
    }
  }
}


.usr-table{
  font-size: 10px;
  font-weight: 600;
  line-height: 1.8;
  letter-spacing: 0.2px;
  text-align: left;
  color: #232b2f;
  background-color: #ffffff;
  box-shadow: 0 2px 2px 0 rgba(36, 53, 65, 0.1), 0 0 2px 0 rgba(36, 53, 65, 0.1);
  border: solid 1px rgba(229, 229, 229, 0.1);
  border-top:none;
  margin: 0 1px;
}


.usr-table__filter{
  background-color: #f3f3f3;


  .usr-table-col{
    justify-content: center;
    padding: 0 20px;
    font-size: 0;
  }
}


.filt-arrow{
  display: flex;


  .filt-arrow__link.active{
    &:before{
      border-color:  #229be6;
    }

  }
}

.filt-arrow__link{
  width: 24px;
  height: 24px;
  position: relative;
  display: block;
  &.arrow-top:before{
    content: '';
    display: block;
    width: 7px;
    height: 7px;
    border: 1px solid  #4a4a4a;
    border-left: none;
    border-bottom: none;
    transform: rotate(-45deg);
    right: 0;
    left: 0;
    margin: auto;
    top: 6px;
    bottom: 0;
    position: absolute;
    transition: .3s;
  }

  &.arrow-bottom:before{
    content: '';
    display: block;
    width: 7px;
    height: 7px;
    border: 1px solid  #4a4a4a;
    border-left: none;
    border-bottom: none;
    -webkit-transform: rotate(134deg);
    transform: rotate(134deg);
    right: 0;
    left: 0;
    margin: auto;
    top: 0;
    bottom: 2px;
    position: absolute;
    transition: .3s;
  }
}


.avatar-table{
  width: 60px;
  height: 60px;
  border: solid 1px rgba(0, 0, 0, 0.1);
  display: block;
  border-radius:50%;
  position: relative;

  img{
    position: absolute;
    display: block;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
.usr-table-row{
  border-bottom: 1px solid #e5e5e5;
}

.show-more__link{
  font-family: 'FiraSans', Sans-Serif;
  font-size: 13px;
  font-weight: 500;
  line-height: 1.69;
  letter-spacing: 0.5px;
  text-align: left;
  color: #1b70b7;
  height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.show-more{
  background-color: #f3f3f3;
  border-top:44px solid #fff;

}

.pagination__item{
  margin-right: 20px;
  &:last-child{
    margin: 0;
  }
}

.usr-table.usr-table--003{
  .usr-table-col{
    text-align: center;
    justify-content: center;
    &:nth-child(1){
      text-align: left;
      justify-content: flex-start;
    }
  }
}

.filt-arrow__link:hover:before{
  border-width:2px;
}

.usr-table-col {
  word-break: break-word;
  word-break: break-all;
}
.c-black{
  color: #000;
}
.c-green{
  color: #7bcc1f;
}
.c-red{
  color: #d0021b;
}
.c-orange {
  color: #f1951c;
}

.usr-table-content{
  overflow-y: auto;
  width: 100%;
}

.usr-table-content-scroll {
  min-height: 40px;
}


.usr-table{
  max-height: calc(100vh - 70px);
  display: flex;
  flex-direction: column;
  border-bottom-width: 15px;
  border-bottom-color: #f3f3f3;
}

`)
//line views/templates/system/routeTable/table_row.css.qtpl:222
}

//line views/templates/system/routeTable/table_row.css.qtpl:222
func WriteTableCSS(qq422016 qtio422016.Writer) {
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	StreamTableCSS(qw422016)
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	qt422016.ReleaseWriter(qw422016)
//line views/templates/system/routeTable/table_row.css.qtpl:222
}

//line views/templates/system/routeTable/table_row.css.qtpl:222
func TableCSS() string {
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	qb422016 := qt422016.AcquireByteBuffer()
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	WriteTableCSS(qb422016)
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	qs422016 := string(qb422016.B)
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	qt422016.ReleaseByteBuffer(qb422016)
	//line views/templates/system/routeTable/table_row.css.qtpl:222
	return qs422016
//line views/templates/system/routeTable/table_row.css.qtpl:222
}