#include <iostream>
using namespace std;

void AddLeftDigit(short D, int& K);

int main()
{
    //Task("Proc9");
	int K;
	cin >> K;

	short D;
	for (int i = 1; i <= 2; i++)
	{
		cin >> D;
		AddLeftDigit(D, K);
		cout << K;
	}
	
	return 0;
}

void AddLeftDigit(short D, int& K)
{
	int temp = K;
	int factor = 1;
	while (temp > 0)
	{
		temp /= 10;
		factor *= 10;
	}

	K = D * factor + K;
}