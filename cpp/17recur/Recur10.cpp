#include <iostream>
using namespace std;

int DigitSum(int K);

int main()
{
	//Task("Recur10");
	int K;
	for (int i = 0; i < 5; i++)
	{
		cin >> K;
		cout << DigitSum(K);
	}
	return 0;
}

int DigitSum(int K)
{
	if (K < 0)
		K *= -1;
	if (K == 0)
		return 0;
	return K % 10 + DigitSum(K / 10);
}