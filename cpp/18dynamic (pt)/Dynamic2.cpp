#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic2");
	TNode *P1, *P2;
	pt >> P1;
	int N = 0;
	while (P1 != NULL)
	{
		pt << P1->Data;
		P2 = P1;
		P1 = P1->Next;
		N++;
	}
	pt << N << P2;
}
