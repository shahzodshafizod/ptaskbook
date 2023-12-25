@echo off
set /p adad="adad: "
if %adad% GTR 0 (
	set /a adad+=1
) else if %adad% LSS 0 (
	set /a adad-=2
) else (
	set /a adad=10
)
echo adad: %adad%
echo on