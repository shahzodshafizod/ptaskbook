#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic5");
	TNode *P1, *P2;
	pt >> P1;
	int D = P1->Data;
	P2 = P1->Next;
	delete P1;
	pt << D;
	pt << P2;
}
