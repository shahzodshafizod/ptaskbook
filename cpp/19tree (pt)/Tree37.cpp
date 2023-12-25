#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Recursion(TNode* P);
void Solve()
{
    Task("Tree37");
	TNode* P1;
	pt >> P1;
	Recursion(P1);
}

void Recursion(TNode* P)
{
	if (P == NULL)
		return;
	P->Data += (P->Left == NULL) && (P->Right == NULL) ? 1 : -1;
	Recursion(P->Left);
	Recursion(P->Right);
}