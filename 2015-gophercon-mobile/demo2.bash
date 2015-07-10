#!/bin/bash

prompt() {
	read -p "$1"
}

export ANDROID_HOME=/Users/hakim/Library/Android/sdk

declare -a arr=(
"cd mypkg"
"tree"
"cat mypkg.go"
"export ANDROID_HOME=/Users/hakim/Library/Android/sdk"
"gomobile bind ."
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
