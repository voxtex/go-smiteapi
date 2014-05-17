package smiteapi

type PlayerLeagueInfo struct {
	Leaves        int
	Losses        int
	Name          string
	Points        int
	PrevRank      int
	Rank          int
	Season        int
	Tier          int
	VictoryPoints int
	Wins          int
}

type Player struct {
	CreatedDateTime   SmiteTime `json:"Created_Datetime"`
	Id                int
	LastLoginDateTime SmiteTime `json:"Last_Login_Datetime"`
	LeagueArena       PlayerLeagueInfo
	LeagueConquest    PlayerLeagueInfo
	Leaves            int
	Level             int
	Losses            int
	MasteryLevel      int
	Name              string
	RankStat          int `json:"Rank_Stat"`
	RankStatArena     int `json:"Rank_Stat_Arena"`
	TeamId            int
	TeamName          string `json:"Team_Name"`
	Wins              int
}

type LeaderboardPlayer struct {
	Leaves        int
	Losses        int
	Name          string
	Points        int
	PrevRank      int
	Rank          int
	Season        int
	Tier          int
	VictoryPoints int
	Wins          int
}
