package exemples

import (
	"context"
	"fmt"
	"time"

	"github.com/jcarlos1999/Golang/Learning/Neo4j/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SimpleQuery(ctx context.Context, session neo4j.SessionWithContext) {
	result, err := session.Run(ctx,
		`MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
	 RETURN p`, map[string]any{"title": "The Matrix"}, func(tc *neo4j.TransactionConfig) {
			tc.Timeout = 3 * time.Second
		})
	utils.PanicOnErr(err)

	people, err := neo4j.CollectTWithContext[neo4j.Node](ctx, result,
		func(record *neo4j.Record) (neo4j.Node, error) {
			person, _, err := neo4j.GetRecordValue[neo4j.Node](record, "p")
			return person, err
		})
	utils.PanicOnErr(err)
	fmt.Println("Person 0: ", people[0])
	fmt.Println("Properties: ", people[0].Props)
	fmt.Println("Elemente id: ", people[0].ElementId)
	fmt.Println("Labels: ", people[0].Labels)
	fmt.Println("ID: ", people[0].GetId())
	fmt.Println("Person 1: ", people[1])
	fmt.Println("Properties: ", people[1].Props)
	fmt.Println("Elemente id: ", people[1].ElementId)
	fmt.Println("Labels: ", people[1].Labels)
	fmt.Println("ID: ", people[1].GetId())

}

func FindTheDirector(ctx context.Context, session neo4j.SessionWithContext, title string) {

	cypher := ` MATCH (p:Person)-[:DIRECTED]->(:Movie {title: $title})
	RETURN p.name AS Director`
	params := map[string]any{"title": title}

	personNode, err := neo4j.ExecuteRead[[]any](ctx, session, func(tx neo4j.ManagedTransaction) ([]any, error) {
		result, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return *new([]any), err
		}

		var record []any
		for result.Next(ctx) {
			record = append(record, result.Record().AsMap())
		}
		return record, nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(personNode)
}

func GetNodesAndRelationships(ctx context.Context, session neo4j.SessionWithContext, title string) ([]any, error) {
	cypher := `MATCH path = (p:Person)-[r:ACTED_IN]->(m:Movie {title: $title}) RETURN p, r, m, path`
	params := map[string]any{"title": title}

	result, err := session.Run(ctx, cypher, params)
	if err != nil {
		return nil, err
	}

	}
}
