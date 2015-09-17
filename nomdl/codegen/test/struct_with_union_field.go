// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// StructWithUnionField

type StructWithUnionFieldDef struct {
	A            float32
	__unionIndex uint32
	__unionValue interface{}
}

type StructWithUnionField struct {
	l types.List
}

func NewStructWithUnionField() StructWithUnionField {
	return StructWithUnionField{types.NewList(
		types.Float32(0),
		types.UInt32(0),
		types.Float64(0),
	)}
}

func (def StructWithUnionFieldDef) New() StructWithUnionField {
	return StructWithUnionField{
		types.NewList(
			types.Float32(def.A),
			types.UInt32(def.__unionIndex),
			def.__unionDefToValue(),
		)}
}

func (self StructWithUnionField) Def() StructWithUnionFieldDef {
	return StructWithUnionFieldDef{
		float32(self.l.Get(0).(types.Float32)),
		uint32(self.l.Get(1).(types.UInt32)),
		self.__unionValueToDef(),
	}
}

func (def StructWithUnionFieldDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	case 2:
		return def.__unionValue.(types.Blob)
	case 3:
		return def.__unionValue.(types.Value)
	case 4:
		return def.__unionValue.(SetOfUInt8Def).New().NomsValue()
	}
	panic("unreachable")
}

func (self StructWithUnionField) __unionValueToDef() interface{} {
	switch uint32(self.l.Get(1).(types.UInt32)) {
	case 0:
		return float64(self.l.Get(2).(types.Float64))
	case 1:
		return self.l.Get(2).(types.String).String()
	case 2:
		return self.l.Get(2).(types.Blob)
	case 3:
		return self.l.Get(2)
	case 4:
		return SetOfUInt8FromVal(self.l.Get(2)).Def()
	}
	panic("unreachable")
}

func StructWithUnionFieldFromVal(val types.Value) StructWithUnionField {
	// TODO: Validate here
	return StructWithUnionField{val.(types.List)}
}

func (self StructWithUnionField) NomsValue() types.Value {
	return self.l
}

func (self StructWithUnionField) Equals(p StructWithUnionField) bool {
	return self.l.Equals(p.l)
}

func (self StructWithUnionField) Ref() ref.Ref {
	return self.l.Ref()
}

func (self StructWithUnionField) A() float32 {
	return float32(self.l.Get(0).(types.Float32))
}

func (self StructWithUnionField) SetA(val float32) StructWithUnionField {
	return StructWithUnionField{self.l.Set(0, types.Float32(val))}
}

func (self StructWithUnionField) B() (val float64, ok bool) {
	if self.l.Get(1).(types.UInt32) != 0 {
		return
	}
	return float64(self.l.Get(2).(types.Float64)), true
}

func (self StructWithUnionField) SetB(val float64) StructWithUnionField {
	return StructWithUnionField{self.l.Set(1, types.UInt32(0)).Set(2, types.Float64(val))}
}

func (def StructWithUnionFieldDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def StructWithUnionFieldDef) SetB(val float64) StructWithUnionFieldDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (self StructWithUnionField) C() (val string, ok bool) {
	if self.l.Get(1).(types.UInt32) != 1 {
		return
	}
	return self.l.Get(2).(types.String).String(), true
}

func (self StructWithUnionField) SetC(val string) StructWithUnionField {
	return StructWithUnionField{self.l.Set(1, types.UInt32(1)).Set(2, types.NewString(val))}
}

func (def StructWithUnionFieldDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def StructWithUnionFieldDef) SetC(val string) StructWithUnionFieldDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}

func (self StructWithUnionField) D() (val types.Blob, ok bool) {
	if self.l.Get(1).(types.UInt32) != 2 {
		return
	}
	return self.l.Get(2).(types.Blob), true
}

func (self StructWithUnionField) SetD(val types.Blob) StructWithUnionField {
	return StructWithUnionField{self.l.Set(1, types.UInt32(2)).Set(2, val)}
}

func (def StructWithUnionFieldDef) D() (val types.Blob, ok bool) {
	if def.__unionIndex != 2 {
		return
	}
	return def.__unionValue.(types.Blob), true
}

func (def StructWithUnionFieldDef) SetD(val types.Blob) StructWithUnionFieldDef {
	def.__unionIndex = 2
	def.__unionValue = val
	return def
}

func (self StructWithUnionField) E() (val types.Value, ok bool) {
	if self.l.Get(1).(types.UInt32) != 3 {
		return
	}
	return self.l.Get(2), true
}

func (self StructWithUnionField) SetE(val types.Value) StructWithUnionField {
	return StructWithUnionField{self.l.Set(1, types.UInt32(3)).Set(2, val)}
}

func (def StructWithUnionFieldDef) E() (val types.Value, ok bool) {
	if def.__unionIndex != 3 {
		return
	}
	return def.__unionValue.(types.Value), true
}

func (def StructWithUnionFieldDef) SetE(val types.Value) StructWithUnionFieldDef {
	def.__unionIndex = 3
	def.__unionValue = val
	return def
}

func (self StructWithUnionField) F() (val SetOfUInt8, ok bool) {
	if self.l.Get(1).(types.UInt32) != 4 {
		return
	}
	return SetOfUInt8FromVal(self.l.Get(2)), true
}

func (self StructWithUnionField) SetF(val SetOfUInt8) StructWithUnionField {
	return StructWithUnionField{self.l.Set(1, types.UInt32(4)).Set(2, val.NomsValue())}
}

func (def StructWithUnionFieldDef) F() (val SetOfUInt8Def, ok bool) {
	if def.__unionIndex != 4 {
		return
	}
	return def.__unionValue.(SetOfUInt8Def), true
}

func (def StructWithUnionFieldDef) SetF(val SetOfUInt8Def) StructWithUnionFieldDef {
	def.__unionIndex = 4
	def.__unionValue = val
	return def
}

// SetOfUInt8

type SetOfUInt8 struct {
	s types.Set
}

type SetOfUInt8Def map[uint8]bool

func NewSetOfUInt8() SetOfUInt8 {
	return SetOfUInt8{types.NewSet()}
}

func (def SetOfUInt8Def) New() SetOfUInt8 {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.UInt8(d)
		i++
	}
	return SetOfUInt8{types.NewSet(l...)}
}

func (s SetOfUInt8) Def() SetOfUInt8Def {
	def := make(map[uint8]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[uint8(v.(types.UInt8))] = true
		return false
	})
	return def
}

func SetOfUInt8FromVal(p types.Value) SetOfUInt8 {
	return SetOfUInt8{p.(types.Set)}
}

func (s SetOfUInt8) NomsValue() types.Value {
	return s.s
}

func (s SetOfUInt8) Equals(p SetOfUInt8) bool {
	return s.s.Equals(p.s)
}

func (s SetOfUInt8) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfUInt8) Empty() bool {
	return s.s.Empty()
}

func (s SetOfUInt8) Len() uint64 {
	return s.s.Len()
}

func (s SetOfUInt8) Has(p uint8) bool {
	return s.s.Has(types.UInt8(p))
}

type SetOfUInt8IterCallback func(p uint8) (stop bool)

func (s SetOfUInt8) Iter(cb SetOfUInt8IterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8IterAllCallback func(p uint8)

func (s SetOfUInt8) IterAll(cb SetOfUInt8IterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8FilterCallback func(p uint8) (keep bool)

func (s SetOfUInt8) Filter(cb SetOfUInt8FilterCallback) SetOfUInt8 {
	ns := NewSetOfUInt8()
	s.IterAll(func(v uint8) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfUInt8) Insert(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfUInt8) Remove(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfUInt8) Union(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfUInt8) Subtract(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfUInt8) Any() uint8 {
	return uint8(s.s.Any().(types.UInt8))
}

func (s SetOfUInt8) fromStructSlice(p []SetOfUInt8) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfUInt8) fromElemSlice(p []uint8) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.UInt8(v)
	}
	return r
}
