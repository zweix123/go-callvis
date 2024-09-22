// 2. ssa -> callgraph
package go_callvis_modern

import (
	"fmt"

	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/callgraph/rta"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/pointer"
	"golang.org/x/tools/go/ssa"
)

func (g *GoCallvisModern) analysis() error {
	switch g.algo {
	case Static:
		g.callgraph = static.CallGraph(g.program)
	case Cha:
		g.callgraph = cha.CallGraph(g.program)
	case Rta:
		mainFuncs, err := getMainFunction(g.mains)
		if err != nil {
			return fmt.Errorf("getMainFunction: %w", err)
		}
		g.callgraph = rta.Analyze(mainFuncs, true).CallGraph //? parameter meaning?
	case Pointer:
		config := &pointer.Config{
			Mains:          g.mains,
			BuildCallGraph: true,
		}
		result, err := pointer.Analyze(config)
		if err != nil {
			return fmt.Errorf("pointer.Analyze: %w", err)
		}
		g.callgraph = result.CallGraph
	default:
		return fmt.Errorf("unknown analysis algorithm: %s", g.algo)
	}
	return nil
}

func getMainFunction(mains []*ssa.Package) ([]*ssa.Function, error) {
	var mainFuncs []*ssa.Function
	for _, pkg := range mains {
		mainFunc := pkg.Func("main")
		if mainFunc != nil {
			mainFuncs = append(mainFuncs, mainFunc)
		}
	}
	return mainFuncs, nil
}
