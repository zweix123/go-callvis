// 1. golang source code -> ast -> ssa
package go_callvis_modern

import (
	"fmt"
	"strings"

	"github.com/zweix123/go-callvis/pkg/util"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func (g *GoCallvisModern) parse() error {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
		Tests:      g.IncludeTest,
		Dir:        g.Dir,
		BuildFlags: util.GetBuildFlags(),
	}
	pkgs, err := packages.Load(cfg, g.PackagesLoadPatterns...)
	if err != nil {
		return fmt.Errorf("failed to load packages: %w", err)
	}

	if packages.PrintErrors(pkgs) > 0 {
		var errors []string
		for _, pkg := range pkgs {
			for _, err := range pkg.Errors {
				errors = append(errors, err.Error())
			}
		}
		return fmt.Errorf("packages contain errors: %s", strings.Join(errors, ";"))
	}

	// Create and build SSA-form program representation.
	g.program, g.pkgs = ssautil.AllPackages(pkgs, 0) //? parameter meaning?
	g.program.Build()

	g.mains = getMainPackages(g.pkgs)

	if g.isMainNecessary() && len(g.mains) == 0 {
		return fmt.Errorf("%s algorithm needs main packages, but no main packages found", g.Algo)
	}

	return nil
}

func getMainPackages(pkgs []*ssa.Package) (mains []*ssa.Package) {
	for _, p := range pkgs {
		if p != nil && p.Pkg.Name() == "main" && p.Func("main") != nil { // 在main包内且存在名为main的函数
			mains = append(mains, p)
		}
	}
	return
}

func (g *GoCallvisModern) isMainNecessary() bool {
	return g.Algo == Rta || g.Algo == Pointer
}
