package algorithm

import (
	"encoding/base64"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	mh "github.com/multiformats/go-multihash"
	"math/bits"
)

const HashLen = 32

type Hash [HashLen]byte

func (h *Hash) String() string {
	return base64.StdEncoding.EncodeToString(h[:])
}

func GetFromString(s string) (*Hash, error) {
	res := &Hash{}
	err := res.LoadString(s)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *Hash) LoadString(s string) error {
	decode, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	if len(decode) != HashLen {
		return fmt.Errorf("wrong size")
	}
	copy(h[:], decode)
	return nil
}

// CommonPrefixLen 计算两个32位byte数组的共同前导
func CommonPrefixLen(a, b []byte) (int, error) {
	if len(a) != 32 || len(b) != 32 {
		return 0, fmt.Errorf("前导零计算，位数错误")
	}
	for i := 0; i < 32; i++ {
		c := a[i] ^ b[i]
		if c != 0 {
			return i*8 + bits.LeadingZeros8(uint8(c)), nil
		}
	}
	return 32 * 8, nil
}

func GetCidHash(c cid.Cid) ([]byte, error) {
	hash := c.Hash()
	decode, err := mh.Decode(hash)
	if err != nil || decode.Length != 32 {
		return nil, err
	}
	return decode.Digest, nil
}

func GetCidHashFromString(cStr string) ([]byte, error) {
	c, err := cid.Decode(cStr)
	if err != nil {
		return nil, err
	}
	return GetCidHash(c)
}

func GetPidByte(pid peer.ID) []byte {
	d, _ := mh.Decode(mh.Multihash(pid))
	pidByte := d.Digest
	return pidByte
}

func GetPidByteFromString(pidStr string) ([]byte, error) {
	pid, err := peer.Decode(pidStr)
	if err != nil {
		return nil, err
	}
	d, _ := mh.Decode(mh.Multihash(pid))
	pidByte := d.Digest
	return pidByte, nil
}
