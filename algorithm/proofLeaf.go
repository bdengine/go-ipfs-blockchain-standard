package algorithm

import (
	"crypto/sha256"
	"github.com/ipfs/go-cid"
)

type ProofLeaf struct {
	C             cid.Cid
	ChallengeHash *Hash
	//PeerChallengeHash *Hash
	*Hash
}

func NewProofLeaf(c cid.Cid, challengeHash *Hash, pidByte []byte) *ProofLeaf {
	pl := ProofLeaf{C: c, ChallengeHash: challengeHash}

	h := sha256.New()
	defer h.Reset()

	h.Write(pidByte)
	h.Write(challengeHash[:])
	hash := h.Sum(nil)
	pl.Hash = GetHashPointer(hash)

	return &pl
}

func GetHashPointer(b []byte) *Hash {
	if len(b) != 32 {
		return nil
	}
	var res Hash
	copy(res[:], b)
	return &res
}

func (pl *ProofLeaf) GetHash(pidByte []byte) *Hash {
	h := sha256.New()
	defer h.Reset()

	h.Write(pidByte)
	h.Write(pl.ChallengeHash[:])
	sum := h.Sum(nil)

	H1 := [32]byte{}
	copy(H1[:], sum)
	hash1 := Hash(H1)

	return &hash1
}

func (h *Hash) Equal(to []byte) bool {
	if len(h) != len(to) {
		return false
	}
	for i, b := range h {
		if b != to[i] {
			return false
		}
	}
	return true
}
