#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* SearchTree(TNode* P, TNode** prev);

void Solve()
{
    Task("Tree57");
	TNode* P1, *prev = NULL;
	pt >> P1;
	TNode* P2 = SearchTree(P1, &prev);
	pt << P2;
}

TNode* SearchTree(TNode* P, TNode** prev)
{
	if (P == NULL)
		return P;
	TNode* node = SearchTree(P->Left, prev);
	if (node != NULL)
		return node;
	
	if (*prev == NULL)
		*prev = P;
	else if ((*prev)->Data > P->Data)
		return P;
	
	*prev = P;
	node = SearchTree(P->Right, prev);
	
	return node;
}