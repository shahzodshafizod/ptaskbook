#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void AddAnotherChild(TNode*);

void Solve()
{
    Task("Tree46");
	TNode* P1;
	pt >> P1;
	AddAnotherChild(P1);
}

void AddAnotherChild(TNode* P)
{
	if (P == NULL)
		return;
	if ( (P->Left == NULL) && (P->Right != NULL) )
	{
		//add left
		P->Left = new TNode;
		P->Left->Left = P->Left->Right = NULL;
		P->Left->Data = 2 * P->Data;

		AddAnotherChild(P->Right);
	}
	else if ( (P->Left != NULL) && (P->Right == NULL) )
	{
		//add right
		P->Right = new TNode;
		P->Right->Left = P->Right->Right = NULL;
		P->Right->Data = 2 * P->Data;

		AddAnotherChild(P->Left);
	}
	else
	{
		AddAnotherChild(P->Left);
		AddAnotherChild(P->Right);
	}
}