#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

pkgArgs=${*:-}

pkgs=${pkgArgs:-$(find cmd/* -type d)}

for dir in $pkgs
do
	filename=$(echo -n "$dir" | cut -d "/" -f 2)
	echo "Building directory $dir to $filename:"
	echo -n "	go build -o \"$filename $dir/\"*.go ... "
	GOARCH=amd64 go build -o "$filename" -ldflags="-s -w" "$dir/"*.go
	GOARCH=arm GOARM=6 go build -o "$filename-arm" -ldflags="-s -w" "$dir/"*.go
	echo "done!"
done

echo "Everything built!"
