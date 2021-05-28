package node

import (
	"crypto/rand"
        "crypto/sha256"

	"github.com/ethereum/go-ethereum/common/hexutil"

        "github.com/ipfs/go-cid"
        mh "github.com/multiformats/go-multihash"
)

func UniqueId() (string) {
	b := make([]byte, 32)	
	_, err := rand.Read(b)
	if err != nil {
		// TODO log
	}
	return hexutil.Encode(b)
}


func String2CID(key string) (c cid.Cid, err error) {
        h := sha256.New()
        h.Write([]byte(key))
        hashed := h.Sum(nil)
        multihash, err := mh.Encode([]byte(hashed), mh.SHA2_256)
        if err != nil {
                return
        }
        c = cid.NewCidV0(multihash)
        return
}

