#!/bin/bash

prompt() {
	read -p "$1"
}

cd mypkg

declare -a arr=(
"tree"
"cat mypkg.go"
"gomobile bind -i ."
"tree"
"unzip -l mypkg.aar"
)

for i in "${arr[@]}"
do
	prompt "\$ $i"
	$i
	echo
done

rm mypkg.aar
