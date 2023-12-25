#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Copy(TNode* parent, TNode *from, bool toLeft = false);
void DeleteTree(TNode**);

void Solve()
{
    Task("Tree55");
	TNode* P1;
	pt >> P1;
	TNode* parent = P1->Parent;
	if (parent != NULL)
	{
		if (parent->Left == P1)
		{
			//i am in left
			if (parent->Right != NULL)
			{
				//delete right
				TNode* right = parent->Right;
				DeleteTree(&right);
				parent->Right = NULL;
			}
			else
			{
				//copy me to right
				Copy(parent, P1);
			}
		}
		else
		{
			//i am in right
			if (parent->Left != NULL)
			{
				//delete left
				TNode* left = parent->Left;
				DeleteTree(&left);
				parent->Left = NULL;
			}
			else
			{
				//copy me to left
				Copy(parent, P1, true);
			}
		}
	}
	pt << parent;
}

void Copy(TNode* parent, TNode *from, bool toLeft)
{
	if (from == NULL)
		return;
	if (toLeft)
	{
		parent->Left = new TNode;
		parent->Left->Left = parent->Left->Right = NULL;
		parent->Left->Data = from->Data;
		parent->Left->Parent = parent;
		Copy(parent->Left, from->Left, true);
		Copy(parent->Left, from->Right);
	}
	else
	{
		parent->Right = new TNode;
		parent->Right->Left = parent->Right->Right = NULL;
		parent->Right->Data = from->Data;
		parent->Right->Parent = parent;
		Copy(parent->Right, from->Left, true);
		Copy(parent->Right, from->Right);
	}
}

void DeleteTree(TNode** P)
{
	if (*P == NULL)
		return;
	TNode* left = (*P)->Left;
	DeleteTree(&left);
	TNode* right = (*P)->Right;
	DeleteTree(&right);
	delete *P;
	*P = NULL;
}