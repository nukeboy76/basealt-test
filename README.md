TODO:

1. *touch* api, look on the raw list of bin packages
2. think about the most suitable json struct for '2)'
# coding
3. make go module and implement json structures
4. and cli with `compare` mode selection:
- if pkg in A and pkg not in B -> [pkg, ]
- if A[pkg].version > B[pkg].version -> [pkg, ]

usage:
```
$ ./cli --arch=<arch> --a=<branch A> --b=<branch B> --compare=<existance|version>
	arch — system architecture;
	a, b — target branch;
	compare — compare mode: existance or version;
```
