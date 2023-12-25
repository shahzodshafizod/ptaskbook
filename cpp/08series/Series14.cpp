#include <iostream>
using namespace std;

int main()
{
	//Task("Series14");
	int K, number, counts = 0;
	cin >> K;
	do
	{
		cin >> number;
		if (number == 0)
			break;
		if (number < K)
			counts++;
	}
	while (true);
	cout << counts;
	
	return 0;
}
