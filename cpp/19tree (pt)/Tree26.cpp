#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N);
void Solve()
{
    Task("Tree26");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N)
{
	bool left = true;
	TNode* node = *P;
	for (int i = 1; i <= N; i++)
	{
		left = !left;
		TNode* temp = new TNode;
		pt >> temp->Data;
		temp->Left = temp->Right = NULL;
		if (node == NULL)
			*P = temp;
		else if (left)
			node->Left = temp;
		else
			node->Right = temp;
		node = temp;
	}
}