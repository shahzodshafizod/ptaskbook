#include <iostream>
using namespace std;

int main()
{
	//Task("While9");
	int N;
	cin >> N;
	int degree3 = 3, K = 1;
	while (degree3 <= N)
	{
		degree3 *= 3;
		K++;
	}
	cout << K;
	
	return 0;
}
