package main

import (
    "flag"
    "fmt"

    "basealt/pkg/basealt"
)

type Args struct {
	Arch string
	BranchA string
	BranchB string
	CompareMode string
}

func main() {
    var args Args
	flag.StringVar(&args.Arch, "arch", "x86_64", "system architecture")
	flag.StringVar(&args.BranchA, "branchA", "p10", "target branch A")
	flag.StringVar(&args.BranchB, "branchB", "sisyphus", "target branch B")
	flag.StringVar(&args.CompareMode, "compare", "existence", "compare mode")
	flag.Parse()

    pkgsA := basealt.GetPackages(args.Arch, args.BranchA)
    pkgsB := basealt.GetPackages(args.Arch, args.BranchB)
    pkgsC := make(basealt.Pkgs)

    switch args.CompareMode {
        case "existence":
            pkgsC = basealt.CompareExistence(pkgsA, pkgsB)
        default:
            fmt.Println("Unexpected value for `compare` mode.")
    }

    for k, _ := range pkgsC {
        fmt.Println(k)
    }
}

