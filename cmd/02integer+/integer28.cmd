@echo off
set /p K="K [1;365] = "
set /p N="N [1;7] = "
set /a hafta=(%K%+(%N%-2))%%7+1
echo ruzi hafta: %hafta%
echo on