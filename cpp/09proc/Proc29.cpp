#include <iostream>
using namespace std;

int DigitCount(int K);

int main()
{
	//Task("Proc29");
	int K;
	for (int i = 1; i <= 5; i++)
	{
		cin >> K;
		cout << DigitCount(K);
	}
	
	return 0;
}

int DigitCount(int K)
{
	int digits = K == 0 ? 1 : 0;
	while (K > 0)
	{
		digits++;
		K /= 10;
	}
	return digits;
}