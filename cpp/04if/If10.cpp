#include <iostream>
using namespace std;

int main()
{
	//Task("If10");
	int A, B;
	cin >> A >> B;
	if (A != B)
	{
		A += B;
		B = A;
	}
	else
		A = B = 0;
	cout << A << B;
	
	return 0;
}
