// This file is automatically generated by qtc from "rooms.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line rooms.qtpl:1
package pages

//line rooms.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line rooms.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line rooms.qtpl:2
type DiscountTermRules struct {
	Main_place_only int
	Value           string
	Condition       int
	Night_count     int
}

type DiscountRules struct {
	Count_use_for int
	Value         string
}

type DiscountData struct {
	Rules                 []DiscountRules
	ChildRules            DiscountRules
	DiscountLongTermRules DiscountTermRules
}

type DiscountAllData struct {
	Min_points_count int
	Max_points_count int
	DiscountMain     DiscountData
	DiscountChildren DiscountData
	DiscountTerm     DiscountData
}

// JSON marshaling

//line rooms.qtpl:31
func (d *DiscountAllData) StreamJSON(qw422016 *qt422016.Writer) {
	//line rooms.qtpl:31
	qw422016.N().S(`{"min_points_count": "`)
	//line rooms.qtpl:33
	qw422016.N().D(d.Min_points_count)
	//line rooms.qtpl:33
	qw422016.N().S(`","max_points_count": "`)
	//line rooms.qtpl:34
	qw422016.N().D(d.Max_points_count)
	//line rooms.qtpl:34
	qw422016.N().S(`",`)
	//line rooms.qtpl:35
	if len(d.DiscountMain.Rules) > 0 {
		//line rooms.qtpl:35
		qw422016.N().S(`"discountMain":{"rules":[`)
		//line rooms.qtpl:39
		for j, s := range d.DiscountMain.Rules {
			//line rooms.qtpl:39
			qw422016.N().S(`{"value":`)
			//line rooms.qtpl:41
			qw422016.N().Q(s.Value)
			//line rooms.qtpl:41
			qw422016.N().S(`,"count_use_for": "`)
			//line rooms.qtpl:42
			qw422016.N().D(s.Count_use_for)
			//line rooms.qtpl:42
			qw422016.N().S(`"}`)
			//line rooms.qtpl:44
			if j+1 < len(d.DiscountMain.Rules) {
				//line rooms.qtpl:44
				qw422016.N().S(`,`)
				//line rooms.qtpl:44
			}
			//line rooms.qtpl:45
		}
		//line rooms.qtpl:45
		qw422016.N().S(`]}`)
		//line rooms.qtpl:49
		if (d.DiscountChildren.ChildRules.Value > "" && d.DiscountChildren.ChildRules.Count_use_for > 0) || (d.DiscountTerm.DiscountLongTermRules.Main_place_only > 0 && d.DiscountTerm.DiscountLongTermRules.Value > "" && d.DiscountTerm.DiscountLongTermRules.Condition > 0 && d.DiscountTerm.DiscountLongTermRules.Night_count > 0) {
			//line rooms.qtpl:49
			qw422016.N().S(`,`)
			//line rooms.qtpl:49
		}
		//line rooms.qtpl:50
	}
	//line rooms.qtpl:52
	if d.DiscountChildren.ChildRules.Value > "" && d.DiscountChildren.ChildRules.Count_use_for > 0 {
		//line rooms.qtpl:52
		qw422016.N().S(`"discountChildren":{"rules":{"value":`)
		//line rooms.qtpl:56
		qw422016.N().Q(d.DiscountChildren.ChildRules.Value)
		//line rooms.qtpl:56
		qw422016.N().S(`,"count_use_for":"`)
		//line rooms.qtpl:57
		qw422016.N().D(d.DiscountChildren.ChildRules.Count_use_for)
		//line rooms.qtpl:57
		qw422016.N().S(`"}}`)
		//line rooms.qtpl:62
		if d.DiscountTerm.DiscountLongTermRules.Main_place_only > 0 && d.DiscountTerm.DiscountLongTermRules.Value > "" && d.DiscountTerm.DiscountLongTermRules.Condition > 0 && d.DiscountTerm.DiscountLongTermRules.Night_count > 0 {
			//line rooms.qtpl:62
			qw422016.N().S(`,`)
			//line rooms.qtpl:62
		}
		//line rooms.qtpl:63
	}
	//line rooms.qtpl:65
	if d.DiscountTerm.DiscountLongTermRules.Main_place_only > 0 && d.DiscountTerm.DiscountLongTermRules.Value > "" && d.DiscountTerm.DiscountLongTermRules.Condition > 0 && d.DiscountTerm.DiscountLongTermRules.Night_count > 0 {
		//line rooms.qtpl:65
		qw422016.N().S(`"discountLongTerm":{"rules":{"main_place_only":"`)
		//line rooms.qtpl:69
		qw422016.N().D(d.DiscountTerm.DiscountLongTermRules.Main_place_only)
		//line rooms.qtpl:69
		qw422016.N().S(`","value":`)
		//line rooms.qtpl:70
		qw422016.N().Q(d.DiscountTerm.DiscountLongTermRules.Value)
		//line rooms.qtpl:70
		qw422016.N().S(`,"condition":"`)
		//line rooms.qtpl:71
		qw422016.N().D(d.DiscountTerm.DiscountLongTermRules.Condition)
		//line rooms.qtpl:71
		qw422016.N().S(`","night_count":"`)
		//line rooms.qtpl:72
		qw422016.N().D(d.DiscountTerm.DiscountLongTermRules.Night_count)
		//line rooms.qtpl:72
		qw422016.N().S(`"}}`)
		//line rooms.qtpl:75
	}
	//line rooms.qtpl:75
	qw422016.N().S(`}`)
//line rooms.qtpl:82
}

//line rooms.qtpl:82
func (d *DiscountAllData) WriteJSON(qq422016 qtio422016.Writer) {
	//line rooms.qtpl:82
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line rooms.qtpl:82
	d.StreamJSON(qw422016)
	//line rooms.qtpl:82
	qt422016.ReleaseWriter(qw422016)
//line rooms.qtpl:82
}

//line rooms.qtpl:82
func (d *DiscountAllData) JSON() string {
	//line rooms.qtpl:82
	qb422016 := qt422016.AcquireByteBuffer()
	//line rooms.qtpl:82
	d.WriteJSON(qb422016)
	//line rooms.qtpl:82
	qs422016 := string(qb422016.B)
	//line rooms.qtpl:82
	qt422016.ReleaseByteBuffer(qb422016)
	//line rooms.qtpl:82
	return qs422016
//line rooms.qtpl:82
}

//line rooms.qtpl:83
func StreamPutRoomsHead(qw422016 *qt422016.Writer) {
	//line rooms.qtpl:83
	qw422016.N().S(`<div class="main-form-wrap">`)
//line rooms.qtpl:85
}

//line rooms.qtpl:85
func WritePutRoomsHead(qq422016 qtio422016.Writer) {
	//line rooms.qtpl:85
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line rooms.qtpl:85
	StreamPutRoomsHead(qw422016)
	//line rooms.qtpl:85
	qt422016.ReleaseWriter(qw422016)
//line rooms.qtpl:85
}

//line rooms.qtpl:85
func PutRoomsHead() string {
	//line rooms.qtpl:85
	qb422016 := qt422016.AcquireByteBuffer()
	//line rooms.qtpl:85
	WritePutRoomsHead(qb422016)
	//line rooms.qtpl:85
	qs422016 := string(qb422016.B)
	//line rooms.qtpl:85
	qt422016.ReleaseByteBuffer(qb422016)
	//line rooms.qtpl:85
	return qs422016
//line rooms.qtpl:85
}

//line rooms.qtpl:87
func StreamPutRoomsEnd(qw422016 *qt422016.Writer) {
	//line rooms.qtpl:87
	qw422016.N().S(`</div>`)
//line rooms.qtpl:89
}

//line rooms.qtpl:89
func WritePutRoomsEnd(qq422016 qtio422016.Writer) {
	//line rooms.qtpl:89
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line rooms.qtpl:89
	StreamPutRoomsEnd(qw422016)
	//line rooms.qtpl:89
	qt422016.ReleaseWriter(qw422016)
//line rooms.qtpl:89
}

//line rooms.qtpl:89
func PutRoomsEnd() string {
	//line rooms.qtpl:89
	qb422016 := qt422016.AcquireByteBuffer()
	//line rooms.qtpl:89
	WritePutRoomsEnd(qb422016)
	//line rooms.qtpl:89
	qs422016 := string(qb422016.B)
	//line rooms.qtpl:89
	qt422016.ReleaseByteBuffer(qb422016)
	//line rooms.qtpl:89
	return qs422016
//line rooms.qtpl:89
}
