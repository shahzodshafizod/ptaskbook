#ifndef ____FIELD____H____
#define ____FIELD____H____
#include "TNode.h"
#include <cstdlib>

enum TChar {SPACE, LEFT, CENTER, RIGHT, NUMBER};

class Field {
	private: const int cellWidth;
	private: int height;
	private: int width;
	private: TNode*** matrix;
	private: TChar** chars;
	
	public: Field(TNode*, int, int);
	private: void init(TNode*);
	public: ~Field();
	private: void fillMatrix(TNode*, int, int&);
	private: void fillChars();
	private: void printChars(char, int);
	private: TNode* getNextNode(int, int);
	private: int getSpaces(int);
	public: void display(TNode* node = NULL);
};

#endif