package go_callvis_modern

import (
	"sync"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type GoCallvisModern struct {
	algo Algorithm

	// packages load config
	dir                  string // analysis dir
	includeTest          bool   // analysis include test file
	packagesLoadPatterns []string

	// analysis result
	program   *ssa.Program
	pkgs      []*ssa.Package
	mains     []*ssa.Package
	callgraph *callgraph.Graph
}

func NewGoCallvisModern() *GoCallvisModern {
	return &GoCallvisModern{}
}

var (
	instance *GoCallvisModern
	once     sync.Once
)

func GetInstance() *GoCallvisModern {
	once.Do(func() {
		instance = NewGoCallvisModern()
	})
	return instance
}

func (g *GoCallvisModern) Process() error {
	// 1. parser: golang source code -> ast -> ssa
	err := g.parse()
	if err != nil {
		return err
	}
	// 2. analyst: ssa -> callgraph
	err = g.analysis()
	if err != nil {
		return err
	}
	// 3. structured: callgraph -> internal representation
	err = g.structure()
	if err != nil {
		return err
	}
	// 4. dot: internal representation -> dot
	err = g.dot()
	if err != nil {
		return err
	}
	return nil
}
