@echo off
set /p A="A = "
set /p B="B (B>%A%) = "
set /a N=0
for /l %%i in (%A%, 1, %B%) do (
	echo %%i
	set /a N+=1
)
echo N = %N%
echo on