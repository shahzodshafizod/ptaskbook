#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N, bool right = false);

void Solve()
{
    Task("Tree30");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N, bool right)
{
	TNode* node = new TNode;
	node->Data = N;
	node->Left = node->Right = NULL;
	if (*P == NULL)
		*P = node;
	else if (right)
		(*P)->Right = node;
	else
		(*P)->Left = node;
	
	if (node->Data == 1)
		return;
	MakeTree(&node, N/2);
	if (node->Data % 2 != 0)
	{
		MakeTree(&node, N-N/2, true);
	}
}