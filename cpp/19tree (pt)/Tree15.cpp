#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void InfixOutputToN(TNode* P, int& index, int N);

void Solve()
{
    Task("Tree15");
	TNode* P1;
	pt >> P1;
	int N, index = 1;
	pt >> N;
	InfixOutputToN(P1, index, N);
}

void InfixOutputToN(TNode* P, int& index, int N)
{
	if (P == NULL)
		return;
	InfixOutputToN(P->Left, index, N);
	if (index <= N)
		pt << P->Data;
	index++;
	InfixOutputToN(P->Right, index, N);
}