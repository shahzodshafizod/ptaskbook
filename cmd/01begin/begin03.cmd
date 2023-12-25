@echo OFF
set /p a="a = "
set /p b="b = "
set /a P=2*(%a%+%b%)
set /a S=%a%*%b%
echo S = %S%
echo P = %P%
echo ON