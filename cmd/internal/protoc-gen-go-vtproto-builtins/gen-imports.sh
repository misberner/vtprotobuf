#!/bin/sh

(
  echo 'package main'
  go list -f 'import _ "{{.ImportPath}}"' google.golang.org/protobuf/types/...
) | gofmt -s >imports.go
