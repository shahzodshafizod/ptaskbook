#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** parent, int& l, int L, int* mass, int& index, int N, bool left = false);

void Solve()
{
    Task("Tree56");
	int L, N;
	pt >> L >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		pt >> mass[i];
	TNode* P1 = NULL;
	int l = 0, index = 0;
	MakeTree(&P1, l, L, mass, index, N);
	pt << P1;
}

void MakeTree(TNode** parent, int& l, int L, int* mass, int& index, int N, bool left)
{
	if (l > L)
		return;
	if (index >= N)
		return;
	l++;
	TNode* node = new TNode;
	node->Parent = *parent;
	node->Left = node->Right = NULL;
	node->Data = mass[index++];
	if (*parent == NULL)
		*parent = node;
	else if (left)
		(*parent)->Left = node;
	else
		(*parent)->Right = node;
	MakeTree(&node, l, L, mass, index, N, true);
	MakeTree(&node, l, L, mass, index, N);
	l--;
}