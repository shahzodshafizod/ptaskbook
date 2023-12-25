#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

int LBDeleteCurrent(TListB&);
bool IsBarrier(TListB&);

void Solve()
{
    Task("Dynamic80");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	for (int i = 1; !IsBarrier(lb) && (i <= 5); i++)
		pt << LBDeleteCurrent(lb);
	
	pt << lb.Current;
}

int LBDeleteCurrent(TListB& L)
{
	int data = 0;
	if (!IsBarrier(L))
	{
		TNode* node = L.Current;
		data = node->Data;
		L.Current = L.Current->Next != L.Barrier ? L.Current->Next : L.Current->Prev;
		node->Next->Prev = node->Prev;
		node->Prev->Next = node->Next;
		delete node;
	}
	return data;
}

bool IsBarrier(TListB& L)
{
	return L.Barrier == L.Current;
}