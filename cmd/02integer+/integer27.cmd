@echo off
set /p K="K [1;365] = "
set /a hafta=(%K%+4)%%7+1
echo ruzi hafta: %hafta%
echo on