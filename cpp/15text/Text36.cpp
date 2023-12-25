#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text36");
	const int stringSize = 51;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	int strLen, spaces;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		strLen = strlen(str);
		if (strLen == 0)
		{
			tempFile << endl;
			continue;
		}
		spaces = 0;
		for (int i = 0; (i < strLen) && (str[i] == ' '); i++)
			spaces++;
		for (int i = spaces/2+spaces%2; i < strLen; i++)
			tempFile << str[i];
		tempFile << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
