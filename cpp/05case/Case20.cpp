#include <iostream>
using namespace std;

int main()
{
	//Task("Case20");
	int D, M;
	cin >> D >> M;
	switch (M)
	{
		case 1:
			if (D <= 19)
				cout << "Козерог";
			else
				cout << "Водолей";
			break;
		case 2:
			if (D <= 18)
				cout << "Водолей";
			else
				cout << "Рыбы";
			break;
		case 3:
			if (D <= 20)
				cout << "Рыбы";
			else
				cout << "Овен";
			break;
		case 4:
			if (D <= 19)
				cout << "Овен";
			else
				cout << "Телец";
			break;
		case 5:
			if (D <= 20)
				cout << "Телец";
			else
				cout << "Близнецы";
			break;
		case 6:
			if (D <= 21)
				cout << "Близнецы";
			else
				cout << "Рак";
			break;
		case 7:
			if (D <= 22)
				cout << "Рак";
			else
				cout << "Лев";
			break;
		case 8:
			if (D <= 22)
				cout << "Лев";
			else
				cout << "Дева";
			break;
		case 9:
			if (D <= 22)
				cout << "Дева";
			else
				cout << "Весы";
			break;
		case 10:
			if (D <= 22)
				cout << "Весы";
			else
				cout << "Скорпион";
			break;
		case 11:
			if (D <= 22)
				cout << "Скорпион";
			else
				cout << "Стрелец";
			break;
		case 12:
			if (D <= 21)
				cout << "Стрелец";
			else
				cout << "Козерог";
			break;
	}
	
	return 0;
}
