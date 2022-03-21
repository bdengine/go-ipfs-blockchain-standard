package dto

import "github.com/bdengine/go-ipfs-blockchain-standard/algorithm"
import "github.com/astaxie/beego/validation"

type MiningDTO struct {
	Cid      string                   `json:"cid" valid:"Required"`
	Address  string                   `json:"address" valid:"Required"`
	Pid      string                   `json:"pid" valid:"Required"`
	SPVProof []*algorithm.MerkleBlock `json:"spvProof" `

	LeadingZero    int    `json:"leadingZero"`
	Challenge      string `json:"challenge" valid:"Required"`
	StoreChallenge string `json:"storeChallenge" valid:"Required"`
	ProofRoot      string `json:"proofRoot" valid:"Required"`
}

func (m *MiningDTO) Valid(v *validation.Validation) {
}

const (
	CodeSuccess    = 1
	CodeFailCommon = 0
)

type ResultDTO struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

type SVProofDTO struct {
	Cid            string                   `json:"cid,omitempty" valid:"Required"`
	Pid            string                   `json:"pid,omitempty" valid:"Required"`
	SvProof        []*algorithm.MerkleBlock `json:"spvProof,omitempty" valid:"Required"`
	StoreChallenge string                   `json:"storeChallenge,omitempty" valid:"Required"`
	ProofRoot      string                   `json:"proofRoot,omitempty" valid:"Required"`
	ChallengeHash  []byte                   `json:"challengeHash,omitempty" valid:"Required"`
	ProofLeaf      *algorithm.Hash          `json:"proofLeaf,omitempty" valid:"Required"`
}

func (svp *SVProofDTO) Valid(v *validation.Validation) {
}

type StoreProofDTO struct {
	ProofRoot      string `json:"proofRoot,omitempty" valid:"Required;"`
	StoreChallenge string `json:"storeChallenge,omitempty" valid:"Required;"`
	PeerId         string `json:"peerId,omitempty" valid:"Required;"`
	PeerAddress    string `json:"peerAddress,omitempty" valid:"Required;"`
}

func (sp *StoreProofDTO) Valid(v *validation.Validation) {
}
