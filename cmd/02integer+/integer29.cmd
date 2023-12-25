@echo off
set /p A="A = "
set /p B="B = "
set /p C="C = "
set /a kvadratho=(%A%/%C%)*(%B%/%C%)
set /a xoligi=(%A%*%B%)-(%C%*%C%*%kvadratho%)
echo miqdori kvadratho: %kvadratho%
echo masohati bandnashuda: %xoligi%
echo on