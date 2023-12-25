@echo off

gcc -std=c11 -c TNode.c
gcc -std=c11 -c Field.c
gcc -std=c11 -c Tree.c
gcc -std=c11 -c %1.c
gcc TNode.o Field.o Tree.o %1.o -o %1.exe

@echo on