#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text4");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	int strings = 0, chars = 0;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		strings++;
		chars += strlen(str);
	}
	cout << chars << strings;
	
	return 0;
}