#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic58");
	int K;
	pt >> K;
	TNode* P1, *P2;
	pt >> P1 >> P2;
	P1->Prev = P2;
	P2->Next = P1;
	TNode* P4 = P1;
	for (int i = 1; i < K; i++)
		P4 = P4->Next;
	TNode* P3 = P4->Next;
	P3->Prev = NULL;
	P4->Next = NULL;
	pt << P3 << P4;
}
