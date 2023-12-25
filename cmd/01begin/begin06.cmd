@echo off
set /p a="a = "
set /p b="b = "
set /p c="c = "
set /a V=%a%*%b%*%c%
set /a S=2*(%a%*%b% + %b%*%c% + %c%*%a%)
echo V = %V%
echo S = %S%
echo on