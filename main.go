package main

import (
	"CortexXDRConnTest/cmd"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// Inilitiaze cli application
	app := &cli.App{

		Name:     "CortexXDRConnTest",
		Version:  "v0.3",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Aleth Bruce",
				Email: "dev@alethlab.com",
			},
		},
		Usage:     "Test connection to Cortex XDR communication servers and storage buckets",
		UsageText: "CortexXDRConnTest -region [VALUE] -tenant [VALUE] \n\t CortexXDRConnTest -region [VALUE] -tenant [VALUE] -proxy [VALUE] \n Short Form \n\t CortexXDRConnTest -r [VALUE] -t [VALUE] \n\t CortexXDRConnTest -r [VALUE] -t [VALUE] -p [VALUE]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "region",
				Usage:    "region name eg. us|ca|uk",
				Aliases:  []string{"r"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "tenant",
				Usage:    "tanant name",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "proxy",
				Usage:    "proxy IP and Port eg. 127.0.0.1:8888",
				Aliases:  []string{"p"},
				Required: false,
			},
		},
		// Entry point for CLI application
		Action: func(c *cli.Context) error {
			//flagLogToFile := false

			if c.String("proxy") != "" {
				// Set HTTP to use a proxy when performing connection tests
				proxyURL, err := url.Parse("http://" + c.String("proxy"))
				if err != nil {
					log.Println("Proxy not reachable: \n", err)
					panic(err)
				}
				http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
				// Print to console the proxy selected to be used
				fmt.Println("Using Proxy:", c.String("proxy"))
			}

			// Print to console the region and tenant selected to be tested
			fmt.Println("Using Region:", c.String("region"))
			fmt.Println("Using Tenant:", c.String("tenant")+"\n")

			cortexURIs := cmd.PopulateUrls(c.String("region"), c.String("tenant"))

			cmd.TestUrlsGet(cortexURIs)
			return nil
		},
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
