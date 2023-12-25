#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void ChangeChildrenDatas(TNode* P);
void Solve()
{
    Task("Tree38");
	TNode *P1;
	pt >> P1;
	ChangeChildrenDatas(P1);
}

void ChangeChildrenDatas(TNode* P)
{
	if (P == NULL)
		return;
	if ( (P->Left != NULL) && (P->Right != NULL) )
	{
		int temp = P->Left->Data;
		P->Left->Data = P->Right->Data;
		P->Right->Data = temp;
	}
	ChangeChildrenDatas(P->Left);
	ChangeChildrenDatas(P->Right);
}