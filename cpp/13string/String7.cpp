#include <iostream>
using namespace std;

int main()
{
	//Task("String7");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	cout << int(str[0]) << int(str[len-1]);
	
	return 0;
}
