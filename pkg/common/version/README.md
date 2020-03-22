version for AVA Framework
======

version is a [Semantic Versioning](http://version.org/) library written in golang. It fully covers spec version `2.0.0`.

Usage
-----
```bash
$ go get github.com/ver13/ava/pkg/common/version
```

Features
-----

- Parsing and validation at all levels
- Comparator-like comparisons
- Compare Helper Methods
- InPlace manipulation
- Ranges `>=1.0.0 <2.0.0 || >=3.0.0 !3.0.1-beta.1`
- Wildcards `>=1.x`, `<=2.5.x`

Ranges
------

A `Range` is a set of conditions which specify which versions satisfy the range.

A condition is composed of an operator and a version. The supported operators are:

- `<1.0.0` Less than `1.0.0`
- `<=1.0.0` Less than or equal to `1.0.0`
- `>1.0.0` Greater than `1.0.0`
- `>=1.0.0` Greater than or equal to `1.0.0`
- `1.0.0`, `=1.0.0`, `==1.0.0` Equal to `1.0.0`
- `!1.0.0`, `!=1.0.0` Not equal to `1.0.0`. Excludes version `1.0.0`.

Note that spaces between the operator and the version will be gracefully tolerated.

A `Range` can link multiple `Ranges` separated by space:

Ranges can be linked by logical AND:

  - `>1.0.0 <2.0.0` would match between both ranges, so `1.1.1` and `1.8.7` but not `1.0.0` or `2.0.0`
  - `>1.0.0 <3.0.0 !2.0.3-beta.2` would match every version between `1.0.0` and `3.0.0` except `2.0.3-beta.2`

Ranges can also be linked by logical OR:

  - `<2.0.0 || >=3.0.0` would match `1.x.x` and `3.x.x` but not `2.x.x`

AND has a higher precedence than OR. It's not possible to use brackets.

Ranges can be combined by both AND and OR

  - `>1.0.0 <2.0.0 || >3.0.0 !4.2.1` would match `1.2.3`, `1.9.9`, `3.1.1`, but not `4.2.1`, `2.1.1`

Range usage:

```
v, err := SemanticVersion.Parse("1.2.3")
expectedRange, err := SemanticVersion.ParseRange(">1.0.0 <2.0.0 || >=3.0.0")
if expectedRange(v) {
    //valid
}

```

Example
-----

Have a look at full examples in [examples/version/main.go](examples/main.go)

```go
import github.com/ver13/ava/pkg/common/version

v, err := version.Make("0.0.1-alpha.preview+123.github")
fmt.Printf("Major: %d\n", v.Major)
fmt.Printf("Minor: %d\n", v.Minor)
fmt.Printf("Patch: %d\n", v.Patch)
fmt.Printf("Pre: %s\n", v.Pre)
fmt.Printf("Build: %s\n", v.Build)

// Prerelease versions array
if len(v.Pre) > 0 {
    fmt.Println("Prerelease versions:")
    for i, pre := range v.Pre {
        fmt.Printf("%d: %q\n", i, pre)
    }
}

// Build meta data array
if len(v.Build) > 0 {
    fmt.Println("Build meta data:")
    for i, build := range v.Build {
        fmt.Printf("%d: %q\n", i, build)
    }
}

v001, err := version.Make("0.0.1")
// Compare using helpers: v.GT(v2), v.LT, v.GTE, v.LTE
v001.GT(v) == true
v.LT(v001) == true
v.GTE(v) == true
v.LTE(v) == true

// Or use v.Compare(v2) for comparisons (-1, 0, 1):
v001.Compare(v) == 1
v.Compare(v001) == -1
v.Compare(v) == 0

// Manipulate Version in place:
v.Pre[0], err = version.NewPRVersion("beta")
if err != nil {
    fmt.Printf("Error parsing pre release version: %q", err)
}

fmt.Println("\nValidate versions:")
v.Build[0] = "?"

err = v.Validate()
if err != nil {
    fmt.Printf("Validation failed: %s\n", err)
}
```

Build
-----

```
go build -X "github.com/ver13/ava/pkg/common/version.name=Golang Microservices Framework" \
         -X "github.com/ver13/ava/pkg/common/version.ServerName=gmfServer" \
         -X "github.com/ver13/ava/pkg/common/version.ClientName=gmfCli" \
         -X "github.com/ver13/ava/pkg/common/version.version=v1.0.0" \
         -X "github.com/ver13/ava/pkg/common/version.Commit=f0f7b7dab7e36c20b757cebce0e8f4fc5b95de60" \
         -X "github.com/ver13/ava/pkg/common/version.BuildTags=linux darwin amd64"
```
