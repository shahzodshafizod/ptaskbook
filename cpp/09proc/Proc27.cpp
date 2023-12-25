#include <iostream>
using namespace std;

bool IsPowerN(int K, int N);

int main()
{
	//Task("Proc27");
	int N;
	cin >> N;

	int K, powers = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> K;
		powers += IsPowerN(K, N);
	}
	cout << powers;
	
	return 0;
}

bool IsPowerN(int K, int N)
{
	int degree = 1;
	while (degree < K)
		degree *= N;

	return (degree == K);
}