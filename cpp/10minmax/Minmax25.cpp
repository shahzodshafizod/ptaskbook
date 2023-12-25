#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax25");
	int N;
	cin >> N;

	double prev, curr, zarb, minZarb;
	int index1, index2;
	cin >> prev >> curr;

	minZarb = prev * curr;
	index1 = 1;
	index2 = 2;

	for (int i = 3; i <= N; i++)
	{
		prev = curr;
		cin >> curr;
		zarb = prev * curr;
		if (zarb < minZarb)
		{
			minZarb = zarb;
			index1 = i-1;
			index2 = i;
		}
	}
	cout << index1 << index2;
	
	return 0;
}
