@echo off
echo 1-decimeters
echo 2-kilometers
echo 3-meters
echo 4-millimeters
echo 5-santimeters
set /p darozi="Raqami vohidi daroziro doxil kuned: "
set /p qimat="Qimati daroziro doxil kuned: "
if %darozi%==1 (
	set /a meters=%qimat%/10
) else if %darozi%==2 (
	set /a meters=%qimat%*1000
) else if %darozi%==3 (
	set /a meters=%qimat%
) else if %darozi%==4 (
	set /a meters=%qimat%/100000
) else if %darozi%==5 (
	set /a meters=%qimat%/100
)
echo meters = %meters%
echo on