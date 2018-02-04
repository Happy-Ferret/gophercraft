package guid

import "testing"

func TestGUID(t *testing.T) {
	var guid1 GUID = 0xDEADBEEF1337BADC
	guid1 = guid1.SetHigh(Player)

	if guid1 != 0x0000BEEF1337BADC {
		t.Errorf("SetHigh failed")
	}

	if guid1.High() != Player {
		t.Errorf("Fail to extract high properly")
	}

	var g2 GUID = 0x0000000000521BC0
	if g2.High() != Player {
		t.Errorf("Fail to extract high from nokturnim GUID")
	}
}
