@echo OFF
set /p seraqama="adadi seraqama: "
set /a avval=%seraqama%/10
set /a vohid=%seraqama%%%10
set /a seraqama=%vohid%*100+%avval%
echo natija: %seraqama%
echo ON