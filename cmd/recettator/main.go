package main

import (
	"fmt"
	"os"
	"time"

<<<<<<< Updated upstream
	"github.com/sirupsen/logrus"
=======
	"github.com/camembertaulaitcrew/recettator"
>>>>>>> Stashed changes
	"github.com/urfave/cli"

	"ultre.me/recettator"
)

func main() {
	app := cli.NewApp()
	app.Name = "recettator"
	app.Usage = "Generate CALC recipes"
	app.Version = "master"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "seed, s",
			Usage: "Set seed value",
		},
		cli.IntFlag{
			Name:  "main-ingredients",
			Usage: "Amount of main-ingredients",
		},
		cli.IntFlag{
			Name:  "secondary-ingredients",
			Usage: "Amount of secondary-ingredients",
		},
		cli.IntFlag{
			Name:  "steps",
			Usage: "Amount of steps",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Use JSON output",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		//panic(err)
		panic(err)
	}
}

func run(c *cli.Context) error {
	seed := int64(c.Int("seed"))
	if seed == 0 {
		seed = time.Now().UTC().UnixNano()
	}
	rctt := recettator.New(seed)

	rctt.SetSettings(recettator.Settings{
		MainIngredients:      uint64(c.Int("main-ingredients")),
		SecondaryIngredients: uint64(c.Int("secondary-ingredients")),
		Steps:                uint64(c.Int("steps")),
	})

	var output string
	var err error

	if c.Bool("json") {
		output = rctt.JSON()
	} else {
		output, err = rctt.Markdown()
		if err != nil {
			//panic(err)
			return err
		}
	}

	fmt.Println(output)

	return nil
}
