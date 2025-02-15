.PHONY: regex flex test smc
CPP_SRC = ./regex/regex.cpp
FLEX_SRC = ./flex/scan.l
TESTS = ./input.txt
CC = gcc

smc:
	cd smc/my && make clean && make all
	smc/my/checkstring $(TESTS)