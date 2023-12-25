#include <iostream>
using namespace std;

bool IsPower5(int K);

int main()
{
	//Task("Proc26");
	int number, degrees = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		degrees += (int)IsPower5(number);
	}
	cout << degrees;
	
	return 0;
}

bool IsPower5(int K)
{
	int degree = 1;
	while (degree < K)
		degree *= 5;
	
	return degree == K;
}