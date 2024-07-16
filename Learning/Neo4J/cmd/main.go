package main

import (
	"context"

	"github.com/jcarlos1999/Golang/Learning/Neo4j/exemples"
	"github.com/jcarlos1999/Golang/Learning/Neo4j/initializers"
	"github.com/jcarlos1999/Golang/Learning/Neo4j/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()

	driver, err := initializers.InitDriver(ctx, "neo4j://localhost:7687", "neo4j", "password")
	utils.PanicOnErr(err)
	defer driver.Close(ctx)

	session, err := initializers.CrateSessions(ctx, driver, neo4j.AccessModeWrite)
	utils.PanicOnErr(err)
	defer session.Close(ctx)

	// exemples.SimpleQuery(ctx, session)
	exemples.FindTheDirector(ctx, session, "The Matrix")
}
