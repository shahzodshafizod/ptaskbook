#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetMaximal(TNode*);
int Maximal(int, int, int);
int Maximal(int, int);

void Solve()
{
    Task("Tree22");
	TNode* P1;
	pt >> P1;
	int maximal = GetMaximal(P1);
	pt << maximal;
}

int GetMaximal(TNode* P)
{
	int left, right, data = P->Data;
	bool hasLeft = false, hasRight = false;
	if (P->Left != NULL)
	{
		if ( (P->Left->Left != NULL) || (P->Left->Right != NULL) )
		{
			hasLeft = true;
			left = GetMaximal(P->Left);
		}
	}
	if (P->Right != NULL)
	{
		if ( (P->Right->Left != NULL) || (P->Right->Right != NULL) )
		{
			hasRight = true;
			right = GetMaximal(P->Right);
		}
	}
	int maximal = hasLeft && hasRight ? Maximal(data, left, right) :
							hasLeft ? Maximal(data, left) :
							hasRight ? Maximal(data, right) :
							data;
	return maximal;
}

int Maximal(int a, int b, int c)
{
	return Maximal( Maximal(a, b), Maximal(b, c) );
}

int Maximal(int a, int b)
{
	return a > b ? a : b;
}