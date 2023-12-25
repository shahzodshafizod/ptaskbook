#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetRightLeavesCount(TNode*);

void Solve()
{
    Task("Tree8");
	TNode* P1;
	pt >> P1;
	int rights = GetRightLeavesCount(P1);
	pt << rights;
}

int GetRightLeavesCount(TNode* P)
{
	return P == NULL ? 0
		: int((P->Right != NULL) && (P->Right->Left == NULL) && (P->Right->Right == NULL))
		+ GetRightsCount(P->Left) + GetRightsCount(P->Right);
}