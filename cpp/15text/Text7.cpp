#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text7");
	const int stringSize = 80;
	char fileName[100], S[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(S, stringSize);
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	tempFile << S;
	while (true)
	{
		file.getline(S, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		tempFile << endl << S;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
