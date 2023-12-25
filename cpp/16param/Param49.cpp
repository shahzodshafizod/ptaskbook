#include <iostream>
#include <fstream>
using namespace std;

int LineCount(const char* S);

int main()
{
	//Task("Param49");
	char fileName[100];
	for (int i = 0; i < 3; i++)
	{
		cin.getline(fileName, 100);
		cout << LineCount(fileName);
	}
	
	return 0;
}

int LineCount(const char* S)
{
	int lines = -1;
	ifstream file(S);
	if (file.is_open())
	{
		const int stringSize = 100;
		char str[stringSize+1];
		lines = 0;
		while (file.peek() != -1)
		{
			file.getline(str, stringSize);
			lines++;
		}
	}
	return lines;
}