#include "TNode.h"
#include <iostream>
using namespace std;

int main() {
	TNode* P1 = new TNode();
	cout << "P1->Data = ";
	cin >> P1->Data;
	
	P1->Left = new TNode();
	cout << "P1->Left->Data = ";
	cin >> P1->Left->Data;
	
	P1->Right = new TNode();
	cout << "P1->Right->Data = ";
	cin >> P1->Right->Data;
	
	cout << "P1->Data = " << P1->Data << endl;
	cout << "P1->Left->Data = " << P1->Left->Data << endl;
	cout << "P1->Right->Data = " << P1->Right->Data << endl;
	cout << "P1->Left = " << P1->Left << endl;
	cout << "P1->Right = " << P1->Right << endl;
	
	delete P1->Left;
	delete P1->Right;
	delete P1;
	P1 = NULL;
	return 0;
}