#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void AddNodesToLeaves(TNode*);

void Solve()
{
    Task("Tree44");
	TNode* P1;
	pt >> P1;
	AddNodesToLeaves(P1);
}

void AddNodesToLeaves(TNode* P)
{
	if (P == NULL)
		return;
	if ( (P->Left == NULL) && (P->Right == NULL) )
	{
		P->Left = new TNode;
		P->Left->Left = P->Left->Right = NULL;
		P->Left->Data = 10;

		P->Right = new TNode;
		P->Right->Left = P->Right->Right = NULL;
		P->Right->Data = 11;
	}
	else
	{
		AddNodesToLeaves(P->Left);
		AddNodesToLeaves(P->Right);
	}
}