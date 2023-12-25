#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic57");
	int K;
	pt >> K;
	TNode* P1, *P2;
	pt >> P1 >> P2;
	P1->Prev = P2;
	P2->Next = P1;
	TNode* P3 = P2;
	for (int i = 1; i < K; i++)
		P3 = P3->Prev;
	TNode* P4 = P3->Prev;
	P3->Prev = NULL;
	P4->Next = NULL;
	pt << P3 << P4;
}
