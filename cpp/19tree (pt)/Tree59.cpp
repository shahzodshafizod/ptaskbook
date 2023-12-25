#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* Search(TNode* P, int K, int& N);

void Solve()
{
    Task("Tree59");
	TNode* P1;
	pt >> P1;
	int K;
	pt >> K;
	int N = 0;
	TNode* P2 = Search(P1, K, N);
	pt << P2;
	pt << N;
}

TNode* Search(TNode* P, int K, int& N)
{
	if (P == NULL)
		return P;
	N++;
	if (P->Data == K)
		return P;
	
	return ( P->Data > K ? Search(P->Left, K, N) : Search(P->Right, K, N) );
}