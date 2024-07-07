package main

import (
	"go-zero-dandan/common/fmtd"
	"go-zero-dandan/pkg/filed"
	"go-zero-dandan/pkg/parsed"
)

func main() {

	api, err := parsed.ParseGoZeroApiByFile("./plat.api")
	if err != nil {
		fmtd.Error(err)
	}
	filed.JsonFile(api, "api")
}
