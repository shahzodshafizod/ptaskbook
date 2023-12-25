@echo off
set /p a="adadi #1: "
set /p b="adadi #2: "
set /p c="adadi #3: "
if %a% == %b% (
	echo 3
) else (
	if %c% == %a% (
		echo 2
	) else (
		echo 1
	)
)
echo on