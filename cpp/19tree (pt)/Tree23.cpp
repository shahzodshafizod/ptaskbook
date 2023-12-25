#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetMin(TNode* P);
TNode* GetNode(TNode* P, int data);

void Solve()
{
    Task("Tree23");
	TNode *P1;
	pt >> P1;
	int minimal = GetMin(P1);
	TNode* P2 = GetNode(P1, minimal);
	pt << P2;
}

int GetMin(TNode* P)
{
	int minimal = P->Data;
	if (P->Left != NULL)
	{
		int left = GetMin(P->Left);
		if (left < minimal) minimal = left;
	}
	if (P->Right != NULL)
	{
		int right = GetMin(P->Right);
		if (right < minimal) minimal = right;
	}
	return minimal;
}

TNode* GetNode(TNode* P, int data)
{
	if (P == NULL)
		return P;
	if (P->Data == data)
		return P;
	TNode* node = GetNode(P->Left, data);
	node = node != NULL ? node : GetNode(P->Right, data);
	return node;
}