#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DeleteLeaves(TNode**);

void Solve()
{
    Task("Tree41");
	TNode* P1;
	pt >> P1;
	DeleteLeaves(&P1);
}

void DeleteLeaves(TNode** P)
{
	if (*P == NULL)
		return;
	TNode* node = (*P)->Left;
	if (node != NULL)
	{
		if ( (node->Left == NULL) && (node->Right == NULL) )
		{
			delete node;
			(*P)->Left = NULL;
		}
		else
			DeleteLeaves(&node);
	}
	node = (*P)->Right;
	if (node != NULL)
	{
		if ( (node->Left == NULL) && (node->Right == NULL) )
		{
			delete node;
			(*P)->Right = NULL;
		}
		else
			DeleteLeaves(&node);
	}
}