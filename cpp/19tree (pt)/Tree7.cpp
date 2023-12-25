#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetLeavesDataSum(TNode*);

void Solve()
{
    Task("Tree7");
	TNode* P1;
	pt >> P1;
	int leavesDataSum = GetLeavesDataSum(P1);
	pt << leavesDataSum;
}

int GetLeavesDataSum(TNode* P)
{
	return P == NULL ? 0
		: ( (P->Left == NULL) && (P->Right == NULL) ? P->Data : 0 )
		+ GetLeavesDataSum(P->Left) + GetLeavesDataSum(P->Right);
}