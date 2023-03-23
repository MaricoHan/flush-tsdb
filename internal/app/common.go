package app

import (
	"github.com/jony-lee/go-progress-bar"
)

var ProcessBarOptions = []func(bar *progress.Bar){
	progress.WithFillerLength(50),
	progress.WithFiller(">"),
	progress.WithTimeFormat("2006-01-02 15:04:05"),
}
