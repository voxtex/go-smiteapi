package smiteapi

type Match struct {
	ActiveId1    int
	ActiveId2    int
	Active1      string `json:"Active_1"`
	Active2      string `json:"Active_2"`
	Active3      string `json:"Active_3"`
	Assists      int
	Creeps       int
	Damage       int
	DamageTaken  int `json:"Damage_Taken"`
	Deaths       int
	God          string
	GodId        int
	Gold         int
	Healing      int
	ItemId1      int
	ItemId2      int
	ItemId3      int
	ItemId4      int
	ItemId5      int
	ItemId6      int
	Item1        string `json:"Item_1"`
	Item2        string `json:"Item_2"`
	Item3        string `json:"Item_3"`
	Item4        string `json:"Item_4"`
	Item5        string `json:"Item_5"`
	Item6        string `json:"Item_6"`
	KillingSpree int    `json:"Killing_Spree"`
	Kills        int
	Level        int
	Match        int
	MatchTime    SmiteTime `json:"Match_Time"`
	Minutes      int
	MultiKillMax int `json:"Multi_kill_Max"`
	Queue        string
	Skin         string
	SkinId       int
	Surrendered  string
	Team1Score   int
	Team2Score   int
	WinStatus    string `json:"Win_Status"`
	PlayerName   string `json:"playerName"`
}

type TopMatch struct {
	Ban1              string
	Ban1Id            int
	Ban2              string
	Ban2Id            int
	EntryDateTime     SmiteTime `json:"Entry_Datetime"`
	LiveSpectators    int
	Match             int64
	MatchTime         int `json:"Match_Time"`
	OfflineSpectators int
	Queue             string
	RecordingFinished SmiteTime
	RecordingStarted  SmiteTime
	Team1AvgLevel     int `json:"Team1_AvgLevel"`
	Team1Gold         int `json:"Team1_Gold"`
	Team1Kills        int `json:"Team1_Kills"`
	Team1Score        int `json:"Team1_Score"`
	Team2AvgLevel     int `json:"Team2_AvgLevel"`
	Team2Gold         int `json:"Team2_Gold"`
	Team2Kills        int `json:"Team2_Kills"`
	Team2Score        int `json:"Team2_Score"`
	WinningTeam       int
}

type MatchDetail struct {
	AccountLevel       int `json:"Account_Level"`
	ActiveId1          int
	ActiveId2          int
	Assists            int
	Ban1               string
	Ban1Id             int
	Ban2               string
	Ban2Id             int
	Ban3               string
	Ban3Id             int
	Ban4               string
	Ban4Id             int
	DamageBot          int `json:"Damage_Bot"`
	DamageDoneMagical  int `json:"Damage_Done_Magical"`
	DamageDonePhysical int `json:"Damage_Done_Physical"`
	DamageMitigated    int `json:"Damage_Mitigated"`
	DamagePlayer       int `json:"Damage_Player"`
	DamageTaken        int `json:"Damage_Taken"`
	Deaths             int
	EntryDateTime      SmiteTime `json:"Entry_Datetime"`
	FinalMatchLevel    int       `json:"Final_Match_Level"`
	GodId              int
	GoldEarned         int `json:"Gold_Earned"`
	GoldPerMinute      int `json:"Gold_Per_Minute"`
	Healing            int
	ItemId1            int
	ItemId2            int
	ItemId3            int
	ItemId4            int
	ItemId5            int
	ItemId6            int
	ItemActive1        string `json:"Item_Active_1"`
	ItemActive2        string `json:"Item_Active_2"`
	ItemActive3        string `json:"Item_Active_3"`
	ItemPurch1         string `json:"Item_Purch_1"`
	ItemPurch2         string `json:"Item_Purch_2"`
	ItemPurch3         string `json:"Item_Purch_3"`
	ItemPurch4         string `json:"Item_Purch_4"`
	ItemPurch5         string `json:"Item_Purch_5"`
	ItemPurch6         string `json:"Item_Purch_6"`
	KillingSpree       int    `json:"Killing_Spree"`
	KillsBot           int    `json:"Kills_Bot"`
	KillsPlayer        int    `json:"Kills_Player"`
	MasteryLevel       int    `json:"Mastery_Level"`
	Match              int64
	Minutes            int
	MultiKillMax       int `json:"Multi_kil_Max"`
	PartyId            int
	God                string `json:"Reference_Name"`
	Skin               string
	SkinId             int
	StructureDamage    int `json:"Structure_Damage"`
	Surrendered        string
	Team1Score         int
	Team2Score         int
	TeamId             int
	TeamName           string `json:"Team_Name"`
	WinStatus          string `json:"Win_Status"`
	Queue              string `json:"name"`
	PlayerName         string `json:"playerName"`
}
