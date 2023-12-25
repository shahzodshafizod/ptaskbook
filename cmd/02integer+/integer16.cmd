@echo OFF
set /p seraqama="adadi seraqama: "
set /a sadi=%seraqama%/100
set /a dahi=%seraqama%/10%%10
set /a vohid=%seraqama%%%10
set /a seraqama=%sadi%*100+%vohid%*10+%dahi%
echo natija: %seraqama%
echo ON