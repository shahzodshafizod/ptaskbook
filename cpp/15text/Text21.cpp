#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text21");
	const int stringSize = 80;
	char fileName[100], str[3][stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	for (int i = 0; i < 3; i++)
		file.getline(str[i], stringSize);
	int index = 0;
	while (file.peek() != -1)
	{
		tempFile << str[index] << endl;
		file.getline(str[index], stringSize);
		index = (index + 1) % 3;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);

	return 0;
}
