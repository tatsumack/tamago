#! /bin/bash

try() {
    expected="$1"
    input="$2"
    echo $input | ./tamago > tmp.ll
    lli tmp.ll
    got="$?"

    if [ "$got" != "$expected" ]; then
        echo "$expected expected, but got $got"
        exit 1
    fi
}

try 0 0
try 42 42
try 255 255

echo OK
