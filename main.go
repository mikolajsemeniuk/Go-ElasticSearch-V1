package main

import (
	"github.com/mikolajsemeniuk/go-react-elasticsearch/application"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/data"
)

func main() {
	data.GetInfo()
	application.Listen()
}
