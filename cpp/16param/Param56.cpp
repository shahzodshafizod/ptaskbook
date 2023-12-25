#include <iostream>
#include <fstream>
using namespace std;

void TextToStringFile(const char* S);

int main()
{
	//Task("Param56");
	char fileName[100];
	for (int i = 0; i < 2; i++)
	{
		cin.getline(fileName, 100);
		TextToStringFile(fileName);
	}
	
	return 0;
}

void TextToStringFile(const char* S)
{
	ifstream ifs(S);
	if ( ifs.is_open() )
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream ofs(tempFileName, ios::binary);
		const int stringSize = 80;
		char str[stringSize+1];
		while (ifs.peek() != -1)
		{
			ifs.getline(str, stringSize+1);
			ofs.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
		}
		ifs.close();
		ofs.close();

		remove(S);
		rename(tempFileName, S);
	}
}