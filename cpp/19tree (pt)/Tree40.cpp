#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DeleteTree(TNode** P, bool deleteThisNode = false);
void Solve()
{
    Task("Tree40");
	TNode* P1;
	pt >> P1;
	DeleteTree(&P1);
}

void DeleteTree(TNode** P, bool deleteThisNode)
{
	if (*P == NULL)
		return;
	DeleteTree(&((*P)->Left), true);
	DeleteTree(&((*P)->Right), true);
	(*P)->Left = (*P)->Right = NULL;
	if (deleteThisNode)
	{
		delete *P;
		*P = NULL;
	}
}