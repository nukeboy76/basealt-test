package basealt

import (
    "encoding/json"
    "fmt"
    "maps"
    "net/http"
    "net/url"
    "slices"

    "github.com/cavaliergopher/rpm"
)

var (
    baseUrl = "https://rdb.altlinux.org/api/"
    endpoint = "/export/branch_binary_packages/"
)

type Package struct {
    Name        string  `json:"name"`
    Epoch       uint32  `json:"epoch"`
    Version     string  `json:"version"`
    Release     string  `json:"release"`
    Arch        string  `json:"arch"`
    Disttag     string  `json:"disttag"`
    Buildtime   uint32  `json:"buildtime"`
    Source      string  `json:"source"`
}

type BranchBinaryPackages struct {
    Packages    []Package   `json:"packages"`
}

type Pkgs map[string]Package

type ComparisonResult struct {
    Arch        string  `json:"arch"`
    Packages    []Package  `json:"packages"`
}

func NewComparisonResult(arch string, pkgs Pkgs) ComparisonResult {
    return ComparisonResult{
        Arch: arch,
        Packages: slices.Collect(maps.Values(pkgs)),
    }
}

func ComparePackagesExistence(a, b Pkgs) Pkgs {
    c := make(Pkgs)
    for k, v := range a {
        if _, ok := b[k]; !ok {
            c[k] = v
        }
    }
    return c
}

func ComparePackagesVersion(a, b Pkgs) Pkgs {
    c := make(Pkgs)
    for k, packageA := range a {
        if packageB, ok := b[k]; (ok && rpm.CompareVersions(packageA.Release, packageB.Release) == 1) {
            c[k] = packageA
        }
    }
    return c
}

func GetPackages(arch, branch string) Pkgs {
    branch_url, err := url.JoinPath(baseUrl, endpoint, branch)
    if err != nil {
        panic(fmt.Sprintf("\033[31m JoinPath error: %#v \033[0m", err))
    }

    req, err := http.NewRequest("GET", branch_url, nil)
    if err != nil {
        panic(fmt.Sprintf("\033[31m Request error: %#v \033[0m", err))
    }
    q := req.URL.Query()
    q.Add("arch", arch)
    req.URL.RawQuery = q.Encode()

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
         panic(fmt.Sprintf("\033[31m Response error: %#v \033[0m", err))
    }
    defer resp.Body.Close()

    var branchPackages BranchBinaryPackages
    err = json.NewDecoder(resp.Body).Decode(&branchPackages)
    if err != nil {
         panic(fmt.Sprintf("\033[31m Decode error: %#v \033[0m", err))
    }

    pkgs := make(Pkgs)

    for _, v := range branchPackages.Packages {
        pkgs[v.Name] = v
    }

    return pkgs
}
