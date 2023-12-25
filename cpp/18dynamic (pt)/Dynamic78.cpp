#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

void LBToFirst(TListB&);
void LBToNext(TListB&);
void LBSetData(TListB&, int);
bool IsBarrier(const TListB&);

void Solve()
{
    Task("Dynamic78");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	LBToFirst(lb);
	int counts = 0;
	while (!IsBarrier(lb))
	{
		LBSetData(lb, 0);
		for (int i = 1; !IsBarrier(lb) && (i <= 2); i++)
		{
			LBToNext(lb);
			counts++;
		}
	}
	pt << counts;
	pt << lb.Current;
}

void LBToFirst(TListB& L)
{
	L.Current = L.Barrier->Next;
}

void LBToNext(TListB& L)
{
	L.Current = L.Current->Next;
}

void LBSetData(TListB& L, int D)
{
	if (!IsBarrier(L))
		L.Current->Data = D;
}

bool IsBarrier(const TListB& L)
{
	return L.Current == L.Barrier;
}