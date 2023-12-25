#include <iostream>
using namespace std;

void DigitCountSum(unsigned int K, int& C, int& S);

int main()
{
	//Task("Proc6");
	int K, C, S;
	for (int i = 5; i > 0; i--)
	{
		cin >> K;
		DigitCountSum(K, C, S);
		cout << C << S;
	}
	
	return 0;
}

void DigitCountSum(unsigned int K, int& C, int& S)
{
	C = S = 0;
	while (K > 0)
	{
		S += K % 10;
		C++;
		K /= 10;
	}
}