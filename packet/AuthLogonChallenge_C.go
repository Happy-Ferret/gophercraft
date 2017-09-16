package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// AuthLogonChallenge_C is the first packet sent by a client
// while initiating a connection to a World of Warcraft authserver.
type AuthLogonChallenge_C struct {
	Cmd          AuthType
	Error        ErrorType
	Size         uint16
	GameName     []byte // Encode in reverse.
	Version1     uint8
	Version2     uint8
	Version3     uint8
	Build        uint16
	Platform     []byte
	OS           []byte
	Country      []byte
	TimezoneBias uint32
	IP           uint32
	ILen         uint8
	I            []byte
}

func (alcc *AuthLogonChallenge_C) Encode() []byte {
	header := []byte{uint8(alcc.Cmd), uint8(alcc.Error)}
	size := make([]byte, 2)
	footer := new(bytes.Buffer)
	footer.Write(alcc.GameName)
	footer.Write([]byte{alcc.Version1, alcc.Version2, alcc.Version3})
	buildno := make([]byte, 2)
	binary.LittleEndian.PutUint16(buildno, alcc.Build)
	footer.Write(buildno)
	footer.Write(reverseBuffer(alcc.Platform))
	footer.Write(reverseBuffer(alcc.OS))
	footer.Write(reverseBuffer(alcc.Country))
	tz := make([]byte, 4)
	binary.LittleEndian.PutUint32(tz, alcc.TimezoneBias)
	footer.Write(tz)
	ip := make([]byte, 4)
	binary.LittleEndian.PutUint32(ip, alcc.IP)
	footer.Write(ip)
	footer.WriteByte(uint8(len(alcc.I)))
	footer.Write(alcc.I)
	body := new(bytes.Buffer)
	binary.LittleEndian.PutUint16(size, uint16(footer.Len()))
	body.Write(header)
	body.Write(size)
	body.Write(footer.Bytes())
	return body.Bytes()
}

func (alcc *AuthLogonChallenge_C) Version() string {
	return fmt.Sprintf("%d.%d.%d", alcc.Version1, alcc.Version2, alcc.Version3)
}

func UnmarshalAuthLogonChallenge_C(data []byte) (*AuthLogonChallenge_C, error) {
	ac := &AuthLogonChallenge_C{}
	ac.Cmd = AuthType(data[0])
	ac.Error = ErrorType(data[1])
	ac.Size = binary.LittleEndian.Uint16(data[2:4])
	ac.GameName = reverseBuffer(data[4:8])
	ac.Version1 = data[8]
	ac.Version2 = data[9]
	ac.Version3 = data[10]
	ac.Build = binary.LittleEndian.Uint16(data[11:13])
	ac.Platform = reverseBuffer(data[13:17])
	ac.OS = reverseBuffer(data[17:21])
	ac.Country = reverseBuffer(data[21:25])
	ac.TimezoneBias = binary.LittleEndian.Uint32(data[25:29])
	ac.IP = binary.LittleEndian.Uint32(data[29:33])
	ac.ILen = data[33]
	ac.I = data[34 : 34+int(ac.ILen)]
	// TODO: bounds check
	return ac, nil
}

// LogonChallengePacket_C is a helper function to simplify the client library.
func LogonChallengePacket_C(username string) []byte {
	usr := []byte(strings.ToUpper(username))
	alcc := &AuthLogonChallenge_C{
		Cmd:          AUTH_LOGON_CHALLENGE,
		Error:        8,
		GameName:     packetString(".WoW"),
		Version1:     3, // 3.3.5a
		Version2:     3,
		Version3:     5,
		Build:        12340,
		Platform:     packetString(".x86"),
		OS:           packetString(".Win"),
		Country:      packetString("enGB"),
		TimezoneBias: 4294966996,
		IP:           16777343,
		I:            usr,
	}

	return alcc.Encode()
}
