#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic32");
	int D1, D2;
	pt >> D1 >> D2;
	TNode* P0;
	pt >> P0;
	TNode *P1 = P0, *P2 = P0;
	//P1 points to the begin of the list
	while (P1->Prev != NULL)
		P1 = P1->Prev;
	//P1 points to the end of the list
	while (P2->Next != NULL)
		P2 = P2->Next;
	TNode* P = new TNode;
	P->Data = D1;
	P->Prev = NULL;
	P->Next = P1;
	P1->Prev = P;
	P1 = P1->Prev;

	P = new TNode;
	P->Data = D2;
	P->Next = NULL;
	P->Prev = P2;
	P2->Next = P;
	P2 = P2->Next;

	pt << P1 << P2;
}
