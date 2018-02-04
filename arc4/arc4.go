package arc4

// Warning: this package is ONLY suitable for Gophercraft.
// Do not use it in any other case: it provides no actual security.
type Cipher struct {
	S    []byte
	i, j byte
}

func ARC4(key []byte) *Cipher {
	this := &Cipher{}
	this.S = make([]byte, 256)
	for i := 0; i < 256; i++ {
		this.S[i] = byte(i)
	}

	var j uint8 = 0
	var t uint8 = 0
	for i := 0; i < 256; i++ {
		j = (j + this.S[i] + key[i%len(key)]) & 255
		t = this.S[i]
		this.S[i] = this.S[j]
		this.S[j] = t
	}

	this.i = 0
	this.j = 0
	return this
}

func (this *Cipher) Next() uint8 {
	var t uint8
	this.i = (this.i + 1) & 255
	this.j = (this.j + this.S[this.i]) & 255
	t = this.S[this.i]
	this.S[this.i] = this.S[this.j]
	this.S[this.j] = t
	return this.S[(t+this.S[this.i])&255]
}

func (this *Cipher) Encrypt(data []byte) {
	for i := 0; i < len(data); i++ {
		data[i] ^= this.Next()
	}
}

func (this *Cipher) Decrypt(data []byte) {
	this.Encrypt(data)
}
