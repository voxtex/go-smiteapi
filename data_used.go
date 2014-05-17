package smiteapi

type DataUsed struct {
	ActiveSessions     int `json:"Active_Sessions"`
	ConcurrentSessions int `json:"Concurrent_Sessions"`
	RequestLimitDaily  int `json:"Request_Limit_daily"`
	SessionCap         int `json:"Session_Cap"`
	SessionTimeLimit   int `json:"Session_Time_Limit"`
	TotalRequestsToday int `json:"TotalRequestsToday"`
	TotalSessionsToday int `json:"Total_Sessions_Today"`
}
