#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutDoubleAfter(TNode* P);

void Solve()
{
    Task("Dynamic38");
	TNode* P1;
	pt >> P1;
	TNode* next, *prev;
	for (next = prev = P1; next != NULL; )
	{
		PutDoubleAfter(next);
		for (int j = 0; j < 3; j++)
		{
			if (next != NULL)
			{
				prev = next;
				next = next->Next;
			}
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