#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic1");
	TNode* P1;
	pt >> P1;
	pt << P1->Data << P1->Next->Data;
	pt << P1->Next;
}
