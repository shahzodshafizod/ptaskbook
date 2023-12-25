#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text12");
	const int stringSize = 80;
	char S[stringSize], str[stringSize], fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(S, stringSize);
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		tempFile << ( strlen(str) == 0 ? S : str ) << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
