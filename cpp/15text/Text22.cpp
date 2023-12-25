#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text22");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	char** str = new char* [K];
	for (int i = 0; i < K; i++)
	{
		str[i] = new char [stringSize];
		file.getline(str[i], stringSize);
	}

	int index = 0;
	while (file.peek() != -1)
	{
		tempFile << str[index] << endl;
		file.getline(str[index], stringSize);
		index = (index+1) % K;
	}
	file.close();
	tempFile.close();

	for (int i = 0; i < K; i++)
		delete [] str[i];
	delete [] str;
	str = NULL;

	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}
