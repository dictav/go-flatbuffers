// automatically generated, do not modify

package Sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Weapon struct {
	_tab flatbuffers.Table
}

func (rcv *Weapon) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Weapon) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Weapon) Damage() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func WeaponStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func WeaponAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0) }
func WeaponAddDamage(builder *flatbuffers.Builder, damage int16) { builder.PrependInt16Slot(1, damage, 0) }
func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
