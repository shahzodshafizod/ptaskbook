#include <iostream>
using namespace std;

int main()
{
	//Task("Series16");
	int K, number, index = 0, i = 0;
	cin >> K;
	do
	{
		i++;
		cin >> number;
		if (number == 0)
			break;
		if (number > K)
			index = i;
	}
	while (true);
	cout << index;
	
	return 0;
}
