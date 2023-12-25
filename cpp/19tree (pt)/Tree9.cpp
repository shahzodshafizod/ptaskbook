#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void HeightCounter(TNode* P, int& height, int& maxHeight);

void Solve()
{
    Task("Tree9");
	TNode* P1;
	pt >> P1;
	int height = 0, maxHeight = 0;
	HeightCounter(P1, height, maxHeight);
	maxHeight--;
	pt << maxHeight;
}

void HeightCounter(TNode* P, int& height, int& maxHeight)
{
	if (P == 0)
		return;
	height++;
	if (height > maxHeight)
		maxHeight = height;
	HeightCounter(P->Left, height, maxHeight);
	HeightCounter(P->Right, height, maxHeight);
	height--;
}