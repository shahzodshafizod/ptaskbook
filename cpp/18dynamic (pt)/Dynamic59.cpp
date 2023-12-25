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

void InsertLast(TList& L, int D);

void Solve()
{
    Task("Dynamic59");
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
		InsertLast(list, number);
	}
	pt << list.First << list.Last << list.Current;
}

void InsertLast(TList& L, int D)
{
	L.Current = L.Last;
	L.Last = new TNode;
	L.Last->Data = D;
	L.Last->Next = NULL;
	L.Last->Prev = L.Current;
	if (L.Current != NULL)
		L.Current->Next = L.Last;
	else
		L.First = L.Last;
	L.Current = L.Last;
}