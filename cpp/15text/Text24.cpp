#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text24");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	int sarxats = 0;
	do
		file.getline(str, stringSize);
	while ( (file.peek() != -1) && (strlen(str) == 0) );
	sarxats++;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		if (strlen(str) != 0)
			continue;
		while ( (file.peek() != -1) && (strlen(str) == 0) )
			file.getline(str, stringSize);
		
		sarxats++;
	}
	file.close();
	cout << sarxats;
	
	return 0;
}
