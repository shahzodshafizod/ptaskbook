#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int& l, int L, int& n, int N, bool right = false);

void Solve()
{
    Task("Tree31");
	int L, N;
	pt >> L >> N;
	TNode* P1 = NULL;
	int l = 0, n = 0;
	MakeTree(&P1, l, L, n, N);
	pt << P1;
}

void MakeTree(TNode** P, int& l, int L, int& n, int N, bool right)
{
	if ( (n == N) || (l > L) )
		return;

	l++;
	TNode* node = new TNode;
	pt >> node->Data;
	n++;
	node->Left = node->Right = NULL;
	if (*P == NULL)
		*P = node;
	else if (right)	
		(*P)->Right = node;
	else
		(*P)->Left = node;

	MakeTree(&node, l, L, n, N);
	MakeTree(&node, l, L, n, N, true);
	l--;
}