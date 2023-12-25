#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text34");
	const int stringSize = 51;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	int strLen;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof()) { file.clear(); break; }
		strLen = strlen(str);
		if (strLen == 0)
		{
			tempFile << endl;
			continue;
		}
		for (int i = stringSize-1-strLen; i > 0; i--)
			tempFile << ' ';
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
