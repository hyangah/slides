#!/bin/bash

prompt() {
	read -p "$1"
}

cd hello

declare -a arr=(
#"less main.go"
"wc -l main.go"
"go run main.go"
"gomobile install ."
)

for i in "${arr[@]}"
do
	prompt "\$ $i"
	$i
	echo
done
