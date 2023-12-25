#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PrefixOutput(TNode*);
void Solve()
{
    Task("Tree13");
	TNode* P1;
	pt >> P1;
	PrefixOutput(P1);
}

void PrefixOutput(TNode* P)
{
	if (P == NULL)
		return;
	pt << P->Data;
	PrefixOutput(P->Left);
	PrefixOutput(P->Right);
}