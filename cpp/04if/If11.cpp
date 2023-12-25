#include <iostream>
using namespace std;

int main()
{
	//Task("If11");
	int A, B;
	cin >> A >> B;
	if (A != B)
	{
		if (A > B)
			B = A;
		else
			A = B;
	}
	else
		A = B = 0;
	cout << A << B;
	
	return 0;
}
