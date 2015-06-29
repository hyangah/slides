#!/bin/bash

prompt() {
	read -p "$1"
}

cd hello

declare -a arr=(
#"less main.go"
"go run main.go"
#"gomobile install -i ."
)

for i in "${arr[@]}"
do
	prompt "\$ $i"
	$i
	echo
done
