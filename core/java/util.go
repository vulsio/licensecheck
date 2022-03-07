package java

import (
	"fmt"
	"strings"
)

// ToMavenPomName is helper method to build name exists in https://repo1.maven.org
//
// example
// name: org.apache.logging.log4j:log4j-core
// version: 2.17.1
// output: org/apache/logging/log4j/log4j-core/2.17.1/log4j-core-2.17.1.pom
func ToMavenPomName(name, version string) string {
	name = strings.ReplaceAll(name, ":", "/")
	name = strings.ReplaceAll(name, ".", "/")
	split := strings.Split(name, "/")
	softwareName := split[len(split)-1]
	return fmt.Sprintf("%s/%s/%s-%s.pom", name, version, softwareName, version)
}
