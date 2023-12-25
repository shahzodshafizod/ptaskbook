#ifndef ____TREE____H____
#define ____TREE____H____
#include "TNode.h"
#include "Field.h"

struct Tree {
	struct TNode* root;
	struct TNode* current;
	struct Field* field;
	int nodeCount;
	int level;
};

void Tree_constr(struct Tree*);
void Tree_destr(struct Tree*);
void Tree_automake(struct Tree*, int);
void Tree_free(struct Tree*);
void Tree_infix(struct Tree*);
void Tree_display(struct Tree*);
int Tree_getNodeCount(struct Tree*, enum Flag, int);
int Tree_getLevel(struct Tree*);
void Tree_setLevel(struct Tree*, int);
int Tree_getSum(struct Tree*, enum Flag);
void Tree_toArray(struct Tree*, int[], int, enum Flag);
void Tree_print(struct Tree*, enum Flag);

void Tree_init(struct Tree*);
int Tree_getRandomData(int, int);
int Tree_getRandomDirection();

#endif