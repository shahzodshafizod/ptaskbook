@echo off
set /p a="a = "
set /p b="b = "
set /p c="c = "
set /a miqdor=0

if %a% GTR 0 set /a miqdor+=1
if %b% GTR 0 set /a miqdor+=1
if %c% GTR 0 set /a miqdor+=1

echo miqdori adadhoi musbat: %miqdor%
echo on