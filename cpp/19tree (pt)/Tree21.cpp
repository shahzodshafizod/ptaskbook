#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetMinimalInLeaves(TNode*);

void Solve()
{
    Task("Tree21");
	TNode* P1;
	pt >> P1;
	int minimal = GetMinimalInLeaves(P1);
	pt << minimal;
}

int GetMinimalInLeaves(TNode* P)
{
	if ( (P->Left == NULL) && (P->Right == NULL) )
		return P->Data;
	if (P->Left == NULL)
		return GetMinimalInLeaves(P->Right);
	if (P->Right == NULL)
		return GetMinimalInLeaves(P->Left);
	
	int left = GetMinimalInLeaves(P->Left);
	int right = GetMinimalInLeaves(P->Right);
	return left < right ? left : right;
}