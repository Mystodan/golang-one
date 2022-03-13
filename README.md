# Golang Assignment 1 - prog2006-2022 Advanced Programming
By Daniel Hao Huynh

## Golang 01: simple string processing
### String Specs

Given a string that is a sequence of numbers followed by dash followed by text, eg: `23-ab-48-caba-56-haha`
   * The numbers can be assumed to be unsigned integers
   * The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

Due to wanting one argument the format for running is:
`go run main.go` with

    * - test -> FLAG::BOOL
    * - avr -> FLAG::BOOL
    * - story -> FLAG::BOOL
    * - stats -> FLAG::BOOL
    * - gen -> FLAG::STRING

if no string is supplied all flags called will use the same randomly generated string
Usable flags are:
-test
-avr
-story
-stats
-gen
