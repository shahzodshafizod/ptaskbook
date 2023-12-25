#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetMinimal(TNode* P);
void LeavesWithMinimalCounter(TNode* P, int minimal, int& counts);

void Solve()
{
    Task("Tree20");
	TNode* P1;
	pt >> P1;
	int minimal = GetMinimal(P1);
	int counts = 0;
	LeavesWithMinimalCounter(P1, minimal, counts);
	pt << minimal << counts;
}

int GetMinimal(TNode* P)
{
	int minimal = P->Data;
	if (P->Left != NULL)
	{
		int left = GetMinimal(P->Left);
		if (left < minimal)
			minimal = left;
	}
	if (P->Right != NULL)
	{
		int right = GetMinimal(P->Right);
		if (right < minimal)
			minimal = right;
	}
	return minimal;
}

void LeavesWithMinimalCounter(TNode* P, int minimal, int& counts)
{
	if (P == NULL)
		return;
	if (P->Data == minimal)
	{
		if ( (P->Left == NULL) && (P->Right == NULL) )
			counts++;
	}
	LeavesWithMinimalCounter(P->Left, minimal, counts);
	LeavesWithMinimalCounter(P->Right, minimal, counts);
}