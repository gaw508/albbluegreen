package cli

import (
	"fmt"
	"github.com/gaw508/albbluegreen"
	"github.com/urfave/cli"
)

func SetStatusHandler(blueGreenService albbluegreen.BlueGreenService) cli.ActionFunc {
	return func(c *cli.Context) error {
		status := c.String("status")
		if status == "" {
			return cli.NewExitError("status is required", 1)
		}

		err := blueGreenService.SetStatus(albbluegreen.BlueGreenStatus(status))
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to set status: %+v", err), 1)
		}
		fmt.Printf("Successfully set service to: %s\n", status)
		return nil
	}
}
