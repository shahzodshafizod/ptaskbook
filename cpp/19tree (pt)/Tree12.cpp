#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void InfixOutput(TNode*);
void Solve()
{
    Task("Tree12");
	TNode* P1;
	pt >> P1;
	InfixOutput(P1);
}

void InfixOutput(TNode* P)
{
	if (P == 0)
		return;
	InfixOutput(P->Left);
	pt << P->Data;
	InfixOutput(P->Right);
}