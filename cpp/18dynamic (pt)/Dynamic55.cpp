#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic55");
	TNode* P1;
	pt >> P1;
	TNode* P2 = P1;
	while (P2->Next != NULL)
		P2 = P2->Next;
	P1->Prev = P2;
	P2->Next = P1;
	
	pt << P2;
}
