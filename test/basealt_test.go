package basealt

import (
    "fmt"
    "testing"

    "basealt/pkg/basealt"
)

func Test_GetPackages(t *testing.T) {
    pkgs := basealt.GetPackages("x86_64", "p10")
    fmt.Print(pkgs)
}

func Test_CompareExistence(t *testing.T) {
    pkgsA := basealt.Pkgs{
        "value1": basealt.Package{},
        "value2": basealt.Package{},
        "value3": basealt.Package{},
        "value4": basealt.Package{},
        "value5": basealt.Package{},
    }

    pkgsB := basealt.Pkgs{
        "value1": basealt.Package{},
        "value2": basealt.Package{},
    }

    pkgsC := basealt.CompareExistence(pkgsA, pkgsB)

    got := len(pkgsC)
    want := 3
    if got != want {
        t.Errorf("len(pkgC) = %d; want %d\n%v", got, want, pkgsC)
    }
}
