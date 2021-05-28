package node

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/ipfs/go-ipfs/core"
)

func FastBootstrap(n *core.IpfsNode) {
	numPeers := 0
	fmt.Println("fast bootstrapping ...")
	for numPeers < 1000 {
		ps := n.Peerstore.Peers()
		numPeers = len(ps)
		time.Sleep(100 * time.Millisecond)
		go func() {
			dht := n.DHT
			rval := make([]byte, 32)
			rand.Read(rval)
			ctx, cancel := context.WithTimeout(
				context.Background(), 2*time.Second)
			defer cancel()
			_, err := dht.GetValue(ctx, string(rval))
			if err != nil {
				//fmt.Println("debug err", err)
			}
		}()
	}
	go func() {
		dht := n.DHT
		ctx, cancel := context.WithTimeout(
			context.Background(), 2*time.Second)
		defer cancel()
		_, err := dht.GetValue(ctx, "SAME_TARGET")
		if err != nil {
			//fmt.Println("debug err", err)
		}
	}()
	fmt.Println("fast bootstrapping done.")
}
