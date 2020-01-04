#!/bin/bash -e

### I would like to generate,

###     Option A
######################################
### 384 bit field with given modulus
#
# MODULUS=0x1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab
# go run . -output ./fp -bit 384 -opt A -modulus $MODULUS


###     Option B
#######################################
### 384 bit field with a random modulus
#
# go run . -output ./fp -bit 384 -opt B
#

###     Option C
#######################################
### 384 bit field with not-predefined modulus
#
go run . -output ./fp -bit 256 -opt C
#

###     Option D
#######################################
### Only x86 assembly operations all supported bit sizes
#
# go run . -output ./generic -opt D
#

#######################################
# format the code
goreturns -w -p ./fp
# run the test
go test ./fp -run Field -v -fuzz 50