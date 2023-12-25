#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void LeftCounter(TNode*, int&);
int GetLeftCount(TNode*);

void Solve()
{
    Task("Tree5");
	TNode* P1;
	pt >> P1;
	int counts = GetLeftCount(P1);
	pt << counts;
}

int GetLeftCount(TNode* P)
{
	return P == NULL ? 0 : int(P->Left != NULL) + GetLeftCount(P->Left) + GetLeftCount(P->Right);
}

void LeftCounter(TNode* pointer, int& counts)
{
	if (pointer == NULL)
		return;
	counts += int(pointer->Left != NULL);
	LeftCounter(pointer->Left, counts);
	LeftCounter(pointer->Right, counts);
}