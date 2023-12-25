#include <iostream>
#include <cmath>
using namespace std;

bool IsPrime(int N);

int main()
{
	//Task("Proc28");
	int number, primes = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		primes += (int)IsPrime(number);
	}
	cout << primes;
	
	return 0;
}

bool IsPrime(int N)
{
	for (int i = sqrt((double)N); i > 1; i--)
		if (N % i == 0)
			return false;

	return true;
}