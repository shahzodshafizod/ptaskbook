#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic6");
	TNode *P1, *P2;
	pt >> P1;
	for (int i = 0; i < 9; i++)
	{
		pt << P1->Data;
		P2 = P1;
		P1 = P1->Next;
		delete P2;
	}
	pt << P1;
}
