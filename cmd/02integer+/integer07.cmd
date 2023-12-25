@echo OFF
set /p duraqama="adadi duraqama: "
set /a dahi=%duraqama%/10
set /a vohid=%duraqama%%%10
set /a jamm=%dahi%+%vohid%
set /a zarb=%dahi%*%vohid%
echo %dahi% + %vohid% = %jamm%
echo %dahi% * %vohid% = %zarb%
echo ON