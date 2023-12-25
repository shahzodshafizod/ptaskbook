#include <iostream>
using namespace std;

void AddRightDigit(short D, int& K);

int main()
{
    //Task("Proc8");
	int K;
	cin >> K;

	short D;
	for (int i = 1; i <= 2; i++)
	{
		cin >> D;
		AddRightDigit(D, K);
		cout << K;
	}
	
	return 0;
}

void AddRightDigit(short D, int& K)
{
	K = K * 10 + D;
}