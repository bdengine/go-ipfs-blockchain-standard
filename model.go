package standard

// Apply 文件请求
type Apply struct {
	Cid         string `json:"cid"`         //文件块cid
	ApplyPeerId string `json:"applyPeerId"` // 申请节点id

	ResponseState  int8   `json:"responseState"`            //  响应状态  0：待响应  1：响应成功  2：响应中
	ResponsePeerId string `json:"responsePeerId,omitempty"` // 响应节点id

	Retain interface{} `json:"retain,omitempty"` // 保留字段
}

// IpfsFileInfo 文件权限信息
type IpfsFileInfo struct {
	Cid          string           `json:"cid"`
	State        uint8            `json:"state"`
	AuthorityMap map[string]Limit `json:"authorityMap"` // userCode:Limit   userCode = orgName + examerId + peerId
	Owner        string           `json:"owner"`
}

// Limit 权限限制信息
type Limit struct {
	ReadCount  int   `json:"readCount"`  // 读限制次数
	ReadTime   int64 `json:"readTime"`   // 读限制时间
	ShareCount int   `json:"shareCount"` // 分享限制次数
	ShareTime  int64 `json:"shareTime"`  // 分享限制时间
	State      uint8 `json:"state"`      // 状态 00000000      // 读限制类型 00 // 分享限制类型 00 // 通知状态 00 // 权限生效状态 00
}

type JoinServerListApply struct {
	CorePeer CorePeer       `json:"serverPeer"`
	VoteMap  map[string]int `json:"voteMap,omitempty"`
}

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
