package gcore

import (
	"fmt"
	"log"
	"strings"

	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-xorm/xorm"
	"github.com/superp00t/gophercraft/authserver"
	"github.com/superp00t/gophercraft/packet"
)

type Core struct {
	DB *xorm.Engine
}

func NewCore(driver, source string) (*Core, error) {
	db, err := xorm.NewEngine(driver, source)
	if err != nil {
		return nil, err
	}

	strcts := []interface{}{
		new(Account),
		new(Realm),
		new(SessionKey),
		new(Character),
	}

	for _, v := range strcts {
		err = db.Sync2(v)
		if err != nil {
			return nil, err
		}
	}

	core := &Core{DB: db}

	return core, nil
}

func (c *Core) StoreKey(user string, K []byte) {
	c.DB.Where("account = ?", user).Delete(new(SessionKey))

	c.DB.Insert(&SessionKey{
		Account: user,
		K:       K,
	})
}

func (c *Core) GetAccount(user string) *authserver.Account {
	var accs []Account
	c.DB.Where("username = ?", strings.ToUpper(user)).Find(&accs)
	if len(accs) == 0 {
		return nil
	}

	return &authserver.Account{
		Username:     accs[0].Username,
		IdentityHash: accs[0].IdentityHash,
	}
}

type RealmPublish struct {
	At      int64
	Listing *packet.RealmListing
}

func (c *Core) ListRealms(user string) []packet.RealmListing {
	var acc []Account
	c.DB.Where("username = ?", user).Find(&acc)
	if len(acc) == 0 {
		log.Println("No user found!")
		return nil
	}

	var rlmState []Realm
	c.DB.Find(&rlmState)

	var rlm []packet.RealmListing
	for i, v := range rlmState {
		pkt := packet.RealmListing{}
		pkt.Type = v.Type
		pkt.Locked = false
		pkt.Flags = 0x00
		if (time.Now().UnixNano() - v.LastUpdated.UnixNano()) > (time.Second * 15).Nanoseconds() {
			pkt.Flags = 0x02
		}
		pkt.Name = v.Name
		pkt.Address = v.Address
		pkt.Population = 2
		pkt.Timezone = 8
		pkt.ID = uint8(i)
		c, _ := c.DB.Where("realm = ?", v.Name).Where("account = ?", acc[0].ID).Count(new(Character))
		pkt.Characters = uint8(c)
		rlm = append(rlm, pkt)
	}

	log.Println(spew.Sdump(rlm))

	return rlm
}

const banner = `
 ____ ____ ___  _  _ ____ ____ ____ ____ ____ ____ ___
 |__, [__] |--' |--| |=== |--< |___ |--< |--| |---  | 

 The programs included with Gophercraft are free software;
the exact distribution terms for each program are described in LICENSE.

`

func PrintLicense() {
	fmt.Println(banner)
}
