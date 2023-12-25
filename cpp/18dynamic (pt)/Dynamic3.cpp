#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic3");
	int D;
	pt >> D;
	TNode* P1;
	pt >> P1;
	TNode* P0 = new TNode;
	P0->Data = D;
	P0->Next = P1;
	pt << P0;
}
