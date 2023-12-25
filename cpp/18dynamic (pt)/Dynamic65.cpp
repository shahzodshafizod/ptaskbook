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

int DeleteCurrent(TList& L);

void Solve()
{
    Task("Dynamic65");
	TNode *P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list = {P1, P2, P3};
	for (int i = 1; i <= 5; i++)
		pt << DeleteCurrent(list);
	pt << list.First << list.Last << list.Current;
}

int DeleteCurrent(TList& L)
{
	TNode* next = L.Current->Next;
	int data = L.Current->Data;

	if (L.Current == L.First)
	{
		L.First = L.First->Next;
	}
	if (L.Current == L.Last)
	{
		L.Last = L.Last->Prev;
	}

	if (L.Current->Next != NULL)
	{
		L.Current->Next->Prev = L.Current->Prev;
	}
	if (L.Current->Prev != NULL)
	{
		L.Current->Prev->Next = L.Current->Next;
	}

	delete L.Current;
	L.Current = next != NULL ? next : L.Last;

	return data;
}