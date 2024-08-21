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

func Test_ComparePackagesExistence(t *testing.T) {
    A := basealt.Pkgs{
        "value1": basealt.Package{},
        "value2": basealt.Package{},
        "value3": basealt.Package{},
        "value4": basealt.Package{},
        "value5": basealt.Package{},
    }

    B := basealt.Pkgs{
        "value1": basealt.Package{},
        "value2": basealt.Package{},
    }

    C := basealt.ComparePackagesExistence(A, B)

    got := len(C)
    want := 3
    if got != want {
        t.Errorf("len(C) = %d; want %d\n%v", got, want, C)
    }
}

func Test_ComparePackagesVersion(t *testing.T) {
    A := basealt.Pkgs{
        "value1": basealt.Package{Name: "ArxLibertatis", Release: "alt2"},
        "value2": basealt.Package{Name: "ArxLibertatis", Release: "alt1.20200607.1"},
        "value3": basealt.Package{},
    }

    B := basealt.Pkgs{
        "value1": basealt.Package{Name: "ArxLibertatis", Release: "alt1.20200607.1"},
        "value2": basealt.Package{Name: "ArxLibertatis", Release: "alt2"},
    }

    C := basealt.ComparePackagesVersion(A, B)

    got := len(C)
    want := 1
    if got != want {
        t.Errorf("len(C) = %d; want %d\n%v", got, want, C)
    }

    gotRelease := C["value1"].Release
    wantRelease := A["value1"].Release
    if gotRelease != wantRelease {
        t.Errorf("C[\"value1\"].release = %s; want %s\n%v", gotRelease, wantRelease, C)
    }
}

