#include <iostream>
using namespace std;

int main()
{
	//Task("Case6");
	int index;
	cin >> index;
	double length, meters = 0;
	cin >> length;
	switch (index)
	{
		case 1:
			meters = length / 10;
			break;
		case 2:
			meters = length * 1000;
			break;
		case 3:
			meters = length;
			break;
		case 4:
			meters = length / 1000;
			break;
		case 5:
			meters = length / 100;
			break;
	}
	cout << meters;
	
	return 0;
}
