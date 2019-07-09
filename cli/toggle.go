package cli

import (
	"fmt"
	"github.com/gaw508/albbluegreen"
	"github.com/urfave/cli"
)

func ToggleHandler(blueGreenService albbluegreen.BlueGreenService) cli.ActionFunc {
	return func(c *cli.Context) error {
		newStatus, err := blueGreenService.Toggle()
		if err != nil {
			return cli.NewExitError(fmt.Sprintf("Failed to toggle status: %+v", err), 1)
		}
		fmt.Printf("Successfully toggled to: %s\n", newStatus)
		return nil
	}
}
