#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Push(TNode** from1, TNode** from2, TNode** to1, TNode** to2);

void Solve()
{
    Task("Dynamic24");
	TNode *P1, *P2, *P3, *P4, *P5, *P6;
	pt >> P1 >> P2 >> P3 >> P4;
	P5 = P6 = NULL;
	int i = 1;
	while ( (P1 != NULL) || (P3 != NULL) )
	{
		if ( (P1 != NULL) && (i % 2 != 0) )
			Push(&P1, &P2, &P5, &P6);
		else
			Push(&P3, &P4, &P5, &P6);
		i++;
	}
	P2 = P4 = NULL;
	pt << P5 << P6;
}

void Push(TNode** from1, TNode** from2, TNode** to1, TNode** to2)
{
	TNode* temp = *from1;
	*from1 = (*from1)->Next;
	temp->Next = NULL;
	if (*to2 == NULL)
		*to1 = *to2 = temp;
	else
	{
		(*to2)->Next = temp;
		*to2 = temp;
	}
	temp = NULL;
}