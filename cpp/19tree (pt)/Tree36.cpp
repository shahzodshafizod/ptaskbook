#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Recursion(TNode* P);
void Solve()
{
    Task("Tree36");
	TNode* P1;
	pt >> P1;
	Recursion(P1);
}

void Recursion(TNode* P)
{
	if (P == NULL)
		return;
	if (P->Data % 2 == 0)
		P->Data /= 2;
	Recursion(P->Left);
	Recursion(P->Right);
}