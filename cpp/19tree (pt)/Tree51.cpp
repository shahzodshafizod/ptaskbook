#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int GetLevel(TNode*);

void Solve()
{
    Task("Tree51");
	TNode* P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	int firstLevel = GetLevel(P1);
	int secondLevel = GetLevel(P2);
	int thirdLevel = GetLevel(P3);
	pt << firstLevel << secondLevel << thirdLevel;
}

int GetLevel(TNode* P)
{
	if (P->Parent == NULL)
		return 0;
	return 1 + GetLevel(P->Parent);
}