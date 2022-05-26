# Commander

[![Go package](https://github.com/efureev/go-commander/actions/workflows/go.yml/badge.svg)](https://github.com/efureev/go-commander/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/efureev/go-commander)](https://goreportcard.com/report/github.com/efureev/go-commander)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/efureev/go-commander)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/efureev/go-commander)
![GitHub](https://img.shields.io/github/license/efureev/go-commander)

## Installation

```shell
go get github.com/efureev/go-commander
```

## Usage

```go
package main

import (
	commander "github.com/efureev/go-commander"
)

func main() {
	cmdr := commander.NewCommander().
		Add(
			commander.NewCommand(`info`).
				OnRun(func(cmd *commander.Command) error {
					println(cmd.Name)
					return nil
				}),
			cmdFinish(`finish`),
			// ... 
		)

	cmdr.Run()
}

func cmdFinish(t string) *commander.Command {
	return commander.NewCommand(`Finish!`).
		OnPrepare(func(cmd *commander.Command) error {
			println(`OnPrepare`)
			return nil
		}).
		OnRun(func(cmd *commander.Command) error {
			println(t)
			return nil
		}).
		OnDone(func(cmd *commander.Command) error {
			println(`OnDone`)
			return nil
		})
}
```
