package worldserver

import (
	"log"

	"github.com/superp00t/gophercraft/packet"
)

func (s *Session) IntroductoryPackets() {
	if len(s.AddonData) != 1 {
		s.WriteCrypt(packet.SendAddonsInfo(s.AddonData))
		log.Println("Addon info sent")
	}

	v2 := packet.NewWorldPacket(packet.SMSG_CLIENTCACHE_VERSION)
	// vbf := make([]byte, 4)
	// // binary.LittleEndian.PutUint32(vbf, 24)
	// binary.LittleEndian.PutUint32(vbf, 12340)
	// v2.Write(vbf)
	v2.WriteUint32(12340)
	s.WriteCrypt(v2.Finish())
	log.Println("Client Cache sent.")

	v3 := packet.NewWorldPacket(packet.SMSG_TUTORIAL_FLAGS)
	for i := 0; i < 8; i++ {
		v3.WriteUint32(0x111111)
	}
	s.WriteCrypt(v3.Finish())
	log.Println("Tutorial flags sent.")
}
