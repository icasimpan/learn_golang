#!/bin/bash

go_src=$1

## block no parameters passed
[[ $# -eq 0 ]] && {
  echo "ERROR: unspecified go source filename to create
  "
  exit 1
}

## make sure $go_src does not exit yet
## to avoid erasing prior works
[[ $# -gt 0 ]] && [[ -e $go_src ]] && {
  echo "ERROR: $go_src exist. Aborting.
  "
  exit 1
} 

cat > $1 <<EOL
package main

import "fmt"

func main(){
  fmt.Println("start you code here")
}
EOL
