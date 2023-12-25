@echo off

gcc -std=c11 -c %1.c
gcc %1.o -o %1.exe

@echo on