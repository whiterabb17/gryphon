package deepfire

import (
	goliath "github.com/whiterabb17/goliath"
)

func GetInterface() string {
	return goliath.ListDevices()
}

func ChumWater(iface string, snapLen string, promisc string, keep bool) {
	goliath.SharkWire(iface, snapLen, promisc, keep)
}
