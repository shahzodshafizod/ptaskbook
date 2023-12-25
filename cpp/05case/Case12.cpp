#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Case12");
	const double PI = 3.14;
	int index;
	cin >> index;
	double value, R, D, L, S;
	cin >> value;
	switch (index)
	{
		case 1:
			R = value;
			D = 2 * R;
			L = 2 * PI * R;
			S = PI * R * R;
			break;
		case 2:
			D = value;
			R = D / 2;
			L = 2 * PI * R;
			S = PI * R * R;
			break;
		case 3:
			L = value;
			R = L / (2 * PI);
			D = 2 * R;
			S = PI * R * R;
			break;
		case 4:
			S = value;
			R = sqrt(S / PI);
			D = 2 * R;
			L = 2 * PI * R;
			break;
	}
	
	if (index != 1)
		cout << R;

	if (index != 2)
		cout << D;

	if (index != 3)
		cout << L;

	if (index != 4)
		cout << S;
	
	return 0;
}
