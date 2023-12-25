#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text25");
	int K;
	cin >> K;
	cin.ignore();
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	
	int sarxats = 0;
	do
		file.getline(str, stringSize);
	while ( (file.peek() != -1) && (strlen(str) == 0) );
	sarxats++;
	
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		if (strlen(str) != 0)
			continue;
		while ( (file.peek() != -1) && (strlen(str) == 0) )
			file.getline(str, stringSize);
		sarxats++;
	}
	
	if (sarxats >= K)
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream tempFile(tempFileName);
		file.clear();
		file.seekg(0);
		sarxats = 0;
		
		do
		{
			file.getline(str, stringSize);
			if (strlen(str) != 0)
				break;
			tempFile << str << endl;
		}
		while (file.peek() != -1);
		sarxats++;
		if (sarxats != K)
			tempFile << str << endl;
		
		while (file.peek() != -1)
		{
			file.getline(str, stringSize);
			if (strlen(str) != 0)
			{
				if (sarxats != K)
					tempFile << str << endl;
				continue;
			}
			tempFile << str << endl;
			
			do
			{
				file.getline(str, stringSize);
				if (strlen(str) != 0)
					break;
				tempFile << str << endl;
			}
			while (file.peek() != -1);
			
			sarxats++;
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
