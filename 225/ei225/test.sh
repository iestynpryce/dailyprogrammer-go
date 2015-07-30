#!/bin/bash

PROGRAM=ei225.go
TEMPDIR=$(mktemp -dt "$0.XXXXXXX")

errcount=0
testcount=0

for i in {1..3}; do
	outfile=${TEMPDIR}/out${i}.txt
	go run ${PROGRAM} > $outfile
	if diff -q solution${i}.txt $outfile > /dev/null; then
		echo "[ERROR] output for test $i incorrect" >&2
		let errcount+=1
	fi
	let testcount+=1
done

passed=$(($testcount - $errcount))
echo "$testcount tests run: PASSED $passed FAILED $errcount"
