#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text14");
	const int stringSize = 80;
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	
	char str1[stringSize], str2[stringSize];
	bool isStr1Full = false, isStr2Full = false;
	file.getline(str1, stringSize);
	isStr1Full = true;
	while (true)
	{
		if (isStr1Full)
		{
			file.getline(str2, stringSize);
			if (file.eof())
			{
				file.clear();
				break;
			}
			isStr2Full = true;
			tempFile << str1 << endl;
			isStr1Full = false;
		}
		else
		{
			file.getline(str1, stringSize);
			if (file.eof())
			{
				file.clear();
				break;
			}
			isStr1Full = true;
			tempFile << str2 << endl;
			isStr2Full = false;
		}
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
