@echo OFF
set /p seraqama="adadi seraqama: "
set /a sadi=%seraqama%/100
set /a boqimonda=%seraqama%%%100
set /a seraqama=%boqimonda%*10+%sadi%
echo natija: %seraqama%
echo ON