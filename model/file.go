package model

// IpfsFileInfo 文件权限信息
type IpfsFileInfo struct {
	Cid   string `json:"cid"`
	Uid   string `json:"uid"`
	State uint8  `json:"state"`
	Owner string `json:"owner"`
	Size  uint64 `json:"size"`
}

type IpfsMining struct {
	Cid  string `json:"cid"`
	Hash string `json:"hash"`
}
