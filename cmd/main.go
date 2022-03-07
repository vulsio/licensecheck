package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vuls-saas/licensecheck"
)

func main() {
	app := &cli.App{
		Name:      "licensecheck",
		Usage:     "License Checker of OSS",
		UsageText: "licensecheck [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Usage:   "--name { name like go }",
				Aliases: []string{"n"},
			},
			&cli.StringFlag{
				Name:    "version",
				Usage:   "--version { version like 1.0.0}",
				Aliases: []string{"v"},
			},
			&cli.StringFlag{
				Name:     "type",
				Usage:    "-type { java|ruby|python|nodejs|go|github }",
				Required: true,
				Aliases:  []string{"t"},
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			version := c.String("version")
			var typ int
			switch c.String("type") {
			case "java":
				typ = licensecheck.Java
			case "ruby":
				typ = licensecheck.Ruby
			case "python":
				typ = licensecheck.Python
			case "nodejs":
				typ = licensecheck.Nodejs
			case "go":
				typ = licensecheck.Go
			case "rust":
				typ = licensecheck.Rust
			case "github":
				typ = licensecheck.GitHub
			default:
				return errors.New("please specify option -type in java/ruby/python/nodejs/go/rust/github")
			}
			result, confidence, err := new(licensecheck.Scanner).Scan(name, version, typ)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Licnese: %s, Confidense: %d%%\n", result, int(confidence*100))
			return nil
		},
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
