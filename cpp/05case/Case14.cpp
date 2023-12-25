#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Case14");
	int index;
	cin >> index;
	double value, a, R1, R2, S;
	cin >> value;
	switch (index)
	{
		case 1:
			a = value;
			R1 = a * sqrt(3.0) / 6;
			R2 = 2 * R1;
			S = a * a * sqrt(3.0) / 4;
			break;
		case 2:
			R1 = value;
			R2 = 2 * R1;
			a = 6 * R1 / sqrt(3.0);
			S = a * a * sqrt(3.0) / 4;
			break;
		case 3:
			R2 = value;
			R1 = R2 / 2;
			a = 6 * R1 / sqrt(3.0);
			S = a * a * sqrt(3.0) / 4;
			break;
		case 4:
			S = value;
			a = sqrt( 4 * S / sqrt(3.0) );
			R1 = a * sqrt(3.0) / 6;
			R2 = 2 * R1;
			break;
	}

	if (index != 1)
		cout << a;

	if (index != 2)
		cout << R1;

	if (index != 3)
		cout << R2;

	if (index != 4)
		cout << S;
	
	return 0;
}
