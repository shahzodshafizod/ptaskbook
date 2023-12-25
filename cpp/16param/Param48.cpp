#include <iostream>
#include <fstream>
using namespace std;

int IntFileSize(const char* S);

int main()
{
	//Task("Param48");
	char fileName[100];
	for (int i = 1; i <= 3; i++)
	{
		cin.getline(fileName, 100);
		cout << IntFileSize(fileName);
	}
	
	return 0;
}

int IntFileSize(const char* S)
{
	int fileSize = -1;
	ifstream file(S, ios::binary);
	if (file.is_open())
	{
		file.seekg(0, ios::end);
		fileSize = file.tellg() / sizeof(int);
		file.close();
	}

	return fileSize;
}