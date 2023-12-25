#include <iostream>
using namespace std;

int main()
{
	//Task("Case8");
	int D, M;
	cin >> D >> M;
	switch (M)
	{
		case 1:
			if (D == 1)
			{
				D = 31;
				M = 12;
			}
			else
				D--;
			break;
		case 3:
			if (D == 1)
			{
				D = 28;
				M--;
			}
			else
				D--;
			break;
		case 2: case 4: case 6: case 8: case 9: case 11:
			if (D == 1)
			{
				D = 31;
				M--;
			}
			else
				D--;
			break;
		case 5: case 7: case 10: case 12:
			if (D == 1)
			{
				D = 30;
				M--;
			}
			else
				D--;
			break;
	}
	cout << D << M;
	
	return 0;
}
