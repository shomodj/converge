#!/usr/bin/env bash
set -eo pipefail

LENGTH=${LENGTH:-60}

# POPULATE CORPUS
test -d corpus || mkdir corpus
printf "a=b"  > corpus/kv
printf "="    > corpus/singleequals
printf "a==b" > corpus/doubleequals
printf "k=v=" > corpus/splitequals
printf "key5=😱" > corpus/unicode

# BUILD FUZZER
echo "-- building fuzzer --"
make cmd-fuzz.zip

echo "-- running fuzzer for $LENGTH seconds --"
go-fuzz -bin=./cmd-fuzz.zip -workdir=. &
PID=$!
sleep "$LENGTH"

echo "-- killing fuzz process --"
kill -2 $PID
sleep 0.25

echo 'fuzzing complete.'
