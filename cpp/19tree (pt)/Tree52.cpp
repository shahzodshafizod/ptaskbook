#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Tree52");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TNode* node = P1;
	int L = 0;
	bool found = false;
	while (node != NULL)
	{
		if (node == P2)
		{
			found = true;
			break;
		}
		L++;
		node = node->Parent;
	}
	L = !found ? 0 : L;
	pt << L;
}