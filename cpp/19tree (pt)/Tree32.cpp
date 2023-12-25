#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N, bool right = false);

void Solve()
{
    Task("Tree32");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N, bool right)
{
	if (N == 0)
		return;
	TNode* node = new TNode;
	pt >> node->Data;
	node->Left = node->Right = NULL;
	if (*P == NULL)
		*P = node;
	else if (right)
		(*P)->Right = node;
	else
		(*P)->Left = node;
	MakeTree(&node, N/2);
	MakeTree(&node,  N-1-N/2, true);
}