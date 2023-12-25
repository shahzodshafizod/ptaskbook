#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetLevel(TNode* P);
void FillTree(TNode* P, int& l, int L);

void Solve()
{
    Task("Tree47");
	TNode* P1;
	pt >> P1;
	int L = GetLevel(P1);
	int l = 1;
	FillTree(P1, l, L);
}

int GetLevel(TNode* P)
{
	if (P == NULL)
		return 0;
	int L = 1 + GetLevel(P->Left);
	int R = 1 + GetLevel(P->Right);
	if (L > R)
		return L;
	else
		return R;
}

void FillTree(TNode* P, int& l, int L)
{
	if (l >= L)
		return;
	l++;
	
	if (P->Left == NULL)
	{
		P->Left = new TNode;
		P->Left->Left = P->Left->Right = NULL;
		P->Left->Data = -1;
	}
	FillTree(P->Left, l, L);
	
	if (P->Right == NULL)
	{
		P->Right = new TNode;
		P->Right->Left = P->Right->Right = NULL;
		P->Right->Data = -1;
	}
	FillTree(P->Right, l, L);
	
	l--;
}