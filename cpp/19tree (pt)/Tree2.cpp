#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void NodesCount(TNode*, int&);
int GetNodesCount(TNode*);

void Solve()
{
    Task("Tree2");
	TNode* P1;
	pt >> P1;
	int counts = GetNodesCount(P1);
	pt << counts;
}

int GetNodesCount(TNode* P)
{
	return P == NULL ? 0 : 1 + GetNodesCount(P->Left) + GetNodesCount(P->Right);
}

void NodesCount(TNode* pointer, int& counts)
{
	if (pointer == NULL)
		return;
	counts++;
	NodesCount(pointer->Left, counts);
	NodesCount(pointer->Right, counts);
}