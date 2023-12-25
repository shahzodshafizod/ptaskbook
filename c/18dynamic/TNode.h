#ifndef ____TNODE____H____
#define ____TNODE____H____

struct TNode {
	int Data;
	struct TNode* Next;
	struct TNode* Prev;
};

void TNode_constr(struct TNode*);

#endif