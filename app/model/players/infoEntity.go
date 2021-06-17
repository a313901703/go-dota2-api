package players

//InfoEntity player info entity
type InfoEntity struct {
	PlayerID    float64     `json:"playerID"`
	Avatar      interface{} `json:"avatar"`
	Mmr         interface{} `json:"mmr"`
	Personaname interface{} `json:"personaname"`
	Steamid     interface{} `json:"steamid"`
}

//FriendInfoEntity friend info entity
type FriendInfoEntity struct {
	WithGames interface{} `json:"withGames"`
	WithWin   interface{} `json:"withWin"`
	*InfoEntity
}

//NewPlayerInfoEntity   玩家详细信息实例
func NewPlayerInfoEntity(item map[string]interface{}) *InfoEntity {
	entity := new(InfoEntity)
	if item["profile"] != nil {
		profile := item["profile"].(map[string]interface{})
		entity.PlayerID = profile["account_id"].(float64)
		entity.Avatar = profile["avatar"]
		entity.Personaname = profile["personaname"]
		entity.Steamid = profile["steamid"]
	}

	if item["mmr_estimate"] != nil {
		mmrEstimate := item["mmr_estimate"].(map[string]interface{})
		entity.Mmr = mmrEstimate["estimate"]
	}
	return entity
}

//CreateFriendInfoEntity  friends entity
func CreateFriendInfoEntity(item map[string]interface{}) *FriendInfoEntity {
	entity := new(InfoEntity)
	entity.PlayerID = item["account_id"].(float64)
	entity.Avatar = item["avatar"]
	entity.Personaname = item["personaname"]
	entity.Steamid = item["steamid"]

	friendEntity := &FriendInfoEntity{InfoEntity: entity}
	friendEntity.WithGames = item["with_games"]
	friendEntity.WithWin = item["with_win"]
	return friendEntity
}

//BatchCreateFriendInfoEntity  barch create friends entity
func BatchCreateFriendInfoEntity(items []interface{}) []*FriendInfoEntity {
	friends := make([]*FriendInfoEntity, len(items))
	for k, v := range items {
		friends[k] = CreateFriendInfoEntity(v.(map[string]interface{}))
	}
	return friends
}
