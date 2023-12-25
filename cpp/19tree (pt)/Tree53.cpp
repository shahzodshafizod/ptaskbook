#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* GetClosestParent(TNode* P1, TNode* P2);

void Solve()
{
    Task("Tree53");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P0 = GetClosestParent(P1, P2);
	pt << P0;
}

TNode* GetClosestParent(TNode* P1, TNode* P2)
{
	if (P1 == NULL)
		return NULL;
	for (TNode* i = P2; i != NULL; i = i->Parent)
		if (i == P1)
			return i;
	return GetClosestParent(P1->Parent, P2);
}