#!/usr/bin/env bash

./goroutines-demo $1 & pid=$!
echo $pid
watch -n 0.5 "ps huH p $pid | wc -l >> threads${1}.txt"