#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetLeafCount(TNode*);

void Solve()
{
    Task("Tree6");
	TNode* P1;
	pt >> P1;
	pt << GetLeafCount(P1);
}

int GetLeafCount(TNode* P)
{
	return P == NULL ? 0 : int((P->Left == NULL) && (P->Right == NULL)) + GetLeafCount(P->Left) + GetLeafCount(P->Right);
}