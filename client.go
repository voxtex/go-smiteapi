package smiteapi

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"time"
)

const (
	timeFormat      = "20060102150405"
	rootUrl         = "http://api.smitegame.com/smiteapi.svc"
	sessionDuration = 15 * time.Minute
)

// api specific values
const (
	// languages
	English = 1
	French  = 3
	// league tiers
	Bronze   = 1
	Silver   = 2
	Gold     = 3
	Platinum = 4
	Diamond  = 5
	// queues
	Conquest5v5         = 423
	NoviceQueue         = 424
	Conquest            = 426
	Practice            = 427
	ConquestChallenge   = 429
	ConquestRanked      = 430
	Domination          = 433
	MOTD                = 434
	Arena               = 435
	ArenaChallenge      = 438
	DominationChallenge = 439
	Joust               = 440
	JoustChallenge      = 441
	Assault             = 445
	AssaultChallenge    = 446
	Joust3v3            = 448
	ConquestLeague      = 451
	ArenaLeague         = 452
)

func NewClient(authKey, devId string) *Client {
	return &Client{language: English, authKey: authKey, devId: devId}
}

type Client struct {
	authKey          string
	devId            string
	language         int
	sessionId        string
	sessionTimestamp time.Time
	sessionLock      sync.Mutex
}

func getJsonBody(url string, data interface{}) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(data); err != nil {
		log.Fatal(err)
	}
}

func cleanPlayerName(name string) string {
	rx, err := regexp.Compile(`^\[.+\]\s`)
	if err != nil {
		log.Fatal(err)
	}
	teamless := rx.ReplaceAllString(name, "")
	return url.QueryEscape(teamless)
}

func timestamp() string {
	return time.Now().UTC().Format(timeFormat)
}

func (c *Client) signature(method string) string {
	h := md5.New()
	fmt.Fprintf(h, "%s%s%s%s", c.devId, method, c.authKey, timestamp())
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *Client) methodUrlWithoutSession(method string, args ...string) string {
	path := fmt.Sprintf("%s/%sjson/%s/%s/%s", rootUrl, method, c.devId, c.signature(method), timestamp())
	for _, v := range args {
		path += "/" + v
	}
	return path
}

func (c *Client) methodUrl(method string, args ...string) string {
	path := fmt.Sprintf("%s/%sjson/%s/%s/%s/%s", rootUrl, method, c.devId, c.signature(method), c.sessionId, timestamp())
	for _, v := range args {
		path += "/" + v
	}
	return path
}

func (c *Client) isSessionValid() bool {
	return (time.Since(c.sessionTimestamp) < sessionDuration)
}

func (c *Client) createSession() (sessionId string, err error) {
	data := map[string]string{}
	getJsonBody(c.methodUrlWithoutSession("createsession"), &data)
	c.sessionId = data["session_id"]
	if c.sessionId == "" {
		err = errors.New(data["ret_msg"])
	} else {
		c.sessionTimestamp = time.Now()
	}
	return c.sessionId, err
}

func (c *Client) checkSession() {
	if !c.isSessionValid() {
		c.sessionLock.Lock()
		defer c.sessionLock.Unlock()
		if !c.isSessionValid() {
			if _, err := c.createSession(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (c *Client) SetLanguage(lang int) {
	c.language = lang
}

func (c *Client) GetDataUsed() *DataUsed {
	c.checkSession()
	data := &[]*DataUsed{}
	getJsonBody(c.methodUrl("getdataused"), data)
	return (*data)[0]
}

func (c *Client) GetItems() []*Item {
	c.checkSession()
	data := &[]*Item{}
	getJsonBody(c.methodUrl("getitems", strconv.Itoa(c.language)), data)
	return *data
}

func (c *Client) GetGods() []*God {
	c.checkSession()
	data := &[]*God{}
	getJsonBody(c.methodUrl("getgods", strconv.Itoa(c.language)), data)
	return *data
}

func (c *Client) GetPlayer(name string) []*Player {
	c.checkSession()
	name = cleanPlayerName(name)
	data := &[]*Player{}
	getJsonBody(c.methodUrl("getplayer", name), data)
	return *data
}

func (c *Client) GetMatchHistory(name string) []*Match {
	c.checkSession()
	name = cleanPlayerName(name)
	data := &[]*Match{}
	getJsonBody(c.methodUrl("getmatchhistory", name), data)
	return *data
}

func (c *Client) GetLeagueLeaderboard(queue int, tier int, season int) []*LeaderboardPlayer {
	c.checkSession()
	data := &[]*LeaderboardPlayer{}
	getJsonBody(c.methodUrl("getleagueleaderboard", strconv.Itoa(queue), strconv.Itoa(tier), strconv.Itoa(season)), data)
	return *data
}

func (c *Client) GetLeagueSeasons(queue int) []*Season {
	c.checkSession()
	data := &[]*Season{}
	getJsonBody(c.methodUrl("getleagueseasons", strconv.Itoa(queue)), data)
	return *data
}

func (c *Client) GetTopMatches() []*TopMatch {
	c.checkSession()
	data := &[]*TopMatch{}
	getJsonBody(c.methodUrl("gettopmatches"), data)
	return *data
}

func (c *Client) GetMatchDetails(matchId int) []*MatchDetail {
	c.checkSession()
	data := &[]*MatchDetail{}
	getJsonBody(c.methodUrl("getmatchdetails", strconv.Itoa(matchId)), data)
	return *data
}
