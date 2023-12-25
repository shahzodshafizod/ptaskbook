#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

LBInsertLast(TListB& L, int D);

void Solve()
{
    Task("Dynamic74");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	int N, number;
	pt >> N;
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		LBInsertLast(lb, number);
	}
	pt << lb.Current;
}

LBInsertLast(TListB& L, int D)
{
	TNode* newNode = new TNode;
	newNode->Data = D;
	newNode->Next = L.Barrier;
	newNode->Prev = L.Barrier->Prev;
	L.Barrier->Prev = newNode;
	newNode->Prev->Next = newNode;
	L.Current = newNode;
}