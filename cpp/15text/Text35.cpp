#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text35");
	const int stringSize = 51;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	int strLen;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		strLen = strlen(str);
		if (strLen == 0)
		{
			tempFile << endl;
			continue;
		}
		for (int i = (stringSize-1-strLen)/2+abs(strLen%2); i > 0; i--)
			tempFile << ' ';
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();

	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
