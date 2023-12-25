#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void SetParentToAllNodes(TNode* P);

void Solve()
{
    Task("Tree49");
	TNode* P1;
	pt >> P1;
	P1->Parent = NULL;
	SetParentToAllNodes(P1);
}

void SetParentToAllNodes(TNode* P)
{
	if (P->Left != NULL)
	{
		P->Left->Parent = P;
		SetParentToAllNodes(P->Left);
	}
	if (P->Right != NULL)
	{
		P->Right->Parent = P;
		SetParentToAllNodes(P->Right);
	}
}