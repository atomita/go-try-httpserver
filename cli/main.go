package main

import (
	"github.com/atomita/go-try-httpserver"
	"github.com/jteeuwen/go-pkg-optarg"
)

func main() {
	optarg.Header("start")
	optarg.Add("P", "port", "port number.", 80)

	ch := optarg.Parse()
	<-ch

	if len(optarg.Remainder) == 1 {
		cmd := optarg.Remainder[0]
		switch cmd {
		case "start":
			start()
		}
	} else {
		optarg.Usage()
	}
}

func start() {
	var port int64
	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "P":
			port = opt.Int64()
		}
	}
	tryHttpserver.Start(port)
}
