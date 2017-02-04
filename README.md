gofmtrlx: Relaxed gofmt
=======================

This is a fork of [gofmt](https://golang.org/cmd/gofmt/).

## What is a problem?

gofmt can't format a code which contains syntax error. For example,

```go
package main

func main() {
	s := []int{
		1,
		2,
		3
	}
}
```

This code can't be formatted because there is no trailing comma at line 6.

## How did gofmtrlx solve the problem?

Actually, Go's parser can parse this without bad AST node. So gofmt can format the code with fixing syntax error.

gofmtrlx ignores syntax error if there is no bad node in AST. In above code, gofmtrlx can fill up trailing comma
on formatting it.

## Installation and usage

```
go get -u github.com/rhysd/gofmtrlx
```

Usage is the same as `gofmt` command.

## License

All codes are licensed the same license as original gofmt.
