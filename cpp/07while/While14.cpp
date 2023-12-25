#include <iostream>
using namespace std;

int main()
{
	//Task("While14");
	double A;
	cin >> A;
	int K = 1;
	double sum = 1.0 / K;
	while (sum < A)
	{
		K++;
		sum += 1.0 / K;
	}
	sum -= 1.0 / K;
	K--;

	cout << K << sum;
	
	return 0;
}
