#include <iostream>
using namespace std;

int main()
{
	//Task("String6");
	char C;
	cin >> C;

	if ( (C >= '0') && (C <= '9') )
		cout << "digit";

	else if ( (C >= 'A') && (C <= 'Z') || (C >= 'a') && (C <= 'z') )
		cout << "lat";

	else if ( (C >= 'А') && (C <= 'Я') || (C >= 'а') && (C <= 'я') || (C == 'Ё') || (C == 'ё') )
		cout << "rus";

	return 0;
}
