package main

import (
	"fmt"
	"time"

	"github.com/georgercarder/pubsub_test/node"
	
	"github.com/georgercarder/alerts"

	"github.com/ipfs/go-ipfs/core"
)

func main() {
	for node.G_Node() == nil {
		time.Sleep(100*time.Millisecond)
	}
	subs := "cats"
	alertName := subs + "Alert"
	fmt.Println("debug Subscribe")
	time.Sleep(1*time.Second) // to allow some connections to be made
	node.Subscribe(subs, alertName)
	fmt.Println("debug subs, alertName", subs, alertName)

        // fast bootstrap
	node.FastBootstrap((*core.IpfsNode)(node.G_Node()))
        ps := node.G_Node().Peerstore.Peers()
	fmt.Println("peers", len(ps))
	alertSub, err := alerts.G_Alerts().NewSubscription(alertName)
	if err != nil {
		// TODO handle err
	}
	for node.G_Node() == nil {
		time.Sleep(100*time.Millisecond)
	}
	for {
		select {
		case a := <- alertSub:
			fmt.Println("debug a", a)
			//handleSub(a.([]byte))
		}
	}

}
