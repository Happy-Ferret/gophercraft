package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AuthLogonProof_S struct {
	Cmd          AuthType
	Error        ErrorType
	M2           []byte
	AccountFlags uint32
	SurveyID     uint32
	Unk3         uint16
}

func (alps *AuthLogonProof_S) Encode() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(uint8(alps.Cmd))
	buf.WriteByte(uint8(alps.Error))
	buf.Write(alps.M2)
	acf := make([]byte, 4)
	binary.LittleEndian.PutUint32(acf, alps.AccountFlags)
	buf.Write(acf)
	sid := make([]byte, 4)
	binary.LittleEndian.PutUint32(sid, alps.SurveyID)
	buf.Write(sid)
	unk3 := make([]byte, 2)
	binary.LittleEndian.PutUint16(unk3, alps.Unk3)
	buf.Write(unk3)
	return buf.Bytes()
}

func UnmarshalAuthLogonProof_S(input []byte) (*AuthLogonProof_S, error) {
	if len(input) < 32 {
		return nil, fmt.Errorf("Packet too small")
	}
	alps := &AuthLogonProof_S{}
	alps.Cmd = AuthType(input[0])
	alps.Error = ErrorType(input[1])
	alps.M2 = (input[2:22])
	alps.AccountFlags = binary.LittleEndian.Uint32(input[22:26])
	alps.SurveyID = binary.LittleEndian.Uint32(input[26:30])
	alps.Unk3 = binary.LittleEndian.Uint16(input[30:32])
	return alps, nil
}
