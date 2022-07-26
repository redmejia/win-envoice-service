package mongodb

import (
	"context"
	"testing"
)

func TestDBConnection(t *testing.T) {

	client, _ := Connection()

	err := client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("connected")

}
