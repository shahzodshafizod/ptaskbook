#include <iostream>
using namespace std;

int main()
{
	//Task("Case9");
	int D, M;
	cin >> D >> M;
	switch (M)
	{
		case 1: case 3: case 5: case 7: case 8: case 10:
			if (D == 31)
			{
				D = 1;
				M++;
			}
			else
				D++;
			break;
		case 4: case 6: case 9: case 11:
			if (D == 30)
			{
				D = 1;
				M++;
			}
			else
				D++;
			break;
		case 2:
			if (D == 28)
			{
				D = 1;
				M++;
			}
			else
				D++;
			break;
		case 12:
			if (D == 31)
				D = M = 1;
			else
				D++;
			break;
	}
	cout << D << M;
	
	return 0;
}
