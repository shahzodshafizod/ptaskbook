@echo off
set /p A="A = "
set /p B="B = "
set /a jamm=A*A+B*B
set /a tarh=A*A-B*B
set /a zarb=A*A*B*B
set /a taqs=A*A/(B*B)
echo %A%*%A%+%B%*%B%=%jamm%
echo %A%*%A%-%B%*%B%=%tarh%
echo %A%*%A%*%B%*%B%=%zarb%
echo %A%*%A%/%B%*%B%=%taqs%
echo on