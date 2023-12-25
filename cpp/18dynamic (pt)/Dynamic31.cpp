#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic31");
	TNode* P0;
	pt >> P0;
	TNode* P1, *P2;
	P1 = P2 = P0;
	int N = 0;
	while ( (P1 != NULL) && (P1->Prev != NULL) )
	{
		P1 = P1->Prev;
		N++;
	}
	while ( (P2 != NULL) && (P2->Next != NULL) )
	{
		P2 = P2->Next;
		N++;
	}
	N++; // !counting the element P0!
	pt << N;
	pt << P1 << P2;
}
