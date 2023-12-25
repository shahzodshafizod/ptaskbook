@echo off
set /p D="Day: "
set /p M="Month: "
set /a ruz=%D%
set /a moh=%M%
if %ruz%==1 (
	if %moh%==1 (
		set /a ruz=31
		set /a moh=12
	) else (
		if %moh%==2 (
			set /a ruz=31
		) else if %moh%==3 (
			set /a ruz=28
		) else if %moh%==4 (
			set /a ruz=31
		) else if %moh%==5 (
			set /a ruz=30
		) else if %moh%==6 (
			set /a ruz=31
		) else if %moh%==7 (
			set /a ruz=30
		) else if %moh%==8 (
			set /a ruz=31
		) else if %moh%==9 (
			set /a ruz=31
		) else if %moh%==10 (
			set /a ruz=30
		) else if %moh%==11 (
			set /a ruz=31
		) else if %moh%==12 (
			set /a ruz=30
		)
		set /a moh=%moh%-1
	)
) else (
	set /a ruz=%ruz%-1
)
echo D = %ruz%
echo M = %moh%
echo on