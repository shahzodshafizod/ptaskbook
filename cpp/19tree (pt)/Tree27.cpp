#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N);

void Solve()
{
    Task("Tree27");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N)
{
	bool parent, left = false;
	TNode* node = *P;
	for (int i = 1; i <= N; i++)
	{
		TNode* temp = new TNode;
		pt >> temp->Data;
		parent = left;
		left = temp->Data % 2 != 0;
		temp->Left = temp->Right = NULL;
		if (node == NULL)
			*P = temp;
		else if (parent)
			node->Left = temp;
		else
			node->Right = temp;
		node = temp;
	}
}