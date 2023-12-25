#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text20");
	const int stringSize = 80;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	int length;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		length = strlen(str);
		for (int i = 0; i < length;)
		{
			if (str[i] == ' ')
			{
				tempFile << str[i];
				while ( (i < length) && (str[i] == ' ') )
					i++;
				continue;
			}
			tempFile << str[i];
			i++;
		}
		tempFile << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
