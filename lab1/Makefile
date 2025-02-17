.PHONY: regex flex test smc
CPP_SRC = ./regex/regex.cpp
FLEX_SRC = ./flex/scan.l
TESTS = ./input.txt
CC = gcc
CXXFLAGS = -std=c++11


smc:
	cd smc/my && make clean && make all
	smc/my/checkstring $(TESTS)

regex:
	python3 regex/regex.py

