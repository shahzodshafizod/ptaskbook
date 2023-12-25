#include <iostream>
using namespace std;

int main()
{
	//Task("For38");
	int N;
	cin >> N;
	double sum = 0;
	int degree;
	for (int i = 1; i <= N; i++)
	{
		degree = 1;
		for (int j = N-i+1; j > 0; j--)
			degree *= i;
		sum += degree;
	}
	cout << sum;
	
	return 0;
}
