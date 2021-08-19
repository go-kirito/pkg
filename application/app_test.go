package application

import (
	"testing"
	"time"

	"github.com/go-kirito/pkg/transport/grpc"
	"github.com/go-kirito/pkg/transport/http"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	gs := grpc.NewServer()
	app := New(
		Name("kirito"),
		Version("v0.0.1"),
		Server(hs, gs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
