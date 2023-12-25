#include <iostream>
using namespace std;

int main()
{
	//Task("While3");
	int N, K;
	cin >> N >> K;
	int mod, div = 0;
	while (N >= K)
	{
		N -= K;
		div++;
	}
	mod = N;

	cout << div << mod;
	
	return 0;
}
