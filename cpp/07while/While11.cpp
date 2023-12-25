#include <iostream>
using namespace std;

int main()
{
	//Task("While11");
	int N;
	cin >> N;
	int K = 1, sum = K;
	while (sum < N)
	{
		K++;
		sum += K;
	}
	cout << K << sum;
	
	return 0;
}
