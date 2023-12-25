#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text8");
	const int stringSize = 80;
	char fileName1[100], fileName2[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName1, 100);
	cin.getline(fileName2, 100);
	ifstream file1(fileName1);
	ifstream file2(fileName2);
	ofstream tempFile(tempFileName);
	while (true)
	{
		file2.getline(str, stringSize);
		if (file2.eof())
		{
			file2.clear();
			break;
		}
		tempFile << str << endl;
	}
	while (true)
	{
		file1.getline(str, stringSize);
		if (file1.eof())
		{
			file1.clear();
			break;
		}
		tempFile << str << endl;
	}
	file1.close();
	file2.close();
	tempFile.close();
	remove(fileName1);
	rename(tempFileName, fileName1);
	
	return 0;
}
