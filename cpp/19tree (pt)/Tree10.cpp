#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void NodesInLevels(TNode* P, int& height, int& maxHeight, int* mass);
void Solve()
{
    Task("Tree10");
	TNode* P1;
	pt >> P1;
	int height = 0, maxHeight = 0;
	const int limit = 10;
	int mass[limit] = {0};
	NodesInLevels(P1, height, maxHeight, mass);
	for (int i = 0; i < maxHeight; i++)
		pt << mass[i];
}

void NodesInLevels(TNode* P, int& height, int& maxHeight, int* mass)
{
	if (P == NULL)
		return;
	height++;
	if (height > maxHeight)
		maxHeight = height;
	mass[height-1]++;
	NodesInLevels(P->Left, height, maxHeight, mass);
	NodesInLevels(P->Right, height, maxHeight, mass);
	height--;
}