#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PostfixOutput(TNode*);
void Solve()
{
    Task("Tree14");
	TNode* P1;
	pt >> P1;
	PostfixOutput(P1);
}

void PostfixOutput(TNode* P)
{
	if (P == 0)
		return;
	PostfixOutput(P->Left);
	PostfixOutput(P->Right);
	pt << P->Data;
}