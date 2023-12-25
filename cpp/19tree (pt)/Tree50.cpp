#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* GetRoot(TNode*);

void Solve()
{
    Task("Tree50");
	TNode* P1;
	pt >> P1;
	TNode* P2 = GetRoot(P1);
	pt << P2;
}

TNode* GetRoot(TNode* P)
{
	if (P->Parent == NULL)
		return P;
	return GetRoot(P->Parent);
}