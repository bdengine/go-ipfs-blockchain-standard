package model

type Peer struct {
	UserCode string `json:"userCode"`
	PeerId   string `json:"peerId"`
}

type CorePeer struct {
	PeerId    string   `json:"peerId"`
	Addresses []string `json:"addresses"`
}

const (
	RoleCore = 1
	RoleLite = 0
)
