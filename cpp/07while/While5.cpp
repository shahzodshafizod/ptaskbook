#include <iostream>
using namespace std;

int main()
{
	//Task("While5");
	int N;
	cin >> N;
	int K = 0, degree2 = 1;
	while (degree2 < N)
	{
		degree2 *= 2;
		K++;
	}
	cout << K;
	
	return 0;
}
