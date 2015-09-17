// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// EnumStruct

type EnumStructDef struct {
	Hand Handedness
}

type EnumStruct struct {
	l types.List
}

func NewEnumStruct() EnumStruct {
	return EnumStruct{types.NewList(
		types.Int32(0),
	)}
}

func (def EnumStructDef) New() EnumStruct {
	return EnumStruct{
		types.NewList(
			types.Int32(def.Hand),
		)}
}

func (self EnumStruct) Def() EnumStructDef {
	return EnumStructDef{
		Handedness(self.l.Get(0).(types.Int32)),
	}
}

func EnumStructFromVal(val types.Value) EnumStruct {
	// TODO: Validate here
	return EnumStruct{val.(types.List)}
}

func (self EnumStruct) NomsValue() types.Value {
	return self.l
}

func (self EnumStruct) Equals(p EnumStruct) bool {
	return self.l.Equals(p.l)
}

func (self EnumStruct) Ref() ref.Ref {
	return self.l.Ref()
}

func (self EnumStruct) Hand() Handedness {
	return Handedness(self.l.Get(0).(types.Int32))
}

func (self EnumStruct) SetHand(val Handedness) EnumStruct {
	return EnumStruct{self.l.Set(0, types.Int32(val))}
}

// Handedness

type Handedness uint32

const (
	Right Handedness = iota
	Left
	Switch
)