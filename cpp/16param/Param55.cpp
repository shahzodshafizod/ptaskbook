#include <iostream>
#include <fstream>
using namespace std;

void StringFileToText(const char* S);

int main()
{
	//Task("Param55");
	char fileName[100];
	for (int i = 1; i < 3; i++)
	{
		cin.getline(fileName, 100);
		StringFileToText(fileName);
	}
	
	return 0;
}

void StringFileToText(const char* S)
{
	ifstream ifs(S, ios::binary);
	if (ifs.is_open())
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream ofs(tempFileName);
		const int stringSize = 80;
		char str[stringSize+1];
		while (ifs.peek() != -1)
		{
			ifs.read(reinterpret_cast<char*>(&str), sizeof(char)*stringSize);
			ofs << str << endl;
		}
		ifs.close();
		ofs.close();

		remove(S);
		rename(tempFileName, S);
	}
}