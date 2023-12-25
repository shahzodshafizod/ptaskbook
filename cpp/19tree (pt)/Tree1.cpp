#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Tree1");
	TNode* P1;
	pt >> P1;
	pt << P1->Data;
	pt << P1->Left->Data << P1->Right->Data;
	pt << P1->Left << P1->Right;
}
