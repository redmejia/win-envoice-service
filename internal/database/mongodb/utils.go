package mongodb

import (
	"context"
	"fmt"
	"time"
)

func (m *MongoDB) InsertOneCollec() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	defer m.Mdb.Disconnect(ctx)
	fmt.Println(databaseName, collection)
	return
}
