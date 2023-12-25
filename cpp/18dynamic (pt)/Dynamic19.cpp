#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

bool Pop(TNode** node, int& D);

void Solve()
{
    Task("Dynamic19");
	int N;
	pt >> N;
	TNode* P1, *P2;
	pt >> P1 >> P2;
	int D;
	for (int i = 1; i <= N; i++)
	{
		if (Pop(&P1, D))
			pt << D;
		else
			break;
	}
	if (P1 == NULL)
		P2 = NULL;

	pt << P1 << P2;
}

bool Pop(TNode** node, int& D)
{
	if (*node != NULL)
	{
		TNode* temp = *node;
		*node = (*node)->Next;
		D = temp->Data;
		delete temp;
		return true;
	}
	
	return false;
}