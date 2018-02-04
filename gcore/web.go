package gcore

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gorilla/mux"
	"github.com/superp00t/gophercraft/guid"
	"github.com/superp00t/gophercraft/srp"
)

type Account struct {
	ID           int64
	Username     string
	IdentityHash []byte
}

type SessionKey struct {
	Account string
	K       []byte
}
type Realm struct {
	Name        string
	Type        uint8
	Address     string
	LastUpdated time.Time
}

type Character struct {
	Name       string    `json:"name"`
	Level      uint8     `json:"level"`
	Realm      string    `json:"realm"`
	Account    string    `json:"account"`
	GUID       guid.GUID `json:"guid"`
	Race       uint8     `json:"race"`
	Class      uint8     `json:"class"`
	Gender     uint8     `json:"gender"`
	Skin       uint8     `json:"skin"`
	Face       uint8     `json:"face"`
	HairStyle  uint8     `json:"hairStyle"`
	HairColor  uint8     `json:"hairColor"`
	FacialHair uint8     `json:"facialHair"`
}

func (c *Core) NewCaptcha(r *Request) {
	cp := captcha.New()
	r.Encode(http.StatusOK, CaptchaResponse{
		Status:    http.StatusOK,
		CaptchaID: cp,
	})
}

func (c *Core) UserExists(r *Request) {
	t := strings.ToUpper(r.Vars["username"])
	var acc []Account
	err := c.DB.Where("username = ?", t).Find(&acc)
	if err != nil {
		r.Respond(http.StatusInternalServerError, "Internal server error")
		return
	}

	r.Encode(http.StatusOK, UserExistsResponse{
		Status:     http.StatusOK,
		UserExists: len(acc) == 1,
	})
}

func (c *Core) Register(r *Request) {
	if r.Vars["username"] == "" || r.Vars["password"] == "" {
		r.Respond(http.StatusBadRequest, "username and password must not be empty")
		return
	}

	if !captcha.VerifyString(r.Vars["capid"], r.Vars["capsol"]) {
		r.Respond(http.StatusOK, "Bad captcha")
		return
	}

	u := strings.ToUpper(r.Vars["username"])
	p := strings.ToUpper(r.Vars["password"])

	var acc []Account
	err := c.DB.Where("username = ?", u).Find(&acc)
	if err != nil {
		r.Respond(http.StatusInternalServerError, "Internal server error")
		return
	}

	if len(acc) > 0 {
		r.Respond(http.StatusOK, "Username in use")
		return
	}

	idhash := srp.HashCredentials(u, p)

	c.DB.Insert(&Account{
		Username:     u,
		IdentityHash: idhash,
	})

	r.Respond(http.StatusOK, "")
}

func (c *Core) APIKey(key string) int {
	if key == "_proper_" {
		return 2
	}

	return 0
}

const (
	PVPServer     = 0x01
	RPServer      = 0x02
	DefaultServer = 0x00
)

func (c *Core) PublishRealmInfo(addr, n string, stype uint8) {
	t := stype

	c.DB.Where("name = ?", n).Delete(new(Realm))
	c.DB.Insert(&Realm{
		Name:        n,
		Type:        t,
		Address:     addr,
		LastUpdated: time.Now(),
	})
}

func (c *Core) RealmState() []Realm {
	var r []Realm
	c.DB.Where("type > 0").Find(&r)
	return r
}

func (c *Core) RealmList(r *Request) {
	r.Encode(http.StatusOK, map[string]interface{}{
		"status":  200,
		"listing": c.RealmState(),
	})
}

func (c *Core) WebAPI() http.Handler {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1/").Subrouter()

	v1.Handle("/newCaptcha", c.Intercept(0, c.NewCaptcha))
	v1.Handle("/userExists/{username}", c.Intercept(0, c.UserExists))
	v1.Handle("/register/{username}/{password}/{capid}/{capsol}", c.Intercept(0, c.Register))
	v1.PathPrefix("/captcha/").Handler(captcha.Server(captcha.StdWidth, captcha.StdHeight))
	v1.Handle("/realmList", c.Intercept(0, c.RealmList))

	// admin/realm RPC functions

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(os.Getenv("GOPATH") + "src/github.com/superp00t/gophercraft/gcore/webapp/public/")))
	return r
}

func (c *Core) Intercept(required int, fn RequestHandler) *Interceptor {
	return &Interceptor{required, c, fn}
}

type Interceptor struct {
	requiredLevel int
	core          *Core
	fn            RequestHandler
}

func (s *Interceptor) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// TODO: Rate limiting and authorization
	lvl := s.core.APIKey(req.URL.Query().Get("a"))

	if lvl < s.requiredLevel {
		r := &Request{
			RW:   rw,
			R:    req,
			Vars: mux.Vars(req),
		}
		r.Respond(http.StatusUnauthorized, "not enough clearance")
		return
	}

	s.fn(&Request{
		RW:   rw,
		R:    req,
		Vars: mux.Vars(req),
	})
}
