#include <iostream>
using namespace std;

int main()
{
	//Task("For10");
	int N;
	cin >> N;
	double sum = 0;
	for (int i = 1; i <= N; i++)
		sum += 1.0 / i;
	cout << sum;
	
	return 0;
}
