package main

import (
	"fmt"
	"io"
	"os"
	"encoding/xml"
)

type Dependency struct {
	GroupId string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version string `xml:"version"`
}

type Dependencies struct {
	Dependencies []Dependency `xml:"dependency"`
}

type DependencyManagement struct {
	DependencyManagement xml.Name `xml:"dependencyManagement"`
	Dependencies Dependencies `xml:"dependencies"`
}

func main() {
	content := make([]byte, 0)
	buf :=make([]byte, 64)
	for {
		count, err := io.ReadFull(os.Stdin, buf)
		if err == nil || err == io.ErrUnexpectedEOF {
			for i := 0; i < count; i += 1 {
				content = append(content, buf[i])
			}
		}
		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		}
	}
	fmt.Printf("Read %d bytes from stdin.\n", len(content))
	var dm DependencyManagement
	xml.Unmarshal(content, &dm)
	for _,dep := range dm.Dependencies.Dependencies {
		fmt.Printf("%s:%s:%s\n", dep.GroupId, dep.ArtifactId, dep.Version)
	}
}
