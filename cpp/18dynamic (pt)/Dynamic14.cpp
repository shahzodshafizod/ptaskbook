#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Push(TNode** queue, int number);

void Solve()
{
	Task("Dynamic14");
	int number;
	TNode* queue = NULL;
	TNode *P1, *P2;
	for (int i = 1; i <= 10; i++)
	{
		pt >> number;
		Push(&queue, number);
		if (i == 1)
			P1 = queue;
	}
	P2 = queue;
	pt << P1 << P2;
}

void Push(TNode** queue, int number)
{
	TNode* temp = new TNode;
	temp->Data = number;
	temp->Next = NULL;
	
	if (*queue != NULL)
		(*queue)->Next = temp;
	
	*queue = temp;
	temp = NULL;
}