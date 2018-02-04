package worldserver

import "github.com/superp00t/gophercraft/packet"

func (s *Session) SendSocialList() {
	p := packet.NewWorldPacket(packet.SMSG_CONTACT_LIST)
	Buf_sociallist := []byte{0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	p.WriteBytes(Buf_sociallist)
	s.Send(p)
}

func (s *Session) SendDanceMoves() {
	p := packet.NewWorldPacket(packet.SMSG_LEARNED_DANCE_MOVES)
	Buf_dance := []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	p.WriteBytes(Buf_dance)
	s.Send(p)
}
