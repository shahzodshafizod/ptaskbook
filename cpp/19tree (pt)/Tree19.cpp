#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetMaximal(TNode*);
void MaximalsCounter(TNode* P, int maximal, int& counts);

void Solve()
{
    Task("Tree19");
	TNode* P1;
	pt >> P1;
	int maximal = GetMaximal(P1);
	int counts = 0;
	MaximalsCounter(P1, maximal, counts);
	pt << maximal << counts;
}

int GetMaximal(TNode* P)
{
	int maximal = P->Data;
	if (P->Left != NULL)
	{
		int left = GetMaximal(P->Left);
		if (left > maximal)
			maximal = left;
	}
	if (P->Right != NULL)
	{
		int right = GetMaximal(P->Right);
		if (right > maximal)
			maximal = right;
	}
	return maximal;
}

void MaximalsCounter(TNode* P, int maximal, int& counts)
{
	if (P == NULL)
		return;
	counts += int(P->Data == maximal);
	MaximalsCounter(P->Left, maximal, counts);
	MaximalsCounter(P->Right, maximal, counts);
}