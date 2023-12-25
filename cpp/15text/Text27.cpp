#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text27");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	int sarxats = 0;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		if (strlen(str) >= 5)
		{
			bool isSarxat = true;
			for (int i = 0; i < 5; i++)
				if (str[i] != ' ')
				{
					isSarxat = false;
					break;
				}
			sarxats += (int)isSarxat;
		}
	}
	if (sarxats >= K)
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream tempFile(tempFileName);
		sarxats = 0;
		int length;
		file.seekg(0);
		while (file.peek() != -1)
		{
			file.getline(str, stringSize);
			length = strlen(str);
			if (length == 0)
			{
				tempFile << str << endl;
				continue;
			}

			if (length >= 5)
			{
				bool isSarxat = true;
				for (int i = 0; i < 5; i++)
					if (str[i] != ' ')
					{
						isSarxat = false;
						break;
					}
				sarxats += (int)isSarxat;
			}
			
			if (sarxats != K)
				tempFile << str << endl;
		}
		file.close();
		tempFile.close();
		remove(fileName);
		rename(tempFileName, fileName);
	}
	else
		file.close();
	
	return 0;
}
