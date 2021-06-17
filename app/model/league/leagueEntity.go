package league

//Entity leagues实例
type Entity struct {
	LeagueID         float64 `json:"leagueId"`         //赛事ID
	IsIntegralLeague bool    `json:"isIntegralLeague"` //是否是积分联赛
	LeagueName       string  `json:"leagueName"`
	Organizer        string  `json:"organizer"`
	StartTime        float64 `json:"startTime"`
	EndTime          float64 `json:"endTime"`
	PrizePoll        string  `json:"prizePoll"`
	LeagueAbbr       string  `json:"leagueAbbr"`
	LeagueType       float64 `json:"leagueType"` //联赛类型  0 都不是  1 major  2 minor
	Integral         float64 `json:"integral"`   //积分
	Thumb            string  `json:"thumb"`      //图片
}
