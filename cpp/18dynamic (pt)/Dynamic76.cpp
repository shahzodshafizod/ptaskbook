#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

void LBInsertBefore(TListB&, int);

void Solve()
{
    Task("Dynamic76");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	int number;
	for (int i = 1; i <= 5; i++)
	{
		pt >> number;
		LBInsertBefore(lb, number);
	}
	pt << lb.Current;
}

void LBInsertBefore(TListB& L, int D)
{
	TNode* newNode = new TNode;
	newNode->Data = D;
	newNode->Next = L.Current;
	newNode->Prev = L.Current->Prev;
	L.Current->Prev = newNode;
	newNode->Prev->Next = newNode;
	L.Current = newNode;
}