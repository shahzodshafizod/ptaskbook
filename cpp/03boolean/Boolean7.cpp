#include <iostream>
using namespace std;

int main()
{
	//Task("Boolean7");
	int A, B, C;
	cin >> A >> B >> C;
	bool progReg = (B > A) && (B < C) || (B > C) && (B < A);
	cout << progReg;
	
	return 0;
}
