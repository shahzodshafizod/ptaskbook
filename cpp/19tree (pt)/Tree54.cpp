#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* GetRoot(TNode* P);
void Copy(TNode* P1, TNode** P2);

void Solve()
{
    Task("Tree54");
	TNode* P1;
	pt >> P1;
	P1 = GetRoot(P1);
	TNode* P2 = NULL;
	Copy(P1, &P2);
	pt << P2;
}

TNode* GetRoot(TNode* P)
{
	if ( (P == NULL) || (P->Parent == NULL) )
		return P;
	return GetRoot(P->Parent);
}

void Copy(TNode* P1, TNode** P2)
{
	if (*P2 == NULL)
	{
		*P2 = new TNode;
		(*P2)->Left = (*P2)->Right = NULL;
		(*P2)->Data = P1->Data;
		(*P2)->Parent = NULL;
	}
	if (P1 == NULL)
		return;
	TNode* left = P1->Left;
	if (left != NULL)
	{
		(*P2)->Left = new TNode;
		(*P2)->Left->Left = (*P2)->Left->Right = NULL;
		(*P2)->Left->Data = left->Data;
		(*P2)->Left->Parent = *P2;
		Copy(left, &(*P2)->Left);
	}
	TNode* right = P1->Right;
	if (right != NULL)
	{
		(*P2)->Right = new TNode;
		(*P2)->Right->Left = (*P2)->Right->Right = NULL;
		(*P2)->Right->Data = right->Data;
		(*P2)->Right->Parent = *P2;
		Copy(right, &(*P2)->Right);
	}
}