@echo off
set /p adad="adad: "
if %adad% GTR 0 (
	set /a adad+=1
) else (
	set /a adad-=2
)
echo adad: %adad%
echo on