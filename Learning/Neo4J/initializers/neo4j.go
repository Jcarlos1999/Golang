package initializers

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitDriver(ctx context.Context, uri, username, password string) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func CrateSessions(ctx context.Context, driver neo4j.DriverWithContext, mode neo4j.AccessMode) (neo4j.SessionWithContext, error) {

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: mode})

	return session, nil
}
