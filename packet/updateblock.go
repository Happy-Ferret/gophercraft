package packet

// import "github.com/superp00t/gophercraft/guid"

// const (

// 	TYPEID_OBJECT        uint8 = 0
// 	TYPEID_ITEM          uint8 = 1
// 	TYPEID_CONTAINER     uint8 = 2
// 	TYPEID_UNIT          uint8 = 3
// 	TYPEID_PLAYER        uint8 = 4
// 	TYPEID_GAMEOBJECT    uint8 = 5
// 	TYPEID_DYNAMICOBJECT uint8 = 6
// 	TYPEID_CORPSE        uint8 = 7
// 	TYPEID_AREATRIGGER   uint8 = 8
// 	TYPEID_SCENEOBJECT   uint8 = 9
// 	TYPEID_CONVERSATION  uint8 = 10

const (
	UPDATEFLAG_NONE                uint16 = 0x0000
	UPDATEFLAG_SELF                uint16 = 0x0001
	UPDATEFLAG_TRANSPORT           uint16 = 0x0002
	UPDATEFLAG_HAS_TARGET          uint16 = 0x0004
	UPDATEFLAG_UNKNOWN             uint16 = 0x0008
	UPDATEFLAG_LOWGUID             uint16 = 0x0010
	UPDATEFLAG_LIVING              uint16 = 0x0020
	UPDATEFLAG_STATIONARY_POSITION uint16 = 0x0040
	UPDATEFLAG_VEHICLE             uint16 = 0x0080
	UPDATEFLAG_POSITION            uint16 = 0x0100
	UPDATEFLAG_ROTATION            uint16 = 0x0200
)

// 	TYPEMASK_OBJECT        uint16 = 0x0001
// 	TYPEMASK_ITEM          uint16 = 0x0002
// 	TYPEMASK_CONTAINER     uint16 = 0x0004
// 	TYPEMASK_UNIT          uint16 = 0x0008
// 	TYPEMASK_PLAYER        uint16 = 0x0010
// 	TYPEMASK_GAMEOBJECT    uint16 = 0x0020
// 	TYPEMASK_DYNAMICOBJECT uint16 = 0x0040
// 	TYPEMASK_CORPSE        uint16 = 0x0080
// 	TYPEMASK_AREATRIGGER   uint16 = 0x0100
// 	TYPEMASK_SCENEOBJECT   uint16 = 0x0200
// 	TYPEMASK_CONVERSATION  uint16 = 0x0400
// 	TYPEMASK_SEER          uint16 = TYPEMASK_PLAYER | TYPEMASK_UNIT | TYPEMASK_DYNAMICOBJECT
// )

// type UpdateBlock struct {
// 	MapID uint16
// }

// type ObjectUpdate struct {
// 	UpdateFlags uint16
// 	data        *EtcBuffer

// 	Obj  *Object
// }

// func (o *ObjectUpdate) IsType(mask uint16) bool {
// 	return (mask & o.ObjectType) != 0
// }

// func (o *ObjectUpdate) Encode(me bool) []byte {
// 	flags := o.UpdateFlags
// 	uType := UPDATETYPE_CREATE_OBJECT2

// 	if me {
// 		flags |= UPDATEFLAG_SELF
// 	}

// 	if (flags & UPDATEFLAG_STATIONARY_POSITION) != 0 {
// 		if o.IsType(TYPEMASK_DYNAMICOBJECT | TYPEMASK_CORPSE | TYPEMASK_PLAYER) {
// 			uType = UPDATETYPE_CREATE_OBJECT2
// 		}
// 	}

// 	o.data.WriteByte(uType)
// 	o.data.WriteBytes(o.ObjectGUID.EncodePacked())
// 	o.data.WriteByte(guid.TYPEID_PLAYER)
// }

// type Unit struct {}

// type Object struct {
// 	GUID    guid.GUID
// 	Self    bool
// 	X, Y, Z, O float32
// }

// func (o *Object) GetTypeID() uint8 {
// 	switch o.GUID.High() {
// 		case guid.Player:
// 		return TYPEID_PLAYER:
// 		case guid.Pet, guid.Unit:
// 		return TYPEID_UNIT:
// 		default:
// 		return 0
// 	}
// }

// func (o *Object) BuildMovementUpdate(data *EtcBuffer, flags uint16) {
// 	// var unit *Unit =

// 	data.WriteUint16(flags)

// 	if (flags & UPDATEFLAG_LIVING) != 0 {
// 		o.BuildMovementPacket(data) // Unit

// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)
// 		data.WriteFloat32(1.0)

// 		// do splines
// 		// Object.cpp#347
// 	} else {
// 		if (flags & UPDATEFLAG_POSITION) != 0 {
// 			// do transports Object.cpp#355

// 			data.WriteByte(0)
// 			data.WriteFloat32(o.X)
// 			data.WriteFloat32(o.Y)
// 			data.WriteFloat32(o.Z)

// 			data.WriteFloat32(o.X)
// 			data.WriteFloat32(o.Y)
// 			data.WriteFloat32(o.Z)

// 			data.WriteFloat32(o.O)

// 			if o.GetTypeID() == TYPEID_CORPSE {
// 				data.WriteFloat32(o.O)
// 			} else {
// 				data.WriteFloat32(0)
// 			}
// 		}
// 	}

// 	if (flags & UPDATEFLAG_UNKNOWN) != 0 {
// 		data.WriteUint32(0)
// 	}

// 	if (flags & UPDATEFLAG_LOWGUID) != 0 {
// 		switch o.GetTypeID() {
// 			case TYPEID_OBJECT, TYPEID_ITEM, TYPEID_CONTAINER, TYPEID_GAMEOBJECT, TYPEID_DYNAMICOBJECT,TYPEID_CORPSE:
// 			data.WriteUint32(o.GUID.GetCounter())
// 			case TYPEID_UNIT:
// 			data.WriteUint32(0x0000000B)
// 			case TYPEID_PLAYER:
// 			if (flags & UPDATEFLAG_SELF) != 0 {
// 				data.WriteUint32(0x0000002F)
// 			} else {
// 				data.WriteUint32(0x0000000B)
// 			}
// 			default:
// 			data.WriteUint32(0x00000000)
// 		}
// 	}

// 	if (flags & UPDATEFLAG_HAS_TARGET) != 0 {
// 	   /* ASSERT(unit);
//         if (Unit* victim = unit->GetVictim())
//             *data << victim->GetPackGUID();
//         else
// 			*data << uint8(0); */
// 		data.WriteByte(0)
// 	}

// 	if (flags & UPDATEFLAG_TRANSPORT) != 0 {
// 		/* GameObject const* go = ToGameObject();
//          @TODO Use IsTransport() to also handle type 11 (TRANSPORT)
//             Currently grid objects are not updated if there are no nearby players,
//             this causes clients to receive different PathProgress
//             resulting in players seeing the object in a different position

//         if (go && go->ToTransport())
//             *data << uint32(go->GetGOValue()->Transport.PathProgress);
//         else
// 			*data << uint32(GameTime::GetGameTimeMS())*/

// 		data.WriteUint32(uint32(time.Now().Unix() / 1000))
// 	}

// }
