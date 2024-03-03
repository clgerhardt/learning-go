package types

type WSMessage struct {
	Type string   `json:"type"`
	Data [] `json:"data"`
}

type Login struct {
	ClientId int    `json:"clientID"`
	Username string `json:"username"`
}
