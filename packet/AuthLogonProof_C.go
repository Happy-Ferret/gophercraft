package packet

import "bytes"

type AuthLogonProof_C struct {
	Cmd          AuthType
	A            []byte // 32 long
	M1           []byte // 20 long
	CRC          []byte // 20 long
	NumberOfKeys uint8
	SecFlags     uint8
}

// Client sends BE (Big Endian)
// Server reinterpret_casts struct, converting it to LE in C++
// Server converts it back to BE with SetBinary
func (alpc *AuthLogonProof_C) Encode() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(uint8(alpc.Cmd))
	buf.Write(alpc.A)
	buf.Write(alpc.M1)
	buf.Write(randomBuffer(20))
	buf.WriteByte(alpc.NumberOfKeys)
	buf.WriteByte(alpc.SecFlags)
	return buf.Bytes()
}

func UnmarshalAuthLogonProof_C(input []byte) (*AuthLogonProof_C, error) {
	alpc := &AuthLogonProof_C{}
	alpc.Cmd = AuthType(input[0])
	alpc.A = input[1:33]
	alpc.M1 = input[33:53]
	alpc.CRC = input[53:73]
	alpc.NumberOfKeys = input[73]
	alpc.SecFlags = input[74]
	return alpc, nil
}
