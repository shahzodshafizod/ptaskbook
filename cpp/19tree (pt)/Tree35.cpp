#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DoubleTreeDatas(TNode* P);
void Solve()
{
    Task("Tree35");
	TNode* P1;
	pt >> P1;
	DoubleTreeDatas(P1);
}

void DoubleTreeDatas(TNode* P)
{
	if (P == NULL)
		return;
	P->Data *= 2;
	DoubleTreeDatas(P->Left);
	DoubleTreeDatas(P->Right);
}