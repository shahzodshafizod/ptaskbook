#include <iostream>
using namespace std;

int main()
{
	//Task("For36");
	int N, K;
	cin >> N >> K;
	double sum = 0;
	int degree;
	for (int i = 1; i <= N; i++)
	{
		degree = 1;
		for (int j = 1; j <= K; j++)
			degree *= i;
		sum += degree;
	}
	cout << sum;
	
	return 0;
}
