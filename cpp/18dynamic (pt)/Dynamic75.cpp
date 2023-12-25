#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

void LBInsertFirst(TListB&, int);

void Solve()
{
    Task("Dynamic75");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	int N, number;
	pt >> N;
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		LBInsertFirst(lb, number);
	}
	pt << lb.Current;
}

void LBInsertFirst(TListB& L, int D)
{
	TNode* newNode = new TNode;
	newNode->Data = D;
	newNode->Prev = L.Barrier;
	newNode->Next = L.Barrier->Next;
	L.Barrier->Next = newNode;
	newNode->Next->Prev = newNode;
	L.Current = newNode;
}