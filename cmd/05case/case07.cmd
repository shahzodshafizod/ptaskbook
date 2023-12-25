@echo off
echo 1 - kilograms
echo 2 - milligrams
echo 3 - grams
echo 4 - tons
echo 5 - centners
set /p vazn="Raqami vohidi vaznro doxil kuned: "
set /p qimat="Qimati vaznro doxil kuned: "
set /a kilos=0
if %vazn%==1 (
	set /a kilos=%qimat%
) else if %vazn%==2 (
	set /a kilos=%qimat%/1000000
) else if %vazn%==3 (
	set /a kilos=%qimat%/1000
) else if %vazn%==4 (
	set /a kilos=%qimat%*1000
) else if %vazn%==5 (
	set /a kilos=%qimat%*100
)
echo kilograms: %kilos%
echo on