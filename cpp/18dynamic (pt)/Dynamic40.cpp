#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutDoubleAfter(TNode* P);

void Solve()
{
    Task("Dynamic40");
	TNode* P1;
	pt >> P1;
	TNode *prev, *next;
	for (prev = next = P1; next != NULL; )
	{
		if (next->Data % 2 != 0)
		{
			PutDoubleAfter(next);
			if (next != NULL)
			{
				prev = next;
				next = next->Next;
			}
		}
		if (next != NULL)
		{
			prev = next;
			next = next->Next;
		}
	}
	pt << prev;
}

void PutDoubleAfter(TNode* P)
{
	TNode* pointer = new TNode;
	pointer->Data = P->Data;
	pointer->Prev = P;
	pointer->Next = P->Next;
	if (P->Next != NULL)
		P->Next->Prev = pointer;
	P->Next = pointer;
}