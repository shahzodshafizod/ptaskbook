#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text28");
	const int stringSize = 80;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	file.getline(str, stringSize);
	tempFile << str << endl;
	bool isSarxat;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		if (strlen(str) >= 5)
		{
			isSarxat = true;
			for (int i = 0; i < 5; i++)
				if (str[i] != ' ')
				{
					isSarxat = false;
					break;
				}
			if (isSarxat)
				tempFile << endl;
		}
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
