#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Search(TNode* P, int K, int& C, int &N);

void Solve()
{
    Task("Tree60");
	TNode* P1;
	pt >> P1;
	int K;
	pt >> K;
	int C = 0, N = 0;
	Search(P1, K, C, N);
	pt << C << N;
}

void Search(TNode* P, int K, int& C, int &N)
{
	if (P == NULL)
		return;
	N++;
	if (P->Data > K)
		Search(P->Left, K, C, N);
	else if (P->Data < K)
		Search(P->Right, K, C, N);
	else
	{
		C++;
		Search(P->Left, K, C, N);
		Search(P->Right, K, C, N);
	}
}