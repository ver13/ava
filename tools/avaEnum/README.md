# go-enum

[![CircleCI](https://circleci.com/gh/ver13/ava-enum.svg?style=svg&circle-token=b44c10ce16bcef76e86da801d67811a5ff71fc72)](https://circleci.com/gh/ver13/ava-enum)
[![Go Report Card](https://goreportcard.com/badge/github.com/ver13/ava-enum)](https://goreportcard.com/report/github.com/ver13/ava-enum)
[![Coverage Status](https://coveralls.io/repos/github/ver13/ava-enum/badge.svg)](https://coveralls.io/github/ver13/ava-enum)
[![GoDoc](https://godoc.org/github.com/ver13/ava-enum?status.svg)](https://godoc.org/github.com/ver13/ava-enum)

An enum generator for go

## How it works

The goal of ava-enum is to create an easy to use enum generator that will take a decorated type declaration like `type EnumName int` and create the associated constant values and funcs that will make life a little easier for adding new values.
It's not perfect, but I think it's useful.

I took the output of the [Stringer](https://godoc.org/golang.org/x/tools/cmd/stringer) command as the `String()` method, and added a way to parse a string value.

## Command options

``` shell
ava-enum --help
Options:

  -h, --help       display help information
  -f, --file      *The file(s) to generate enums.  Use more than one flag for more files.
      --noprefix   Prevents the constants generated from having the Enum as a prefix.
      --lower      Adds lowercase variants of the enum strings for lookup.
      --marshal    Adds text (and inherently json) marshalling functions.
      --sql        Adds SQL database scan and value functions.
      --flag       Adds golang flag functions.
      --prefix     Replaces the prefix with a user one.
      --names      Generates a 'Names() []string' function, and adds the possible enum values in the error response during parsing
      --nocamel    Removes the snake_case to CamelCase name changing
```

### Syntax

The parser looks for comments on your type defs and parse the enum declarations from it.
The parser will look for `ENUM(` and continue to look for comma separated values until it finds a `)`.  You can put values on the same line, or on multiple lines.\
If you need to have a specific value jump in the enum, you can now specify that by adding `=numericValue` to the enum declaration.  Keep in mind, this resets the data for all following values.  So if you specify `50` in the middle of an enum, each value after that will be `51, 52, 53...`

#### Comments

You can use comments inside enum that start with `//`\
The comment must be at the end of the same line as the comment value, only then it will be added as a comment to the generated constant.

```go
// Commented is an enumeration of commented values
/*
ENUM(
value1 // Commented value 1
value2
value3 // Commented value 3
)
*/
type Commented int
```

The generated comments in code will look something like:

```go
...
const (
    // CommentedValue1 is a Commented of type Value1
    // Commented value 1
    CommentedValue1 Commented = iota
    // CommentedValue2 is a Commented of type Value2
    CommentedValue2
    // CommentedValue3 is a Commented of type Value3
    // Commented value 3
    CommentedValue3
)
...
```

## Adding it to your project

1. `go get github.com/ver13/ava/tools/ava-enum`
1. Add a go:generate line to your file like so... `//go:generate ava-enum -f=$GOFILE --marshal`
1. Run go generate like so `go generate ./...`
1. Enjoy your newly created Enumeration
