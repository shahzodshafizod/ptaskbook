@echo OFF
set /p seraqama="adadi seraqama: "
set /a sadi=%seraqama%/100
set /a dahi=%seraqama%/10%%10
set /a vohid=%seraqama%%%10
set /a jamm=%sadi%+%dahi%+%vohid%
set /a zarb=%sadi%*%dahi%*%vohid%
echo %sadi% + %dahi% + %vohid% = %jamm%
echo %sadi% * %dahi% * %vohid% = %zarb%
echo ON