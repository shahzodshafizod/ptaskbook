@echo off
set /p K="K = "
set /p N="N (N>0) = "
for /L %%i in (1, 1, %N%) do echo %K%
echo on