@echo off
set /p moh="adad [1;12] = "
if %moh%==1 (
	echo January
) else if %moh%==2 (
	echo February
) else if %moh%==3 (
	echo March
) else if %moh%==4 (
	echo April
) else if %moh%==5 (
	echo May
) else if %moh%==6 (
	echo June
) else if %moh%==7 (
	echo July
) else if %moh%==8 (
	echo August
) else if %moh%==9 (
	echo September
) else if %moh%==10 (
	echo October
) else if %moh%==11 (
	echo November
) else echo December
echo on