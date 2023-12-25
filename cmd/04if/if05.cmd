@echo off
set /p a="a = "
set /p b="b = "
set /p c="c = "
set /a musbatho=0
set /a manfiho=0

if %a% GTR 0 set /a musbatho+=1
if %b% GTR 0 set /a musbatho+=1
if %c% GTR 0 set /a musbatho+=1

if %a% LSS 0 set /a manfiho+=1
if %b% LSS 0 set /a manfiho+=1
if %c% LSS 0 set /a manfiho+=1

echo miqdori adadhoi musbat: %musbatho%
echo miqdori adadhoi manfi: %manfiho%
echo on