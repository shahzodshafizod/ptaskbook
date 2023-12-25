#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MakeTree(TNode**, int);

void Solve()
{
    Task("Tree29");
	int N;
	pt >> N;
	TNode* P1 = NULL;
	MakeTree(&P1, N);
	pt << P1;
}

void MakeTree(TNode** P, int N)
{
	TNode* node = *P;
	for (int i = 1; i <= N; i++)
	{
		TNode* temp = new TNode;
		pt >> temp->Data;
		temp->Left = temp->Right = NULL;
		if (node == NULL)
		{
			*P = node = temp;
			continue;
		}
		if ( (node->Left == NULL) && (node->Right == NULL) )
		{
			if (node->Data % 2 != 0)
				node->Left = temp;
			else
				node->Right = temp;
		}
		else
		{
			if (node->Left == NULL)
				node->Left = temp;
			else
				node->Right = temp;
			node = temp;
		}
	}
}