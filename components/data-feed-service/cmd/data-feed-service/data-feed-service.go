package main

import (
	"fmt"
	"os"

	"github.com/chef/automate/components/data-feed-service/cmd/data-feed-service/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
