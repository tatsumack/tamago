#! /bin/bash

try () {
    expected="$1"
    input="$2"

    echo test "$input"

    echo "$input" | ./bin/tamago > ./bin/tmp.ll
    lli ./bin/tmp.ll
    got="$?"

    if [ "$got" != "$expected" ]; then
        echo -e "\033[31mNG $expected expected, but got $got\033[m"
        return
    fi
    echo -e '\033[32mOK\033[m'
}

try 0 0
try 42 42
try 255 255

try 2 '1 + 1'
try 11 '1 + 10'
try 0 '1 - 1'
try 9 '10 - 1'
try 9 '3 * 3'
try 12 '3 * 4'
try 3 '9 / 3'
try 2 '5 / 2'

try 23 '3 + 4 * 5'
try 80 '100 - 4 * 5'
try 98 '100 - 10 / 5'
try 23 '3 * 2 + 4 * 5 - 9 / 3'

