@echo off
set /p A="A = "
set /p B="B = "
if %A% NEQ %B% (
	if %A% GTR %B% (
		set /a B=%A%
	) else (
		set /a A=%B%
	)
) else (
	set /a A=0
	set /a B=0
)
echo A = %A%
echo B = %B%
echo on