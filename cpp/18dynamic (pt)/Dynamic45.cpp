#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

bool CanMovePrev(TNode* P);
void MovePrev(TNode* P);

void Solve()
{
    Task("Dynamic45");
	TNode* P0;
	pt >> P0;
	
	TNode* P2 = P0;
	while (P2->Next != NULL)
		P2 = P2->Next;
	
	if (P2 == P0)
	{
		P2 = P2->Prev;
		P2->Next = NULL;
	}

	TNode* P1 = P0;
	while (CanMovePrev(P1))
		MovePrev(P1);
	
	pt << P1 << P2;
}

bool CanMovePrev(TNode* P)
{
	return ( (P != NULL) && (P->Prev != NULL) );
}

void MovePrev(TNode* P)
{
	if (P->Next != NULL)
		P->Next->Prev = P->Prev;
	
	if (P->Prev != NULL)
	{
		P->Prev->Next = P->Next;
		if (P->Prev->Prev != NULL)
			P->Prev->Prev->Next = P;
		P->Next = P->Prev;
		P->Prev = P->Prev->Prev;
		P->Next->Prev = P;
	}
}