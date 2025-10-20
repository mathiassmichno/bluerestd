package main

import (
	"context"
	"os"

	"github.com/bluetuith-org/bluerestd/cmd/app"
)

func main() {
	app.New().Run(context.Background(), os.Args)
}
