#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutDoubleBefore(TNode* P);

void Solve()
{
    Task("Dynamic35");
	TNode *P1, *P2;
	pt >> P1 >> P2;

	PutDoubleBefore(P1);
	PutDoubleBefore(P2);
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