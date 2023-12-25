#include <iostream>
using namespace std;

int main()
{
	//Task("Case3");
	int monthNo;
	cin >> monthNo;
	switch (monthNo)
	{
		case 1: case 2: case 12:
			cout << "зима";
			break;
		case 3: case 4: case 5:
			cout << "весна";
			break;
		case 6: case 7: case 8:
			cout << "лето";
			break;
		case 9: case 10: case 11:
			cout << "осень";
			break;
	}
	
	return 0;
}
