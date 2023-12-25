#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Push(TNode** queue, int number);

void Solve()
{
    Task("Dynamic16");
	TNode *yakum = NULL, *duyum = NULL;
	TNode *P1 = NULL, *P2 = NULL;
	TNode *P3 = NULL, *P4 = NULL;
	int number;
	for (int i = 1; i < 11; i++)
	{
		pt >> number;
		if (number % 2 != 0)
		{
			Push(&yakum, number);
			if (P1 == NULL)
				P1 = yakum;
		}
		else
		{
			Push(&duyum, number);
			if (P3 == NULL)
				P3 = duyum;
		}
	}
	P2 = yakum;
	P4 = duyum;
	pt << P1 << P2 << P3 << P4;
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