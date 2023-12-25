#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PrefixOutputInN1andN2(TNode* P, int& index, int N1, int N2);

void Solve()
{
    Task("Tree17");
	TNode* P1;
	pt >> P1;
	int N1, N2, index = 1;
	pt >> N1 >> N2;
	PrefixOutputInN1andN2(P1, index, N1, N2);
}

void PrefixOutputInN1andN2(TNode* P, int& index, int N1, int N2)
{
	if (P == NULL)
		return;
	if ( (index >= N1) && (index <= N2) )
		pt << P->Data;
	index++;
	PrefixOutputInN1andN2(P->Left, index, N1, N2);
	PrefixOutputInN1andN2(P->Right, index, N1, N2);
}