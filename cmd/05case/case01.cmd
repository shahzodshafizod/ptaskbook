@echo off
set /p adad="adad [1;7]: "
if %adad%==1 (
	echo Monday
) else if %adad%==2 (
	echo Tuesday
) else if %adad%==3 (
	echo Wednesday
) else if %adad%==4 (
	echo Thursday
) else if %adad%==5 (
	echo Friday
) else if %adad%==6 (
	echo Saturday
) else echo Sunday
echo on