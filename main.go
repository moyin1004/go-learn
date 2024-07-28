package main

import (
	redisaction "go-learn/redis"
)

func main() {
	// design.TestMap()
	// design.TestPipeline()
	redisaction.ArtleVote("moyin", "text:1")
}
