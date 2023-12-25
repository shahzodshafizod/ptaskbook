@echo off
set /p A="A = "
set /p B="B (B>%A%) = "
set /a N=0
set /a avval=%B%-1
set /a oxir=%A%+1
for /l %%i in (%avval%, -1, %oxir%) do (
	echo %%i
	set /a N+=1
)
echo N = %N%
echo on