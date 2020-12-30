// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/wasm/wasm.proto

package wasm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *PluginSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PluginSource)
	if !ok {
		that2, ok := that.(PluginSource)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetFilters()) != len(target.GetFilters()) {
		return false
	}
	for idx, v := range m.GetFilters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFilters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFilters()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WasmFilter) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WasmFilter)
	if !ok {
		that2, ok := that.(WasmFilter)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFilterStage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFilterStage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFilterStage(), target.GetFilterStage()) {
			return false
		}
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetRootId(), target.GetRootId()) != 0 {
		return false
	}

	if m.GetVmType() != target.GetVmType() {
		return false
	}

	switch m.Src.(type) {

	case *WasmFilter_Image:

		if strings.Compare(m.GetImage(), target.GetImage()) != 0 {
			return false
		}

	case *WasmFilter_FilePath:

		if strings.Compare(m.GetFilePath(), target.GetFilePath()) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *FilterStage) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FilterStage)
	if !ok {
		that2, ok := that.(FilterStage)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetStage() != target.GetStage() {
		return false
	}

	if m.GetPredicate() != target.GetPredicate() {
		return false
	}

	return true
}
