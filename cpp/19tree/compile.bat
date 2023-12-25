g++ -c TNode.cpp
g++ -c Field.cpp
g++ -c Tree.cpp
g++ -c %1.cpp
g++ TNode.o Field.o Tree.o %1.o -o %1.exe