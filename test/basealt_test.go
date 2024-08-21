package basealt

import (
    "fmt"
    "testing"

    "basealt/pkg/basealt"
)

func Test_GetBranchBinaryPackages(t *testing.T) {
    packages := basealt.GetBranchBinaryPackages("x86_64", "p10")
    fmt.Print(packages)
}
