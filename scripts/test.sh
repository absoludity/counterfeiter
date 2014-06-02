#!/usr/bin/env bash

set -e

counterfeiter='go run main.go'

$counterfeiter fixtures Something
$counterfeiter fixtures HasVarArgs
$counterfeiter fixtures HasImports
$counterfeiter fixtures/another_package InAliasedPackage

go test -v .
