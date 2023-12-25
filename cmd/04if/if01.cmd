@echo off
set /p adad="adad: "
if %adad% GTR 0 (
	set /a adad+=1
)
echo adad: %adad%
echo on