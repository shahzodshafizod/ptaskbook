#include <iostream>
using namespace std;

int GCD(int A, int B);

int main()
{
	//Task("Recur9");
	int A, B;
	cin >> A;
	for (int i = 1; i < 4; i++)
	{
		cin >> B;
		cout << GCD(A, B);
	}
	return 0;
}

int GCD(int A, int B)
{
	if (B == 0)
		return A;
	
	return GCD(B, A % B);
}