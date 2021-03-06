package worldserver

import (
	"time"

	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/packet"
)

func (s *Session) SendMOTD(motd ...string) {
	pkt := packet.NewWorldPacket(packet.SMSG_MOTD)
	pkt.Write(packet.PutU32(uint32(len(motd))))
	for _, v := range motd {
		pkt.Write(append([]byte(v), 0))
	}
	s.Send(pkt)
}

func (s *Session) SendInitialPackets() {
	dif := packet.NewWorldPacket(packet.MSG_SET_DUNGEON_DIFFICULTY)
	dif.Write(make([]byte, 12))
	s.Send(dif)

	p := packet.NewWorldPacket(packet.SMSG_LOGIN_VERIFY_WORLD)
	p.WriteUint32(s.Char.Map)
	p.WriteFloat32(s.Char.X)
	p.WriteFloat32(s.Char.Y)
	p.WriteFloat32(s.Char.Z)
	p.WriteFloat32(0)
	s.Send(p)

	s.SendDataTimes()

	pk2 := packet.NewWorldPacket(packet.SMSG_FEATURE_SYSTEM_STATUS)
	pk2.Write([]byte{2, 0})
	s.Send(pk2)

	s.SendDanceMoves()

	s.SendMOTD("<-- Gophercraft Core -->", "Welcome. Keep hacking, slaves.")
	log.Println("Sent MOTD")

	s.SendSocialList()

	p = packet.NewWorldPacket(packet.SMSG_BINDPOINTUPDATE)
	p.WriteFloat32(s.Char.X)
	p.WriteFloat32(s.Char.Y)
	p.WriteFloat32(s.Char.Z)
	p.WriteUint32(s.Char.Map)
	p.WriteUint32(s.Char.Zone)
	s.Send(p)

	s.SendPlayerTalentsInfoData()

	p = packet.NewWorldPacket(packet.SMSG_INSTANCE_DIFFICULTY)
	p.WriteUint32(0)
	p.WriteUint32(0)
	s.Send(p)

	// TODO: actually show spells from server database state
	p = packet.NewWorldPacket(packet.SMSG_INITIAL_SPELLS)
	p.Write(Buf_spelllist)
	s.Send(p)

	log.Println("Sent spell list")

	p = packet.NewWorldPacket(packet.SMSG_SEND_UNLEARN_SPELLS)
	p.WriteUint32(0)
	s.Send(p)

	s.SendInitialActionButtons()
	log.Println("Sent action button list")

	s.SendReputations()
	log.Println("Sent initial faction list")

	s.SendAllAcheivementData()

	s.SendEquipmentSetList()

	pkt8 := packet.NewWorldPacket(packet.SMSG_LOGIN_SETTIMESPEED)
	pkt8.WriteUint32(uint32(time.Now().Unix() / 1000))
	pkt8.WriteFloat32(0.01666667)
	pkt8.WriteUint32(0)
	s.Send(pkt8)
	log.Println("Sent world time speed")

	p = packet.NewWorldPacket(packet.SMSG_SET_FORCED_REACTIONS)
	p.WriteUint32(0)
	s.Send(p)

	p = packet.CreateLoginPacket(s.Char.GUID, s.Char.X, s.Char.Y, s.Char.Z, 0)
	s.Send(p)

	p = packet.NewWorldPacket(packet.SMSG_TIME_SYNC_REQ)
	p.WriteUint32(0)
	s.Send(p)

	s.SendDataTimes()
}

func (s *Session) SendDataTimes() {
	p := packet.NewWorldPacket(packet.SMSG_ACCOUNT_DATA_TIMES)
	p.WriteUint32(uint32(time.Now().Unix() / 1000))
	p.Write([]byte{1})
	p.WriteUint32(0x00000015)
	p.WriteUint32(0x0)
	p.WriteUint32(0x0)
	p.WriteUint32(0x0)
	p.WriteUint32(0x0)
	s.Send(p)
}

var spellsPacket = []byte{
	0x00, 0x2C, 0x00, 0x0A, 0x02, 0x00, 0x00, 0x00,
	0x00, 0x75, 0x23, 0x00, 0x00, 0x00, 0x00, 0xA5, 0x23,
	0x00, 0x00, 0x00, 0x00, 0x76, 0x23, 0x00, 0x00,
	0x00, 0x00, 0xFD, 0xEF, 0x00, 0x00, 0x00, 0x00,
	0x67, 0xB3, 0x00, 0x00, 0x00, 0x00, 0xC4, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x45, 0x50, 0x00, 0x00,
	0x00, 0x00, 0xC6, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x46, 0x50, 0x00, 0x00, 0x00, 0x00, 0xC7, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x9D, 0x02, 0x00, 0x00,
	0x00, 0x00, 0x47, 0x50, 0x00, 0x00, 0x00, 0x00,
	0x9E, 0x02, 0x00, 0x00, 0x00, 0x00, 0x2E, 0x0B,
	0x01, 0x00, 0x00, 0x00, 0x48, 0x50, 0x00, 0x00,
	0x00, 0x00, 0x99, 0x09, 0x00, 0x00, 0x00, 0x00,
	0x6B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1A, 0x59,
	0x00, 0x00, 0x00, 0x00, 0xCB, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xCC, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xD7, 0x7D, 0x00, 0x00, 0x00, 0x00, 0xC2, 0x20,
	0x00, 0x00, 0x00, 0x00, 0xBB, 0x1C, 0x00, 0x00,
	0x00, 0x00, 0xCB, 0x19, 0x00, 0x00, 0x00, 0x00,
	0x62, 0x1C, 0x00, 0x00, 0x00, 0x00, 0x25, 0x0D,
	0x00, 0x00, 0x00, 0x00, 0x63, 0x1C, 0x00, 0x00,
	0x00, 0x00, 0x59, 0x18, 0x00, 0x00, 0x00, 0x00,
	0x0B, 0x56, 0x00, 0x00, 0x00, 0x00, 0x93, 0x54,
	0x00, 0x00, 0x00, 0x00, 0x4E, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x94, 0x54, 0x00, 0x00, 0x00, 0x00,
	0x4E, 0x09, 0x00, 0x00, 0x00, 0x00, 0x51, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xAF, 0x09, 0x00, 0x00,
	0x00, 0x00, 0xB5, 0x14, 0x00, 0x00, 0x00, 0x00,
	0x4D, 0x19, 0x00, 0x00, 0x00, 0x00, 0x4E, 0x19,
	0x00, 0x00, 0x00, 0x00, 0x21, 0x22, 0x00, 0x00,
	0x00, 0x00, 0xEA, 0x0B, 0x00, 0x00, 0x00, 0x00,
	0x66, 0x18, 0x00, 0x00, 0x00, 0x00, 0x67, 0x18,
	0x00, 0x00, 0x00, 0x00, 0x9C, 0x23, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00}
