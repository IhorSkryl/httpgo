// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apis

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type DefValueCalcFnc = func(ctx *fasthttp.RequestCtx) interface{}

// InParam implement params on request
type InParam struct {
	Name              string
	Desc              string
	Req               bool
	PartReq           []string
	Type              APIRouteParamsType
	DefValue          interface{}
	IncompatibleWiths []string
	TestValue         string
}

func (param *InParam) isPartReq() bool {
	return len(param.PartReq) > 0
}

// defaultValueOfParams return value as default for param, it is only for single required param
func (param *InParam) defaultValueOfParams(ctx *fasthttp.RequestCtx) interface{} {
	switch def := param.DefValue.(type) {
	case DefValueCalcFnc:
		if ctx != nil {
			return def(ctx)
		}
		fnc := runtime.FuncForPC(reflect.ValueOf(def).Pointer())
		fName, len := fnc.FileLine(0)
		return fmt.Sprintf("%s:%d %s()", fName, len, getLastSegment(fnc.Name()))
	default:
		return param.DefValue
	}
}

// inParamToJSON produces a human-friendly description of Apis.
//Based on real data of the executable application, does not require additional documentation.
func inParamToJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	// todo: add description of the test-based return data
	param := (*InParam)(ptr)
	stream.WriteObjectStart()
	defer stream.WriteObjectEnd()

	FirstFieldToJSON(stream, "Descriptor", param.Desc)

	if param.Type != nil {
		if t, ok := param.Type.(jsoniter.ValEncoder); ok {
			stream.WriteMore()
			t.Encode(unsafe.Pointer(&t), stream)
		} else {
			AddFieldToJSON(stream, "Type", param.Type.String())
		}
	}

	if param.Req {
		if len(param.PartReq) > 0 {
			s := strings.Join(param.PartReq, ", ")
			AddFieldToJSON(stream, "Required", "one of {"+s+" and "+param.Name+"} is required")
		} else {
			AddObjectToJSON(stream, "Required", true)
		}
	}

	if param.DefValue != nil {
		AddObjectToJSON(stream, "Default", param.defaultValueOfParams(nil))
	}

	if len(param.IncompatibleWiths) > 0 {
		s := strings.Join(param.IncompatibleWiths, ", ")
		AddFieldToJSON(stream, "IncompatibleWith", "only one of {"+s+" and "+param.Name+"} may use for request")
	}

}

func init() {
	jsoniter.RegisterTypeEncoderFunc("apis.InParam", inParamToJSON, func(pointer unsafe.Pointer) bool {
		return false
	})
}
