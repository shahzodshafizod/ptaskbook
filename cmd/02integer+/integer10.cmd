@echo OFF
set /p seraqama="adadi seraqama: "
set /a vohid=%seraqama%%%10
set /a dahi=%seraqama%/10%%10
echo vohid: %vohid%
echo dahi: %dahi%
echo ON