#ifndef ____TNODE____H____
#define ____TNODE____H____

struct TNode {
	int Data;
	struct TNode* Left;
	struct TNode* Right;
	struct TNode* Parent;
};

void TNode_constr(struct TNode*);

#endif