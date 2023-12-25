#include <iostream>
using namespace std;

int main()
{
	//Task("String22");
	char str[100];
	cin.getline(str, 100);

	int index = 0;
	int sum = 0;
	while (str[index] != '\0')
	{
		sum += int(str[index]) - int('0');
		index++;
	}

	cout << sum;

	return 0;
}
