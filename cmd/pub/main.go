package main 

import (
	"fmt"
	"os"
	"time"
	"github.com/georgercarder/pubsub_test/node"

	"github.com/ipfs/go-ipfs/core"
)

func main() {
	for node.G_Node() == nil {
		time.Sleep(100*time.Millisecond)
	}
        // fast bootstrap
	node.FastBootstrap((*core.IpfsNode)(node.G_Node()))
        ps := node.G_Node().Peerstore.Peers()
	fmt.Println("peers", len(ps))
	b := make([]byte, 1)
	for {
		fmt.Println("Enter any key to publish")
		os.Stdin.Read(b)
		node.Publish("cats", []byte("smell good"))
		fmt.Println("debug published")
	}
}
