package go_callvis_modern

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zweix123/go-callvis/pkg/log"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type GoCallvisModern struct {
	Algo Algorithm

	// packages load config
	Dir                  string   // analysis dir
	IncludeTest          bool     // analysis include test file
	PackagesLoadPatterns []string //! unused

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
	log.Log("start process")
	// 1. parser: golang source code -> ast -> ssa
	err := g.parse()
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}
	log.Log("parse success")
	// 2. analyst: ssa -> callgraph
	err = g.analysis()
	if err != nil {
		return fmt.Errorf("analysis error: %w", err)
	}
	log.Log("analysis success")
	// 3. structured: callgraph -> internal representation
	err = g.structure()
	if err != nil {
		return fmt.Errorf("structure error: %w", err)
	}
	log.Log("structure success")
	// 4. dot: internal representation -> dot
	err = g.dot()
	if err != nil {
		return fmt.Errorf("dot error: %w", err)
	}
	log.Log("dot success")
	return nil
}

func (g *GoCallvisModern) Argument() string {
	return strings.Join([]string{
		"algo: " + string(g.Algo),
		"dir: " + fmt.Sprintf("%#v", g.Dir),
		"test: " + fmt.Sprintf("%t", g.IncludeTest),
	}, "; ")
}
