#ifndef ____FIELD____H____
#define ____FIELD____H____
#include "TNode.h"
#include "Flag.h"

struct Field {
	struct TNode*** nodes;
	enum Flag** chars;
	int cellWidth;
	int height;
	int width;
	char left;
	char center;
	char right;
};

void Field_constr(struct Field*, struct TNode*, int, int);
void Field_init(struct Field*, struct TNode*);
void Field_fillNodes(struct Field*, struct TNode*, int, int*);
void Field_fillFlags(struct Field*);
void Field_display(struct Field*);

struct TNode* Field_getNextNode(struct Field*, int, int);
void Field_printChars(char, int);
int Field_getSpaces(int);

#endif