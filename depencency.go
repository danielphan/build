package build

import (
	"crypto/sha1"
	"path/filepath"
	"sort"
)

type dependency interface {
	dependencies() []dependency
	output []string
	signature() []byte
}

// sources.go
type Sources []string

func (s Sources) dependencies() []dependency {
	return nil
}

func (s Sources) output() []string {
	return s.paths()
}

func (s Sources) signature() []byte {
	ps := s.paths()
	hash := sha1.New()
	for _, d := range Dependencies {
		hash.Write(d.Signature())
	}
	return hash.Sum(nil)
}

func (s Sources) paths() (ps []string) {
	for _, p := range s {
		matches, err := filepath.Glob(p)
		if err != nil {
		}
		ps := append(ps, matches)
	}
	return ps
}

// target.go
type Target struct {
	Dependencies []dependency
	Output []string
}

func (t *Target) dependencies() []dependency {
	return t.Dependencies;
}

func (t *Target) output() []string {
	return t.Output
}

func (t *Target) Signature() []byte {
	hash := sha1.New()
	// TODO include metadata in hash
	for _, d := range Dependencies {
		hash.Write(d.Signature())
	}
	return hash.Sum(nil)
}

func (t *Target) Apply(outDir string) error {
	
}

// build/copy/copy.go
package copy

import (
	"build"
)

type Copy struct {
	build.Target
}

func (c *Copy) Apply() error {
	c.(build.Target).Apply()
	
}

// Usage
// build/examples/copy/build.go
package copy

import (
	"github.com/danielphan/build"
	"github.com/danielphan/build/copy"
)

const package = "example/copy"

sources := build.Sources{
	package + "*.txt",
}

Copy := copy.Target{
	Dependencies: {
		sources,
	},
}

func init() {
	build.Add(package, "Copy", &Copy)
}

// build/main.go
package main

import (
	"flags"
)

type package map[string]*Target
var packages map[string]package
var targetPackageNames map[*Target]string
var dependees map[*Target][]*Target

func main() {
	flags.Parse()
	packageName := flags.Arg(0)
	targetName := flags.Arg(1)
	
	package := packages[packageName]
	if package == nil {
	
	}
	target := targetName[t]
	if target == nil {
	
	}
	
	targets, err := topSort(target)
	if err != nil {
	}
	for _, t := range targets {
		err := t.Apply()
	}
}
