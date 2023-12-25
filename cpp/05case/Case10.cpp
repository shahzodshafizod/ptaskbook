#include <iostream>
using namespace std;

int main()
{
	//Task("Case10");
	char C;
	cin >> C;
	int N;
	cin >> N;
	switch (C)
	{
		case 'С':
			if (N == 1)
				C = 'З';
			else if (N == -1)
				C = 'В';
			break;
		case 'В':
			if (N == 1)
				C = 'С';
			else if (N == -1)
				C = 'Ю';
			break;
		case 'Ю':
			if (N == 1)
				C = 'В';
			else if (N == -1)
				C = 'З';
			break;
		case 'З':
			if (N == 1)
				C = 'Ю';
			else if (N == -1)
				C = 'С';
			break;
	}
	cout << C;
	
	return 0;
}
