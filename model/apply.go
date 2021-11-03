package model

// Apply 文件请求
type Apply struct {
	Cid         string `json:"cid"`         //文件块cid
	ApplyPeerId string `json:"applyPeerId"` // 申请节点id

	ResponseState  int8   `json:"responseState"`            //  响应状态  0：待响应  1：响应成功  2：响应中
	ResponsePeerId string `json:"responsePeerId,omitempty"` // 响应节点id

	Retain interface{} `json:"retain,omitempty"` // 保留字段
}

type JoinServerListApply struct {
	CorePeer CorePeer       `json:"serverPeer"`
	VoteMap  map[string]int `json:"voteMap,omitempty"`
}
