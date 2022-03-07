package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vuls-saas/licensecheck/license"
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
				typ = license.Java
			case "ruby":
				typ = license.Ruby
			case "python":
				typ = license.Python
			case "nodejs":
				typ = license.Nodejs
			case "go":
				typ = license.Go
			case "rust":
				typ = license.Rust
			case "github":
				typ = license.GitHub
			default:
				return errors.New("please specify option -type in java/ruby/python/nodejs/go/rust/github")
			}
			result, confidence, err := license.Scan(name, version, typ)
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
