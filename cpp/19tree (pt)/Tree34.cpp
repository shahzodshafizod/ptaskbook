#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void CopyTree(TNode* P1, TNode** P2, bool right = false);

void Solve()
{
    Task("Tree34");
	TNode* P1;
	pt >> P1;
	TNode* P2 = NULL;
	CopyTree(P1, &P2);
	pt << P2;
}

void CopyTree(TNode* P1, TNode** P2, bool right)
{
	if (P1 == NULL)
		return;
	TNode* node = new TNode;
	node->Data = P1->Data;
	node->Left = node->Right = NULL;
	if (*P2 == NULL)
		*P2 = node;
	else if (right)
		(*P2)->Right = node;
	else
		(*P2)->Left = node;
	CopyTree(P1->Left, &node);
	CopyTree(P1->Right, &node, true);
}