package _interface

import "github.com/ipfs/go-ipfs-auth/standard/model"

type AuthAPI interface {
	Server
	LitePeer
}

type BasicPeer interface {
	ResponseApply(cid string, pid string) (*model.Apply, error)
	GetServerList() ([]model.CorePeer, error)
	NewPeer(pid string) (*model.Peer, error)
}

type ServerPeer interface {
	BasicPeer
	Server
}

type LitePeer interface {
	BasicPeer
	FileOwner
	FileApplyer
	FilePermissionFinder
}

type Server interface {
	ApplyJoinServerNet(addresses []string) error
	AuthJoinApply(pid string, i int) (model.JoinServerListApply, error)
	JoinServerNet() ([]model.Peer, error)
	GetApply(peerId string) ([]model.JoinServerListApply, error)
}

// FilePermissionFinder 文件权限查询相关接口
type FilePermissionFinder interface {
	// FindFileAuth 获取文件权限信息
	FindFileAuth(cid string) (*model.Apply, error)
	// FindFileList 获取已拥有文件列表
	FindFileList(pid string) ([]string, error)
	// FindFileReaderList 获取文件阅读者列表
	FindFileReaderList(cid string) ([]string, error)
	// FindPeerReadingHistory 获取节点阅读历史
	FindPeerReadingHistory() ([]string, error)
}

// FilePermissionController 文件权限操作相关接口
type FilePermissionController interface {
	// NewIpfsFilePermission 新增ipfs文件权限信息
	NewIpfsFilePermission(cid string, state int) (*model.IpfsFileInfo, error)
	// Share 分享文件
	Share(cid string, pid string, limit model.Limit) (*model.IpfsFileInfo, error)

	Change(cid string, ipfs model.IpfsFileInfo) (*model.IpfsFileInfo, error)

	Delete(cid string) error
}

type FileOwner interface {
	FilePermissionController
}

type FileApplyer interface {
	// 	ApplyIpfsRemote 远程文件申请
	ApplyIpfsRemote(cid string) (*model.Apply, error)
	// 	ApplyIpfsRemote 本地文件申请
	ApplyIpfsLocal(cid string) (string, error)
}
