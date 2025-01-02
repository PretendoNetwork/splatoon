package main

import (
	nex2 "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/splatoon/globals"
	"sync"
	"time"

	"github.com/PretendoNetwork/splatoon/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server
	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()
	go startSelfTesting()

	wg.Wait()
}

func selfTest() {
	globals.Logger.Info("Self-testing...")
	var errors = 0

	var seenPids map[uint64]struct{}
	// apparently we hold a lock all the way thru this
	globals.SecureEndpoint.Connections.Each(func(key string, pc *nex2.PRUDPConnection) bool {
		if pc.PID() == nil || pc.PID().Value() == 0 {
			if pc.ConnectionState == nex2.StateConnected {
				globals.Logger.Warningf("PID connection invariant failed: %v %#v", key, pc)
				errors++
			}
			// nil entry is ok but kinda weird
			return false
		}

		pid := pc.PID().Value()

		// expected invariant: valid connections do not have PIDs
		if pc.ConnectionState != nex2.StateConnected {
			globals.Logger.Warningf("Stale connection invariant failed: %v %#v", key, pc)
			errors++
		}
		if _, ok := seenPids[pid]; ok {
			globals.Logger.Warningf("Duplicate connection for PID: %v %#v", key, pc)
			errors++
		}
		seenPids[pid] = struct{}{}

		// check SocketConnections
		var found = false
		globals.SecureServer.Connections.Each(func(key string, sc *nex2.SocketConnection) bool {
			return sc.Connections.Each(func(key uint8, pc2 *nex2.PRUDPConnection) bool {
				if pc2 == pc {
					if found {
						globals.Logger.Warningf("Duplicate SocketConnection: %v %#v", key, pc)
						errors++
					}
					found = true
					return true
				}
				return false
			})
		})

		if !found {
			globals.Logger.Warningf("Connection has no SocketConnection: %v %#v", key, pc)
			errors++
		}

		return false
	})

	globals.Logger.Infof("Self-test finished with %v errors", errors)
}

func startSelfTesting() {
	for range time.Tick(10 * time.Second) {
		selfTest()
	}
}
