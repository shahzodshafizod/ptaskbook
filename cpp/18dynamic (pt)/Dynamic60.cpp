#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TList
{
	TNode* First;
	TNode* Last;
	TNode* Current;
};

void InsertFirst(TList& L, int D);

void Solve()
{
    Task("Dynamic60");
	TNode* P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list = {P1, P2, P3};
	/*
	list.First = P1;
	list.Last = P2;
	list.Current = P3;
	*/
	int N, number;
	pt >> N;
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		InsertFirst(list, number);
	}
	pt << list.First << list.Last << list.Current;
}

void InsertFirst(TList& L, int D)
{
	L.Current = L.First;
	L.First = new TNode;
	L.First->Data = D;
	L.First->Prev = NULL;
	L.First->Next = L.Current;
	if (L.Current != NULL)
		L.Current->Prev = L.First;
	else
		L.Last = L.First;
	L.Current = L.First;
}