#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode** P, int N);

void Solve()
{
    Task("Tree28");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N)
{
	TNode* node = *P;
	bool lefted = false;
	for (int i = 1; i <= N; i++)
	{
		TNode* temp = new TNode;
		pt >> temp->Data;
		temp->Left = temp->Right = NULL;
		if (node == NULL)
			*P = node = temp;
		else if (lefted)
		{
			node->Right = temp;
			lefted = false;
			node = temp;
		}
		else
		{
			node->Left = temp;
			lefted = true;
		}
	}
}