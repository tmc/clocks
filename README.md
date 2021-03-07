# clocks

[![Project status](https://img.shields.io/github/release/tmc/clocks.svg?style=flat-square)](https://github.com/tmc/clocks/releases/latest)
[![Build Status](https://github.com/tmc/clocks/workflows/Test/badge.svg)](https://github.com/tmc/clocks/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/tmc/clocks?cache=0)](https://goreportcard.com/report/tmc/clocks)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/tmc/clocks)

Command clocks places additional time zones into your Mac status bar.

## Installation

clocks is a [Go](https://golang.org/) program for MacOS systems.

## Usage

```console
$ clocks -h
Usage of clocks:
  -fmt string
    	time format string (golang time pkg) (default "Mon Jan 2 15:04 MST")
  -tzs string
    	time zones (default "UTC,America/Los_Angeles")
```

## Example Output

This is example output from running `clocks`:

![example output](./docs/example-statusbar-image.png)
