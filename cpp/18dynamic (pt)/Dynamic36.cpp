#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutDoubleAfter(TNode* P);

void Solve()
{
    Task("Dynamic36");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P = new TNode;
	
	PutDoubleAfter(P1);
	PutDoubleAfter(P2);
	P2 = P2->Next;

	pt << P2;
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