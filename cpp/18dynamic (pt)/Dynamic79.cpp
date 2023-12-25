#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TListB
{
	TNode* Barrier;
	TNode* Current;
};

void LBToLast(TListB&);
void LBToPrev(TListB&);
int LBGetData(const TListB&);
bool IsBarrier(const TListB&);

void Solve()
{
    Task("Dynamic79");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TListB lb = {P1, P2};
	LBToLast(lb);
	int data, counts = 0;
	while (!IsBarrier(lb))
	{
		counts++;
		data = LBGetData(lb);
		if (data % 2 == 0)
			pt << data;
		LBToPrev(lb);
	}
	pt << counts;
}

void LBToLast(TListB& L)
{
	L.Current = L.Barrier->Prev;
}

void LBToPrev(TListB& L)
{
	L.Current = L.Current->Prev;
}

int LBGetData(const TListB& L)
{
	return L.Current->Data;
}

bool IsBarrier(const TListB& L)
{
	return L.Current == L.Barrier;
}