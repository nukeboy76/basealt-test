package main

import (
    "encoding/json"
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
	flag.StringVar(&args.BranchA, "A", "p10", "target branch A")
	flag.StringVar(&args.BranchB, "B", "sisyphus", "target branch B")
	flag.StringVar(&args.CompareMode, "compare", "existence", "compare mode")
	flag.Parse()

    A := basealt.GetPackages(args.Arch, args.BranchA)
    B := basealt.GetPackages(args.Arch, args.BranchB)
    C := make(basealt.Pkgs)

    switch args.CompareMode {
        case "existence":
            C = basealt.ComparePackagesExistence(A, B)
        case "version":
            C = basealt.ComparePackagesVersion(A, B)
        default:
            fmt.Println("Unexpected value for `compare` mode.")
    }

    res := basealt.NewComparisonResult(args.Arch, C)
    resJSON, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		fmt.Errorf("%v", err)
	}

    fmt.Println(string(resJSON))
}

