#include <iostream>
using namespace std;

int main()
{
	//Task("If10");
	int A, B;
	cin >> A >> B;
	A = B = A != B ? A+B : 0;
	cout << A << B;
	
	return 0;
}
