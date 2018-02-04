package client

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/packet"
)

type ClientHandler struct {
	Type packet.WorldType
	Fn   func([]byte)
}

func (cl *Client) HandleFunc(t packet.WorldType, fn func([]byte)) {
	cl.Handlers[t] = &ClientHandler{
		Type: t,
		Fn:   fn,
	}
}

func (cl *Client) HandleCharList(b []byte) {
	data, err := packet.UnmarshalCharacterList(b)
	if err != nil {
		log.Fatal(err)
	}

	glogger.Warnln(spew.Sdump(data))
	if len(b) > 100 {
		glogger.Warnln("Weird packet: ", b)
	}

	for _, v := range data.Characters {
		if v.Name == cl.Player {
			pkt := packet.NewGamePacket(packet.CMSG_PLAYER_LOGIN)
			b := packet.NewEtcBuffer(nil)
			b.WriteUint64(v.GUID.U64())
			pkt.Buf.Write(b.Encode())
			cl.SendCrypt(pkt.Finish())
			return
		}
	}
}

func (cl *Client) HandleMOTD(b []byte) {
	glogger.Println(spew.Sdump(b))
}

func (cl *Client) HandleLogin(b []byte) {
	d, _ := packet.UnmarshalSMSGAuthResponse(b)
	glogger.Println(spew.Sdump(d))
	if d.Cmd == packet.AUTH_OK {
		pkt := packet.NewGamePacket(packet.CMSG_CHAR_ENUM)
		cl.SendCrypt(pkt.Finish())
	}
}

func pbt(fix string, input []byte) {
	d := input[4:]
	stb := "Buf_" + fix + " := []byte{ "
	for _, v := range d {
		stb += fmt.Sprintf("0x%X, ", v)
	}
	stb += " }"
	glogger.Println(stb)
}

func (cl *Client) HandleEquip(d []byte) {
	pbt("equip", d)
}

func (cl *Client) HandleActions(d []byte) {
	pbt("actionbuttons", d)
}

func (cl *Client) HandleReputations(d []byte) {
	pbt("reps", d)
}

func (cl *Client) HandleSocialList(d []byte) {
	pbt("sociallist", d)
}

func (cl *Client) HandleDanceMoves(d []byte) {
	pbt("dance", d)
}

func (cl *Client) HandleForcedReactions(d []byte) {
	pbt("forced", d)
}

func (cl *Client) HandleSpellList(d []byte) {
	pbt("spelllist", d)
}

func (cl *Client) HandleUpdateData(d []byte) {
	s, err := packet.UnmarshalObjectUpdate(d)
	if err != nil {
		log.Fatal(err)
	}

	glogger.Println(spew.Sdump(s))
}

func (cl *Client) Handle() {
	cl.HandleFunc(packet.SMSG_WARDEN_DATA, cl.HandleWarden)
	cl.HandleFunc(packet.SMSG_AUTH_RESPONSE, cl.HandleLogin)
	cl.HandleFunc(packet.SMSG_CHAR_ENUM, cl.HandleCharList)
	cl.HandleFunc(packet.SMSG_MOTD, cl.HandleMOTD)

	cl.HandleFunc(packet.SMSG_INITIAL_SPELLS, cl.HandleSpellList)
	cl.HandleFunc(packet.SMSG_EQUIPMENT_SET_LIST, cl.HandleEquip)
	cl.HandleFunc(packet.SMSG_ACTION_BUTTONS, cl.HandleActions)
	cl.HandleFunc(packet.SMSG_INITIALIZE_FACTIONS, cl.HandleReputations)
	cl.HandleFunc(packet.SMSG_CONTACT_LIST, cl.HandleSocialList)
	cl.HandleFunc(packet.SMSG_LEARNED_DANCE_MOVES, cl.HandleDanceMoves)
	cl.HandleFunc(packet.SMSG_SET_FORCED_REACTIONS, cl.HandleForcedReactions)

	cl.HandleFunc(packet.SMSG_UPDATE_OBJECT, cl.HandleUpdateData)
	cl.HandleFunc(packet.SMSG_COMPRESSED_UPDATE_OBJECT, cl.HandleUpdateData)

	for {
		bufs := cl.ReadCrypt()
		for _, buf := range bufs {
			// cl.Steal <- buf
			opcode := packet.WorldType(binary.LittleEndian.Uint16(buf[2:4]))
			glogger.Warnln(opcode)
			if h := cl.Handlers[opcode]; h != nil {
				glogger.Warnln("Handling", opcode)
				h.Fn(buf)
			} else {
				glogger.Warnln("No handler for ", opcode)
			}
		}
	}
}
