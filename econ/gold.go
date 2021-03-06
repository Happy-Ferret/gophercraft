package econ

import (
	"fmt"
	"strconv"
	"strings"
)

type Money int32

const (
	Copper Money = 1
	Silver Money = 100 * Copper
	Gold   Money = 100 * Silver
)

func (m Money) Int32() int32 {
	return int32(m)
}

func (m Money) Coins() []Money {
	tgold := m / Gold
	silver := (m - (tgold * Gold)) / Silver
	copper := (m - (silver * Silver)) - (tgold * Gold)
	return []Money{tgold, silver, copper}
}

func (m Money) String() string {
	c := m.Coins()
	return fmt.Sprintf("%d Gold, %d Silver, %d Copper", c[0], c[1], c[2])
}

func (m Money) ShortString() string {
	c := m.Coins()
	return fmt.Sprintf("%dg %ds %dc", c[0], c[1], c[2])
}

func getCoinKey(s string) (string, string) {
	l := len(s)
	key := string(s[l-1])
	val := s[:l-1]
	return key, val
}

func ParseShortString(input string) (Money, error) {
	t := strings.Split(input, " ")
	if len(t) != 3 {
		return 0, fmt.Errorf("econ: parse error")
	}

	gk, gv := getCoinKey(t[0])
	sk, sv := getCoinKey(t[1])
	ck, cv := getCoinKey(t[2])

	if gk != "g" || sk != "s" || ck != "c" {
		return 0, fmt.Errorf("econ: invalid coin type")
	}

	gM, err := strconv.ParseInt(gv, 10, 64)
	if err != nil {
		return 0, err
	}
	sM, err := strconv.ParseInt(sv, 10, 64)
	if err != nil {
		return 0, err
	}
	cM, err := strconv.ParseInt(cv, 10, 64)
	if err != nil {
		return 0, err
	}

	var out Money
	out = out + Money(cM)
	out = out + (Money(sM) * Silver)
	out = out + (Money(gM) * Gold)

	return out, nil
}
