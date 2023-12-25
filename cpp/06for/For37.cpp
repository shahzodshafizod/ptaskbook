#include <iostream>
using namespace std;

int main()
{
	//Task("For37");
	int N;
	cin >> N;
	double sum = 0;
	int degree;
	for (int i = 1; i <= N; i++)
	{
		degree = 1;
		for (int j = 1; j <= i; j++)
			degree *= i;
		sum += degree;
	}
	cout << sum;
	
	return 0;
}
