package model

// IpfsFileInfo 文件权限信息
type IpfsFileInfo struct {
	Cid       string `json:"cid"`
	Uid       string `json:"uid"`
	Size      int64  `json:"size"`
	Owner     string `json:"owner"`
	State     uint8  `json:"state"`
	StoreDays int64  `json:"storeDays"`
}

type IpfsFile struct {
	Cid         string `json:"cid"`
	Owner       string `json:"owner"`
	Size        uint64 `json:"size"`
	ExpireBlock uint64 `json:"expireBlockNum"`
}

type IpfsMining struct {
	Cid  string `json:"cid"`
	Hash string `json:"hash"`
}
