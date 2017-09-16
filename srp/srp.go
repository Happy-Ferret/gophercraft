package srp

import (
	"bytes"
	"crypto/sha1"
	"strings"
)

// Warning: this package is ONLY suitable for Gophercraft.
// Do not use it in any other case: it provides no actual security.

func Credentials(username, password string) []byte {
	I := strings.ToUpper(username)
	P := strings.ToUpper(password)
	return []byte(I + ":" + P)
}

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}

func SRPCalculate(username, password string, _B, n, salt []byte) ([]byte, []byte) {
	g := BigNumFromInt(7)
	k := BigNumFromInt(3)

	N := BigNumFromArray(n)
	s := BigNumFromArray(salt)

	a := BigNumFromRand(19)
	A := g.ModExp(a, N)

	B := BigNumFromArray(_B)

	auth := hash(Credentials(username, password))

	x := BigNumFromArray(hash(s.ToArray(), auth))

	// v := g.ModExp(x, N)

	uh := hash(A.ToArray(), B.ToArray())
	u := BigNumFromArray(uh)

	kgx := k.Multiply(g.ModExp(x, N))
	aux := a.Add(u.Multiply(x))

	_S := B.Subtract(kgx).ModExp(aux, N)
	S := _S.ToArray()

	S1, S2 := make([]byte, 16), make([]byte, 16)

	for i := 0; i < 16; i++ {
		S1[i] = S[i*2]
		S2[i] = S[i*2+1]
	}

	S1h := hash(S1)
	S2h := hash(S2)

	K := make([]byte, 40)

	for i := 0; i < 20; i++ {
		K[i*2] = S1h[i]
		K[i*2+1] = S2h[i]
	}

	userh := hash([]byte(strings.ToUpper(username)))

	Nh := hash(N.ToArray())
	gh := hash(g.ToArray())

	Ngh := make([]byte, 20)
	for i := 0; i < 20; i++ {
		Ngh[i] = Nh[i] ^ gh[i]
	}

	M1 := hash(
		Ngh,
		userh,
		s.ToArray(),
		A.ToArray(),
		B.ToArray(),
		K,
	)

	return A.ToArray(), M1
}
