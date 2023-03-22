package main

import (
	"flush-tsdb/cmd"
	_ "github.com/taosdata/driver-go/v3/taosSql"
)

func main() {
	cmd.Execute()
}
