package build

import (
	"crypto/sha1"
)

type Cacheable interface {
	IsCached() bool // whether the object has already been created
	Get() string // path to cached object
}

type Depender interface {
	Dependencies() []Depender
}

type Sources []string

func (s *Sources) Hash() []byte {
	
}

type Target struct {
	Dependencies []Target
}

func (t *Target) Hash() []byte {
	hash := sha1.New()
	// TODO include metadata in hash
	for _, d := range Dependencies {
		hash.Write(d.Hash())
	}
	return hash.Sum(nil)
}

func init() {
	
}
