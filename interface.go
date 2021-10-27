package standard

type AuthAPI interface {
	Server
	LitePeer
}

type BasicPeer interface {
	ResponseApply(cid string, pid string) (*Apply, error)
	GetServerList() ([]CorePeer, error)
	NewPeer(pid string) (*Peer, error)
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
	AuthJoinApply(pid string, i int) (JoinServerListApply, error)
	JoinServerNet() ([]Peer, error)
	GetApply(peerId string) ([]JoinServerListApply, error)
}

// FilePermissionFinder 文件权限查询相关接口
type FilePermissionFinder interface {
	// FindFileAuth 获取文件权限信息
	FindFileAuth(cid string) (*Apply, error)
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
	NewIpfsFilePermission(cid string, state int) (*IpfsFileInfo, error)
	// Share 分享文件
	Share(cid string, pid string, limit Limit) (*IpfsFileInfo, error)

	Change(cid string, ipfs IpfsFileInfo) (*IpfsFileInfo, error)

	Delete(cid string) error
}

type FileOwner interface {
	FilePermissionController
}

type FileApplyer interface {
	// 	ApplyIpfsRemote 远程文件申请
	ApplyIpfsRemote(cid string) (*Apply, error)
	// 	ApplyIpfsRemote 本地文件申请
	ApplyIpfsLocal(cid string) (string, error)
}
