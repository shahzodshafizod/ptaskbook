#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutDoubleBefore(TNode* P);

void Solve()
{
    Task("Dynamic39");
	TNode* P1;
	pt >> P1;
	for (TNode* i = P1; i != NULL; i = i->Next)
		if (i->Data % 2 != 0)
			PutDoubleBefore(i);
	if (P1->Prev != NULL)
		P1 = P1->Prev;
	pt << P1;
}

void PutDoubleBefore(TNode* P)
{
	TNode* pointer = new TNode;
	pointer->Data = P->Data;
	pointer->Next = P;
	pointer->Prev = P->Prev;
	if (P->Prev != NULL)
		P->Prev->Next = pointer;
	P->Prev = pointer;
}