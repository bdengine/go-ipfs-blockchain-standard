package model

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
