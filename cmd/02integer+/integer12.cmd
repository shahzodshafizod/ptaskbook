@echo OFF
set /p seraqama="adadi seraqama: "
set /a sadi=%seraqama%/100
set /a dahi=%seraqama%/10%%10
set /a vohid=%seraqama%%%10
set /a seraqama=%vohid%*100+%dahi%*10+%sadi%
echo natija: %seraqama%
echo ON