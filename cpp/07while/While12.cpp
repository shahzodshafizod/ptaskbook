#include <iostream>
using namespace std;

int main()
{
	//Task("While12");
	int N;
	cin >> N;
	int K = 1, sum = K;
	while (sum <= N)
	{
		K++;
		sum += K;
	}
	sum -= K;
	K--;
	
	cout << K << sum;
	
	return 0;
}
