#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text49");
	const int stringSize = 100;

	char textFileName[100], intFileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(textFileName, 100);
	cin.getline(intFileName, 100);

	ifstream textFile(textFileName);
	ifstream intFile(intFileName, ios::binary);
	ofstream tempFile(tempFileName);

	char str[stringSize+1];
	int number;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		tempFile << str;
		if (intFile.peek() != -1)
		{
			intFile.read(reinterpret_cast<char*>(&number), sizeof(number));
			tempFile << number;
		}
		tempFile << endl;
	}
	textFile.close();
	intFile.close();
	tempFile.close();

	remove(textFileName);
	rename(tempFileName, textFileName);

	return 0;
}
