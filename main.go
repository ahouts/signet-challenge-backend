package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0alpha"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Andrew Houts",
			Email: "ahouts4@gmail.com",
		},
	}
	app.Usage = "Signet Challenge Backend"

	var webPort int
	var scheduleDataFile string
	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "serves schedule data via json api",
			Action: func(c *cli.Context) error {
				schedule, err := NewSchedule(scheduleDataFile)
				if err != nil {
					log.Fatalln(err)
				}
				setupRoutes(schedule)
				serve(webPort)
				return nil
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "p, port",
					Usage:       "`port` to serve on",
					Destination: &webPort,
					Value:       8000,
				},
				cli.StringFlag{
					Name:        "d, data",
					Value:       "./schedule_data.json",
					Usage:       "json data `file` to load",
					Destination: &scheduleDataFile,
				},
			},
		},
	}

	app.Run(os.Args)
}

func serve(port int) {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
