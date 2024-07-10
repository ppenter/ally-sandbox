package main

import "github.com/ppenter/ally-sandbox/internal/core/lib/nodejs"
import "C"

//export DifySeccomp
func DifySeccomp(uid int, gid int, enable_network bool) {
	nodejs.InitSeccomp(uid, gid, enable_network)
}

func main() {}
