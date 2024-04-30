package main

import (
	"context"
	"log"
	"os"

	b64 "encoding/base64"
	"encoding/json"

	"github.com/orbit-ops/rocket-google/controller"
	cli "github.com/urfave/cli/v2"
)

type Request struct {
	Target string
	Person string
}

func main() {
	app := &cli.App{
		Name:  "rocket-google",
		Usage: "Manage privilege elevation in google for mission-control",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					arg := cCtx.Args().First()
					sDec, _ := b64.StdEncoding.DecodeString(arg)
					var req *Request
					err := json.Unmarshal(sDec, &req)
					if err != nil {
						return err
					}

					c, err := controller.NewGoogleController(context.TODO())
					if err != nil {
						return err
					}

					return c.AddPersonToGroup(req.Person, req.Target)
				},
			},
			{
				Name:  "delete",
				Usage: "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					arg := cCtx.Args().First()
					sDec, _ := b64.StdEncoding.DecodeString(arg)
					var req *Request
					err := json.Unmarshal(sDec, &req)
					if err != nil {
						return err
					}

					c, err := controller.NewGoogleController(context.TODO())
					if err != nil {
						return err
					}

					return c.RemoveUserFromGroup(req.Person, req.Target)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
