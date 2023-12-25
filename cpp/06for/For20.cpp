#include <iostream>
using namespace std;

int main()
{
	//Task("For20");
	int N;
	cin >> N;
	double fact = 1, sum = 0;
	for (int i = 1; i <= N; i++)
	{
		fact *= i;
		sum += fact;
	}
	cout << sum;
	
	return 0;
}
