#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void ChangeChildrenDatas(TNode* P);
void Solve()
{
    Task("Tree39");
	TNode* P1;
	pt >> P1;
	ChangeChildrenDatas(P1);
}

void ChangeChildrenDatas(TNode* P)
{
	if (P == NULL)
		return;
	if ( (P->Left != NULL) || (P->Right != NULL) )
	{
		TNode* temp = P->Left;
		P->Left = P->Right;
		P->Right = temp;
	}
	ChangeChildrenDatas(P->Left);
	ChangeChildrenDatas(P->Right);
}