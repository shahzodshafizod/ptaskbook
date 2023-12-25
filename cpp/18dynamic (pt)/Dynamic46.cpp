#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

bool CanMoveNext(TNode* P);
void MoveNext(TNode* P);

void Solve()
{
    Task("Dynamic46");
	int K;
	pt >> K;
	TNode* P0;
	pt >> P0;
	
	TNode *P1 = P0;
	while (P1->Prev != NULL)
		P1 = P1->Prev;
	
	if ( (K > 0) && (P1 == P0) )
	{
		P1 = P1->Next;
		P1->Prev = NULL;
	}
	
	int i = 1;
	while ((i <= K) && CanMoveNext(P0))
	{
		MoveNext(P0);
		i++;
	}
	TNode* P2 = P0;
	while (P2->Next != NULL)
		P2 = P2->Next;

	pt << P1 << P2;
}

bool CanMoveNext(TNode* P)
{
	return ( (P != NULL) && (P->Next != NULL) );
}

void MoveNext(TNode* P)
{
	if (P->Prev != NULL)
		P->Prev->Next = P->Next;
	
	if (P->Next != NULL)
	{
		P->Next->Prev = P->Prev;
		
		if (P->Next->Next != NULL)
			P->Next->Next->Prev = P;

		P->Prev = P->Next;
		P->Next = P->Next->Next;
		P->Prev->Next = P;
	}	
}