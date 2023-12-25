@echo OFF
set /p duraqama="adadi duraqama: "
set /a dahi=%duraqama%/10
set /a vohid=%duraqama%%%10
set /a duraqama=%vohid%*10+%dahi%
echo natija = %duraqama%
echo ON