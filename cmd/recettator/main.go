package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/camembertaulaitcrew/recettator"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "recettator"
	app.Usage = "Generate CALC recipes"
	app.Version = "master"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.IntFlag{
			Name:  "seed, s",
			Usage: "Set seed value",
		},
		cli.IntFlag{
			Name:  "ingredients, i",
			Usage: "Amount of ingredients",
		},
		cli.IntFlag{
			Name:  "steps",
			Usage: "Amount of steps",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("%v", err)
	}
}

func run(c *cli.Context) error {
	rctt := recettator.New(uint64(c.Int("seed")))

	fmt.Println(rctt)

	return nil
}
