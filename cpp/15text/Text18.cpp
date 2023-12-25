#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text18");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
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
		for (int i = K; i < length; i++)
			tempFile << str[i];
		
		tempFile << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
