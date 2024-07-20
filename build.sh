#!/bin/bash

# Array of operating systems and architectures
OS_ARCHES=("windows/amd64" "linux/amd64" "darwin/amd64" "linux/arm" "linux/arm64")

# Iterate over each OS/ARCH combination and build the application
for os_arch in "${OS_ARCHES[@]}"
do
  IFS="/" read -r -a arr <<< "$os_arch"
  os=${arr[0]}
  arch=${arr[1]}
  output_name="togo-$os-$arch"
  if [ "$os" = "windows" ]; then
    output_name+=".exe"
  fi

  echo "Building for $os/$arch..."
  GOOS=$os GOARCH=$arch go build -o $output_name main.go
done

echo "Builds completed."
