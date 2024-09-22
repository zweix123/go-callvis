// 1. golang source code -> ast -> ssa
package go_callvis_modern

import (
	"fmt"

	"github.com/zweix123/go-callvis/pkg/util"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func (g *GoCallvisModern) parse() error {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
		Tests:      g.includeTest,
		Dir:        g.dir,
		BuildFlags: util.GetBuildFlags(),
	}
	pkgs, err := packages.Load(cfg, g.packagesLoadPatterns...)
	if err != nil {
		return fmt.Errorf("failed to load packages: %w", err)
	}

	if packages.PrintErrors(pkgs) > 0 {
		return fmt.Errorf("packages contain errors")
	}

	// Create and build SSA-form program representation.
	g.program, g.pkgs = ssautil.AllPackages(pkgs, 0) //? parameter meaning?
	g.program.Build()

	g.mains, err = getMainPackages(g.pkgs)
	if err != nil {
		return fmt.Errorf("failed to get main packages: %w", err)
	}

	return nil
}

func getMainPackages(pkgs []*ssa.Package) (mains []*ssa.Package, err error) {
	for _, p := range pkgs {
		if p != nil && p.Pkg.Name() == "main" && p.Func("main") != nil { // 在main包内且存在名为main的函数
			mains = append(mains, p)
		}
	}
	if len(mains) == 0 {
		err = fmt.Errorf("no main packages")
	}
	return
}
