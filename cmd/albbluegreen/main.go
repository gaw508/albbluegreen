package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	bgaws "github.com/gaw508/albbluegreen/aws"
	bgcli "github.com/gaw508/albbluegreen/cli"
	"github.com/urfave/cli"
	"log"
	"os"
)

const version = "0.1.0"

func main() {
	app := cli.NewApp()

	app.Name = "albbluegreen"
	app.Version = version
	app.Usage = "tool for creating blue green deployments using AWS EC2 ALB"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "aws-region",
			Usage:  "AWS region to use",
			EnvVar: "BG_AWS_REGION",
		},
		cli.StringFlag{
			Name:   "aws-profile",
			Usage:  "AWS profile to use",
			EnvVar: "BG_AWS_PROFILE",
		},
		cli.StringFlag{
			Name:   "rule-arn",
			Usage:  "ARN of the mutable 'green' rule to use",
			EnvVar: "BG_RULE_ARN",
		},
	}

	bgService := &bgaws.BlueGreenService{}
	app.Before = func(c *cli.Context) error {
		var region *string
		if r := c.GlobalString("aws-region"); r != "" {
			region = aws.String(r)
		}

		awsSession, err := session.NewSessionWithOptions(session.Options{
			Config:  aws.Config{Region: region},
			Profile: c.GlobalString("aws-profile"),
		})
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		bgService.RuleArn = c.GlobalString("rule-arn")
		bgService.AlbClient = elbv2.New(awsSession)

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "toggle",
			Aliases: []string{"t"},
			Usage:   "toggle the active target group for blue green deployment",
			Action:  bgcli.ToggleHandler(bgService),
		},
		{
			Name:    "get-status",
			Aliases: []string{"g"},
			Usage:   "get the active target group for blue green deployment",
			Action:  bgcli.GetStatusHandler(bgService),
		},
		{
			Name:    "set-status",
			Aliases: []string{"s"},
			Usage:   "set the active target group for blue green deployment",
			Action:  bgcli.SetStatusHandler(bgService),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "status",
					Usage: "the status to set 'blue' or 'green'",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
