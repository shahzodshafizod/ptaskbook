#include "TNode.h"

#ifndef NULL
#define NULL 0
#endif

void
TNode_constr(struct TNode* this) {
	this->Data = 0;
	this->Left = NULL;
	this->Right = NULL;
	this->Parent = NULL;
}