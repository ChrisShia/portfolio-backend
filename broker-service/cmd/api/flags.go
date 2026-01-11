package main

import "flag"

func (c *config) flags() {
	flag.IntVar(&c.port, "p", 8080, "http port")
	flag.StringVar(&c.mongo.URI, "mongo", "mongodb://localhost:27017", "mongodb uri")

	flag.Parse()
}
