#include <iostream>
using namespace std;

int main()
{
	//Task("Case7");
	int index;
	cin >> index;
	double weight, kg = 0;
	cin >> weight;
	switch (index)
	{
		case 1:
			kg = weight;
			break;
		case 2:
			kg = weight / 1000000;
			break;
		case 3:
			kg = weight / 1000;
			break;
		case 4:
			kg = weight * 1000;
			break;
		case 5:
			kg = weight * 100;
	}
	cout << kg;
	
	return 0;
}
