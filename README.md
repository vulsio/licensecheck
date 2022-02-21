# licnese-scanner

license-scanner is tool to detect license of OSS.

It suprots java/ruby/python/nodejs/go/github.

## Usage

```
NAME:
   license-scanner - License Scanner of OSS

USAGE:
   license-scanner [global options]

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
$ license-scanner --name rails --version 7.0.1 --type ruby
Licnese: MIT, Confidense: 100%
```

## as a Package

```main.go
import (
	"github.com/vuls-saas/license-scanner/license"
	"github.com/vuls-saas/license-scanner/license/java"
)

func detect(name, version) {
	name = java.ToMavenPomName(name)
	result, confidence, err := license.Scan(name, version, license.Java)
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
| Go     | https://raw.githubusercontent.com |
| GitHub | https://raw.githubusercontent.com |

## LICENSE

MIT License

Copyright (c) 2022 FutureVuls

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
