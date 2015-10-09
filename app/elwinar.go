package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	var app = cli.NewApp()

	app.Name = "elwinar"
	app.Version = "4"
	app.Author = "Romain Baugue"
	app.Email = "romain.baugue@elwinar.com"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Usage:  "listening port",
			EnvVar: "ELWINAR_PORT",
		},
		cli.StringFlag{
			Name:   "base",
			Usage:  "base url",
			EnvVar: "ELWINAR_BASE",
		},
		cli.StringFlag{
			Name:   "database",
			Usage:  "sqlite database file",
			EnvVar: "ELWINAR_DATABASE",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "encryption secret",
			EnvVar: "ELWINAR_SECRET",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "administrative password",
			EnvVar: "ELWINAR_PASSWORD",
		},
		cli.StringFlag{
			Name:   "public",
			Usage:  "public directory",
			EnvVar: "ELWIANR_PUBLIC",
		},
		cli.StringFlag{
			Name:   "views",
			Usage:  "views repository",
			EnvVar: "ELWINAR_VIEWS",
		},
		cli.StringFlag{
			Name:   "google-analytics-id",
			Usage:  "google analytics tracking id",
			EnvVar: "ELWINAR_GOOGLE_ANALYTICS_ID",
		},
	}

	app.Before = Bootstrap
	app.Action = Run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
