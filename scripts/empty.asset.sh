#!/bin/sh

filepath=$1
echo "" > $filepath/static.go
echo "package web" >> $filepath/static.go
echo "import \"fmt\"" >> $filepath/static.go
echo "func Asset(name string) ([]byte, error) { return nil, fmt.Errorf(\"Asset %s not found\", name) }" >> $filepath/static.go
echo "func AssetNames() []string              { return []string{} }" >> $filepath/static.go
