package basealt

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
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
    Length      uint32      `json:"length"`
    Packages    []Package   `json:"packages"`
}

type Pkgs map[string]Package

func GetBranchBinaryPackages(arch string, branch string) Pkgs {
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
