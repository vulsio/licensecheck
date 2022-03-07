# licensecheck

licensecheck is tool to detect license of OSS.

It supports java/ruby/python/nodejs/go/rust/github.

# feature

The purpose of this tool is to collect license information without actually installing the software.

Instead of exploring the file tree, this tool explores information published on the Internet.

## Usage

```
NAME:
   licensecheck - License Checker of OSS

USAGE:
   licensecheck [global options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value, -n value     --name { name like go }
   --version value, -v value  --version { version like 1.0.0}
   --type value, -t value     -type { java|ruby|python|nodejs|go|github }
   --help, -h                 show help (default: false)
```

## example
```
$ licensecheck --name rails --version 7.0.1 --type ruby
Licnese: MIT, Confidense: 100%
```

## as a Package

```main.go
import (
	"github.com/vuls-saas/licensecheck"
	"github.com/vuls-saas/licensecheck/core/java"
)

func detect(name, version) {
	name = java.ToMavenPomName(name)
	result, confidence, err := new(licensecheck.Scanner).Scan(name, version, license.Java)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Licnese: %s, Confidense: %d\n", result, int(confidence*100))
}
```

```main.go
import	"github.com/vuls-saas/licensecheck/core/python"

func detectWithMinimumImport(name, version) {
	result, confidence, err := new(python.Scanner).ScanLicense(name, version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Licnese: %s, Confidense: %d\n", result, int(confidence*100))
}
```

## Scan Data Source

Information of License will be fetched Data Sources below.

| target | data source                       |
| ------ | --------------------------------- |
| Java   | https://repo1.maven.org           |
| Ruby   | https://rubygems.org              |
| Python | https://pypi.org                  |
| Nodejs | https://registry.npmjs.org        |
| Go     | https://pkg.go.dev                |
| Rust   | https://crates.io                 |
| GitHub | https://raw.githubusercontent.com |
