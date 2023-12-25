#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Tree48");
	TNode* P1;
	pt >> P1;
	TNode* PL = P1->Left;
	TNode* PR = P1->Right;
	TNode* P0 = P1->Parent;
	TNode* P2 = P0 == NULL ? NULL : P0->Left == P1 ? P0->Right : P0->Left;
	pt << PL << PR << P0 << P2;
}
