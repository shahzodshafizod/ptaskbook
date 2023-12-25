#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DataSum(TNode*, int&);
int GetDataSum(TNode*);

void Solve()
{
    Task("Tree4");
	TNode* P1;
	pt >> P1;
	int sum = GetDataSum(P1);
	pt << sum;
}

int GetDataSum(TNode* P)
{
	return P == NULL ? 0 : P->Data + GetDataSum(P->Left) + GetDataSum(P->Right);
}

void DataSum(TNode* pointer, int& sum)
{
	if (pointer == NULL)
		return;
	sum += pointer->Data;
	DataSum(pointer->Left, sum);
	DataSum(pointer->Right, sum);
}