package cli

import (
	"fmt"
	"github.com/gaw508/albbluegreen"
	"github.com/urfave/cli"
	"log"
)

func GetStatusHandler(log *log.Logger, blueGreenService albbluegreen.BlueGreenService) cli.ActionFunc {
	return func(c *cli.Context) error {
		status, err := blueGreenService.Status()
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to get status: %+v", err), 1)
		}
		log.Printf("Current active service: %s\n", status)
		return nil
	}
}
