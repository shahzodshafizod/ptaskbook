@echo off
set /p D="D = "
set /p M="M = "
set /a ruz=%D%
set /a moh=%M%

if %ruz%LEQ31 (
	if %ruz%GEQ31 (
		if %moh%==1 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==3 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==5 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==7 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==8 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==10 (
			set /a moh=%moh%+1
			set /a ruz=1
		) else if %moh%==12 (
			set /a moh=%moh%+1
			set /a ruz=1
		) 
	)
)

echo D = %ruz%
echo M = %moh%
echo on