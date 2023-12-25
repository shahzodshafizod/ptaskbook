@echo off
set /p N="N [1;4] = "
set /p A="A = "
set /p B="B (B is not 0) = "
if %N%==1 (
	set /a result=%A%+%B%
	set amal=+
) else if %N%==2 (
	set /a result=%A%-%B%
	set amal=-
) else if %N%==3 (
	set /a result=%A%*%B%
	set amal=*
) else if %N%==4 (
	set /a result=%A%/%B%
	set amal=/
)
echo %A% %amal% %B% = %result%
echo on