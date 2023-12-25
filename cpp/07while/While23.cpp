#include <iostream>
using namespace std;

int main()
{
	//Task("While23");
	int A, B;
	cin >> A >> B;
	int temp;
	while (B != 0)
	{
		temp = B;
		B = A % B;
		A = temp;
	}
	cout << A;
	
	return 0;
}
