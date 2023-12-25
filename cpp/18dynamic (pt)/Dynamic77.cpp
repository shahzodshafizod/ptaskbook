#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

void LBInsertAfter(TListB&, int);

void Solve()
{
    Task("Dynamic77");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	int number;
	for (int i = 1; i <= 5; i++)
	{
		pt >> number;
		LBInsertAfter(lb, number);
	}
	pt << lb.Current;
}

void LBInsertAfter(TListB& L, int D)
{
	TNode* newNode = new TNode;
	newNode->Data = D;
	newNode->Prev = L.Current;
	newNode->Next = L.Current->Next;
	L.Current->Next = newNode;
	newNode->Next->Prev = newNode;
	L.Current = newNode;
}