rem:>"Tree book.txt"

echo TNode.h>>"Tree book.txt"
type TNode.h>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

echo TNode.cpp>>"Tree book.txt"
type TNode.cpp>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

echo Field.h>>"Tree book.txt"
type Field.h>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

echo Field.cpp>>"Tree book.txt"
type Field.cpp>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

echo Tree.h>>"Tree book.txt"
type Tree.h>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

echo Tree.cpp>>"Tree book.txt"
type Tree.cpp>>"Tree book.txt"
echo.>>"Tree book.txt"
echo.>>"Tree book.txt"

for /l %%i in (1, 1, 30) do (
	
	if %%i lss 10 (
		echo Tree %%i>>"Tree book.txt"
		type tree00%%i.cpp>>"Tree book.txt"
	) else (
		echo Tree %%i>>"Tree book.txt"
		type tree0%%i.cpp>>"Tree book.txt"
	)
	echo.>>"Tree book.txt"
	echo.>>"Tree book.txt"
)