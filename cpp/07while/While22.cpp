#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("While22");
	int N;
	cin >> N;
	bool isPrime = true;
	int i = sqrt((double)N);
	while (i > 1)
	{
		if (N % i == 0)
			isPrime = false;
		i--;
	}
	cout << isPrime;
	
	return 0;
}
