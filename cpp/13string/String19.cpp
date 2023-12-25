#include <iostream>
using namespace std;

int main()
{
	//Task("String19");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	int minus = 0, digits = 0, points = 0;
	if (str[len] == '-')
	{
		len++;
		minus++;
	}
	while (str[len] != '\0')
	{
		if ( (str[len] >= '0') && (str[len] <= '9') )
			digits++;
		else if (str[len] == '.')
			points++;
		len++;
	}
	if ( len == digits+minus )
		cout << 1;
	else if ( (points == 1) && (len == digits+minus+points) )
		cout << 2;
	else
		cout << 0;
	
	return 0;
}
