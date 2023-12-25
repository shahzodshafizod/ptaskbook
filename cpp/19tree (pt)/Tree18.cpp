#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PrefixLevelL(TNode* P, int& level, int L, int& counts);
void PostfixLevelL(TNode* P, int& level, int L, int& counts);
void InfixLevelL(TNode* P, int& level, int L, int& counts);

void Solve()
{
    Task("Tree18");
	TNode* P1;
	pt >> P1;
	int L;
	pt >> L;
	int level = 0, counts = 0;
	//PrefixLevelL(P1, level, L, counts);
	//PostfixLevelL(P1, level, L, counts);
	InfixLevelL(P1, level, L, counts);
	pt << counts;
}

void PrefixLevelL(TNode* P, int& level, int L, int& counts)
{
	if (P == NULL)
		return;
	level++;
	
	if (level-1 == L)
	{
		pt << P->Data;
		counts++;
	}
	PrefixLevelL(P->Left, level, L, counts);
	PrefixLevelL(P->Right, level, L, counts);
	
	level--;
}

void PostfixLevelL(TNode* P, int& level, int L, int& counts)
{
	if (P == NULL)
		return;
	level++;
	
	PostfixLevelL(P->Left, level, L, counts);
	PostfixLevelL(P->Right, level, L, counts);
	if (level-1 == L)
	{
		pt << P->Data;
		counts++;
	}
	
	level--;
}

void InfixLevelL(TNode* P, int& level, int L, int& counts)
{
	if (P == NULL)
		return;
	level++;
	
	InfixLevelL(P->Left, level, L, counts);
	if (level-1 == L)
	{
		pt << P->Data;
		counts++;
	}
	InfixLevelL(P->Right, level, L, counts);
	
	level--;
}