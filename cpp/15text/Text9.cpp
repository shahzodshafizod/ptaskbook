#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text9");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	for (int i = 1; true; i++)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		if (i == K)
			tempFile << endl;
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
