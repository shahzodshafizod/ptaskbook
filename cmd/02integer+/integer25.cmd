@echo off
set /p K="K [1;365] = "
set /a hafta=(%K%+3)%%7
echo ruzi hafta: %hafta%
echo on