#include <iostream>
using namespace std;

int main()
{
	//Task("For12");
	int N;
	cin >> N;
	double elem = 1.1, zarb = 1;
	for (int i = 1; i <= N; i++)
	{
		zarb *= elem;
		elem += 0.1;
	}
	cout << zarb;
	
	return 0;
}
