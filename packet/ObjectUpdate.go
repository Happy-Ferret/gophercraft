package packet

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/guid"
)

//go:generate stringer -type=UpdateType

type UpdateType uint8

const (
	UPDATETYPE_VALUES UpdateType = iota
	UPDATETYPE_MOVEMENT
	UPDATETYPE_CREATE_OBJECT
	UPDATETYPE_CREATE_OBJECT2
	UPDATETYPE_OUT_OF_RANGE_OBJECTS
	UPDATETYPE_NEAR_OBJECTS

	TYPEID_OBJECT        = 0
	TYPEID_ITEM          = 1
	TYPEID_CONTAINER     = 2
	TYPEID_UNIT          = 3
	TYPEID_PLAYER        = 4
	TYPEID_GAMEOBJECT    = 5
	TYPEID_DYNAMICOBJECT = 6
	TYPEID_CORPSE        = 7
	TYPEID_AREATRIGGER   = 8
	TYPEID_SCENEOBJECT   = 9
	TYPEID_CONVERSATION  = 10
)

type ObjectUpdate struct {
	Opcode     WorldType
	BlockCount uint32
	Upd        []UpdateBlock
}

type UpdateBlock struct {
	Type    UpdateType
	GUID    guid.GUID
	ObjType uint8
	Flags   uint16

	MovementUpdate *MovementUpdate

	ValuesLength uint8
	Values       []byte
}

type MoveSpline struct {
	Flags      uint32
	Angle      float32
	TimePassed uint32
	Duration   uint32
	ID         uint32

	DurationMod,
	DurationModNext float32

	VerticalAcceleration float32
	EffectStartTime      uint32

	NodeSize uint32
	Nodes    []float32
	Mode     uint8

	ZeroVector []float32
}

func DecodeSplineData(b *EtcBuffer) *MoveSpline {
	m := &MoveSpline{}
	m.Flags = b.ReadUint32()
	m.Angle = b.ReadFloat32()
	m.TimePassed = b.ReadUint32()
	m.Duration = b.ReadUint32()
	m.ID = b.ReadUint32()
	m.DurationMod = b.ReadFloat32()
	m.DurationModNext = b.ReadFloat32()
	m.VerticalAcceleration = b.ReadFloat32()
	m.EffectStartTime = b.ReadUint32()
	m.NodeSize = b.ReadUint32()
	log.Fatal(m.NodeSize)
	e := int(m.NodeSize)
	m.Nodes = make([]float32, e)
	for i := 0; i < e; i++ {
		m.Nodes[i] = b.ReadFloat32()
	}
	m.Mode = b.ReadByte()
	m.ZeroVector = make([]float32, 3)
	for i := 0; i < 3; i++ {
		m.ZeroVector[i] = b.ReadFloat32()
	}
	return m
}

type MovementData struct {
	Flags      uint32
	ExtraFlags uint16
	Time       uint32
	Position   *PackedXYZ

	// Optional transport
	TransportGUID     guid.GUID
	TransportPosition *PackedXYZ
	TransportTime     uint32
	TransportSeat     uint8
	TransportTime2    uint32

	Pitch    float32
	FallTime uint32

	JumpZSpeed, JumpSin, JumpCos, JumpXYSpeed float32

	SplineElevation float32
}

type MovementUpdate struct {
	Flags uint16

	MovePacket *MovementData
	Speeds     []float32
	SplineData *MoveSpline

	TransportGUID                   guid.GUID
	PositionX, PositionY, PositionZ float32
	OffsetX, OffsetY, OffsetZ       float32
	Orientation                     float32

	StatX, StatY, StatZ, StatO float32

	LowGUID uint32

	TargetVictim guid.GUID

	TransportProgress uint32

	MountID uint32
	MountO  float32

	Rotation uint64
}

func DecodeMovementUpdate(b *EtcBuffer) *MovementUpdate {
	m := &MovementUpdate{}
	m.Flags = b.ReadUint16()
	if (m.Flags & UPDATEFLAG_LIVING) != 0 {
		m.Speeds = make([]float32, 9)
		m.MovePacket = DecodeMovementData(b)
		for i := 0; i < 9; i++ {
			m.Speeds[i] = b.ReadFloat32()
		}
		if (m.MovePacket.Flags & MOVEMENTFLAG_SPLINE_ENABLED) != 0 {
			m.SplineData = DecodeSplineData(b)
		}
	} else {
		if (m.Flags & UPDATEFLAG_POSITION) != 0 {
			m.TransportGUID = b.ReadPackedGUID()
			m.PositionX = b.ReadFloat32()
			m.PositionY = b.ReadFloat32()
			m.PositionZ = b.ReadFloat32()

			m.OffsetX = b.ReadFloat32()
			m.OffsetY = b.ReadFloat32()
			m.OffsetZ = b.ReadFloat32()

			m.Orientation = b.ReadFloat32()
		} else {
			if (m.Flags & UPDATEFLAG_STATIONARY_POSITION) != 0 {
				m.StatX = b.ReadFloat32()
				m.StatY = b.ReadFloat32()
				m.StatZ = b.ReadFloat32()
				m.StatO = b.ReadFloat32()
			}
		}
	}

	if (m.Flags & UPDATEFLAG_UNKNOWN) != 0 {
		b.ReadUint32()
	}

	if (m.Flags & UPDATEFLAG_LOWGUID) != 0 {
		m.LowGUID = b.ReadUint32()
	}

	if (m.Flags & UPDATEFLAG_HAS_TARGET) != 0 {
		m.TargetVictim = b.ReadPackedGUID()
	}

	if (m.Flags & UPDATEFLAG_TRANSPORT) != 0 {
		m.TransportProgress = b.ReadUint32()
	}

	if (m.Flags & UPDATEFLAG_VEHICLE) != 0 {
		m.MountID = b.ReadUint32()
		m.MountO = b.ReadFloat32()
	}

	if (m.Flags & UPDATEFLAG_ROTATION) != 0 {
		m.Rotation = b.ReadUint64()
	}

	return m
}

func DecodeMovementData(b *EtcBuffer) *MovementData {
	m := &MovementData{}
	m.Flags = b.ReadUint32()
	m.ExtraFlags = b.ReadUint16()
	m.Time = b.ReadUint32()
	m.Position = b.ReadPackedXYZ()

	if (m.Flags & MOVEMENTFLAG_ONTRANSPORT) != 0 {
		m.TransportGUID = b.ReadPackedGUID()
		m.TransportPosition = b.ReadPackedXYZ()
		m.TransportTime = b.ReadUint32()
		m.TransportSeat = b.ReadByte()

		if (m.ExtraFlags & uint16(MOVEMENTFLAG2_INTERPOLATED_MOVEMENT)) != 0 {
			m.TransportTime2 = b.ReadUint32()
		}
	}

	if (m.Flags&(MOVEMENTFLAG_SWIMMING|MOVEMENTFLAG_FLYING)) != 0 || (m.ExtraFlags&uint16(MOVEMENTFLAG2_ALWAYS_ALLOW_PITCHING)) != 0 {
		m.Pitch = b.ReadFloat32()
	}

	if (m.Flags & MOVEMENTFLAG_FALLING) != 0 {
		m.JumpZSpeed = b.ReadFloat32()
		m.JumpSin = b.ReadFloat32()
		m.JumpCos = b.ReadFloat32()
		m.JumpXYSpeed = b.ReadFloat32()
	}

	if (m.Flags & MOVEMENTFLAG_SPLINE_ELEVATION) != 0 {
		m.SplineElevation = b.ReadFloat32()
	}

	return m
}

func CreateLoginPacket(g guid.GUID, x, y, z, o float32) *WorldPacket {
	p := NewWorldPacket(SMSG_UPDATE_OBJECT)
	p.WriteUint32(1)

	p.WriteByte(uint8(UPDATETYPE_CREATE_OBJECT))
	p.WritePackedGUID(g)
	p.WriteByte(TYPEID_PLAYER)

	flags := uint32(UPDATEFLAG_SELF | UPDATEFLAG_POSITION)
	p.WriteUint32(flags)
	p.WriteUint16(0)

	p.WriteByte(0)

	p.WriteFloat32(x)
	p.WriteFloat32(y)
	p.WriteFloat32(z)

	p.WriteFloat32(x)
	p.WriteFloat32(y)
	p.WriteFloat32(z)

	p.WriteFloat32(o)
	p.WriteFloat32(0)

	return p
}

func UnmarshalObjectUpdate(in []byte) (*ObjectUpdate, error) {
	op := WorldType(binary.LittleEndian.Uint16(in[2:4]))
	var b []byte
	// Decompress if necessary
	if op == SMSG_COMPRESSED_UPDATE_OBJECT {
		lng := binary.LittleEndian.Uint32(in[4:8])
		glogger.Println("Compressed length", lng)
		b = uncompress(in[8:])
	} else {
		if op == SMSG_UPDATE_OBJECT {
			b = in[4:]
		} else {
			return nil, fmt.Errorf("Invalid opcode")
		}
	}

	o := &ObjectUpdate{}
	o.Opcode = op
	bf := NewEtcBuffer(b)

	o.BlockCount = bf.ReadUint32()
	for i := 0; i < int(o.BlockCount); i++ {
		u := UpdateBlock{}
		u.Type = UpdateType(bf.ReadByte())
		pkGUID := bf.ReadPackedGUID()
		u.GUID = pkGUID
		u.ObjType = bf.ReadByte()
		u.MovementUpdate = DecodeMovementUpdate(bf)
		u.ValuesLength = bf.ReadByte()
		u.Values = bf.ReadBytes(int(u.ValuesLength))
		o.Upd = append(o.Upd, u)
	}

	return o, nil
}
