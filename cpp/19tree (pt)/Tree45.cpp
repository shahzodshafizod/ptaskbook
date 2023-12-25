#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void AddNodeToLeaves(TNode*);

void Solve()
{
    Task("Tree45");
	TNode* P1;
	pt >> P1;
	AddNodeToLeaves(P1);
}

void AddNodeToLeaves(TNode* P)
{
	if (P == NULL)
		return;
	if ( (P->Left == NULL) && (P->Right == NULL) )
	{
		if (P->Data % 2 == 0)
		{
			//add right
			P->Right = new TNode;
			P->Right->Left = P->Right->Right = NULL;
			P->Right->Data = P->Data;
		}
		else
		{
			//add left
			P->Left = new TNode;
			P->Left->Left = P->Left->Right = NULL;
			P->Left->Data = P->Data;
		}
	}
	else
	{
		AddNodeToLeaves(P->Left);
		AddNodeToLeaves(P->Right);
	}
}