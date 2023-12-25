#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PostfixOutputToN(TNode* P, int& index, int N);

void Solve()
{
    Task("Tree16");
	TNode* P1;
	pt >> P1;
	int N, index = 1;
	pt >> N;
	PostfixOutputToN(P1, index, N);
}

void PostfixOutputToN(TNode* P, int& index, int N)
{
	if (P == NULL)
		return;
	PostfixOutputToN(P->Left, index, N);
	PostfixOutputToN(P->Right, index, N);
	if (index >= N)
		pt << P->Data;
	index++;
}