#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N, int& level, bool right = false);

void Solve()
{
    Task("Tree33");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	int level = -1;
	MakeTree(&P1, N, level);
	pt << P1;
}

void MakeTree(TNode** P, int N, int& level, bool right)
{
	if (N == 0)
		return;
	level++;
	TNode* node = new TNode;
	node->Data = level;
	node->Left = node->Right = NULL;
	if (*P == NULL)
		*P = node;
	else if (right)
		(*P)->Right = node;
	else
		(*P)->Left = node;
	MakeTree(&node, N/2, level);
	MakeTree(&node, N-1-N/2, level, true);
	level--;
}