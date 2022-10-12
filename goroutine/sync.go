package main

import "sync"

var resource map[int]int
var loadResource sync.Once

func main() {

}
func loadRe() {
	loadResource.Do(func() {
		resource[1] = 1
	})
}
