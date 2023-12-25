#include <iostream>
using namespace std;

int DigitN(int K, int N);

int main()
{
	//Task("Proc30");
	int K;
	for (int i = 1; i <= 5; i++)
	{
		cin >> K;
		for (int j = 1; j <= 5; j++)
			cout << DigitN(K, j);
	}
	
	return 0;
}

int DigitN(int K, int N)
{
	if (K == 0)
		return (N == 1 ? 0 : -1);
	
	for (int i = 1; K > 0; i++)
	{
		if (i == N)
		{
			return K%10;
		}
		K /= 10;
	}
	return -1;
}