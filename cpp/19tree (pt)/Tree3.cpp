#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DataKCounter(TNode* pointer, int K, int& counts);
int GetDataKCount(TNode*, int);

void Solve()
{
    Task("Tree3");
	TNode* P1;
	pt >> P1;
	int K;
	pt >> K;
	int counts = GetDataKCount(P1, K);
	pt << counts;
}

int GetDataKCount(TNode* P, int K)
{
	return P == NULL ? 0 : int(P->Data == K) + GetDataKCount(P->Left, K) + GetDataKCount(P->Right, K);
}

void DataKCounter(TNode* pointer, int K, int& counts)
{
	if (pointer == NULL)
		return;
	if (pointer->Data == K)
		counts++;
	DataKCounter(pointer->Left, K, counts);
	DataKCounter(pointer->Right, K, counts);
}