# available arch

|**arch / branch**| p10 | sisyphus |
|-----------------|-----|----------|
| aarch64         | +   | +        |
| armh            | +   |          |
| i586            | +   | +        |
| noarch          | +   | +        |
| ppc64le         | +   | +        |
| srpm            | +   | +        |
| x86_64          | +   | +        |

# build
```
go mod tidy
go build ./cmd/cli/cli.go
```

# usage
```./cli --arch=<arch> --A=<branch A> --B=<branch B> --compare=<mode>```
- arch — system architecture;
- A, B — target branches;
- compare — compare mode: existence or version.

# examples of usage:
- `./cli -arch aarch64 -B sisyphus -A p10 --compare=existence > test1.json`
- `./cli -arch i586 -A sisyphus -B p10 --compare=existence > test2.json`
- `./cli -A sisyphus -B p10 --compare=version > test3.json`
