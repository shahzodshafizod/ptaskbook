#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text15");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	int number = 0;
	while (true)
	{
		number++;
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		if (number == K)
			continue;
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
