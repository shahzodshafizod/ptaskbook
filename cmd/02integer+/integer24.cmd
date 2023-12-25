@echo OFF
set /p K="K [1;365] = "
set /a hafta=%K%%%7
echo raqami ruzi hafta: %hafta%
echo ON