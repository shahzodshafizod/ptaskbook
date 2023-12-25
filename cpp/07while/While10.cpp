#include <iostream>
using namespace std;

int main()
{
	//Task("While10");
	int N;
	cin >> N;
	int degree3 = 3, K = 1;
	while (degree3 < N)
	{
		degree3 *= 3;
		K++;
	}
	--K;
	
	cout << K;
	
	return 0;
}
