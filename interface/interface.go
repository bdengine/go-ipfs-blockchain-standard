package _interface

import (
	"github.com/bdengine/go-ipfs-blockchain-standard/dto"
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
)

type FileApplyer interface {
	ApplyLocal(cid string) error
	ApplyRemote(cid string) error
}

/*type FileProvider interface {
	Response
}*/

type FileOwner interface {
	AddFile(info model.IpfsFileInfo) error
	DeleteFile(cid string) error
	RechargeFile(cid string, storeDays int64) error
}

type FileUploader interface {
	CheckTrustStatus(owner string) error
	FileOwner
}

// ChainInfoReader 该接口方法应使用读方法实现，在select层设有缓存
type ChainInfoReader interface {
	// 获取白名单
	GetPeerList(num int) ([]model.CorePeer, error)
	// 获取链上唯一id
	GetUserCode() (string, error)
	// 获取指定pid的节点信息
	GetPeer(id string) (model.CorePeer, error)
}

type Cacher interface {
	//GetCache()(*cache.Cache,error)
}

type Miner interface {
	//ReportContribute(num int64)error
	GetChallenge() (string, error)
	Mining(m dto.MiningDTO) error
	UpdateAddress(addrList []string) error
	Heartbeat() error
	GetFileList(n int64) ([]string, error)
	GetStoreChallenge() (string, error)
	GetChallengeStage() (int64, string, int64, error)

	UpdateOrGen(m dto.StoreProofDTO) error
	Prove(m dto.SVProofDTO) error
}

type Peer interface {
	FileOwner
	FileApplyer
	ChainInfoReader
	//Cacher
	Miner
	InitPeer(peer model.CorePeer) error
	DaemonPeer() ([]string, error)
}
