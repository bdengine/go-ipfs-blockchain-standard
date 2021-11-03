package model

type Peer struct {
	UserCode string `json:"userCode"`
	PeerId   string `json:"peerId"`
	Role     uint8  `json:"role"`
}

type CorePeer struct {
	Org       string `json:"org"`
	Peer      `json:"peer"`
	Addresses []string `json:"addresses"`
}
