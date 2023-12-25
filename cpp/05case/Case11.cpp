#include <iostream>
using namespace std;

int main()
{
	//Task("Case11");
	char C;
	cin >> C;
	int N1, N2;
	cin >> N1 >> N2;
	int command = N1 + N2;
	command *= command < 0 ? -1 : 1;
	switch (C)
	{
		case 'С':
			switch ( command )
			{
				case 1: C = 'З'; break;
				case 2: C = 'Ю'; break;
				case 3: C = 'В'; break;
			}
			break;
		case 'В':
			switch ( command )
			{
				case 1: C = 'С'; break;
				case 2: C = 'З'; break;
				case 3: C = 'Ю'; break;
			}
			break;
		case 'Ю':
			switch ( command )
			{
				case 1: C = 'В'; break;
				case 2: C = 'С'; break;
				case 3: C = 'З'; break;
			}
			break;
		case 'З':
			switch ( command )
			{
				case 1: C = 'Ю'; break;
				case 2: C = 'В'; break;
				case 3: C = 'С'; break;
			}
			break;
	}
	cout << C;
	
	return 0;
}
