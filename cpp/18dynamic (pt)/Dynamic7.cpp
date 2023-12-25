#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic7");
	TNode *P1;
	pt >> P1;
	int N = 0;
	while (P1 != NULL)
	{
		TNode* temp = P1;
		pt << P1->Data;
		P1 = P1->Next;
		delete temp;
		N++;
	}
	pt << N;
}
