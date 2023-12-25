@echo off

gcc -std=c11 -c TNode.c
gcc -std=c11 -c TStack.c
gcc -std=c11 -c %1.c
gcc TNode.o TStack.o %1.o -o %1.exe

@echo on