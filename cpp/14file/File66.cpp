#include <iostream>
#include <fstream>
using namespace std;

int strCompare(const char* str1, const char* str2);

int main()
{
	//Task("File66");
	char fileInName[100], fileOutName[100];
	
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);
	ifstream fileIn(fileInName, ios::binary);
	ofstream fileOut(fileOutName, ios::binary);
	
	const int stringSize = 80;
	fileIn.seekg(0, ios::end);
	int strings = fileIn.tellg() / (sizeof(char)*stringSize);
	fileIn.seekg(0);
	char** str = new char* [strings];
	for (int i = 0; i < strings; i++)
	{
		str[i] = new char [stringSize];
		fileIn.read(reinterpret_cast<char*>(str[i]), sizeof(char)*stringSize);
	}
	
	char* strPtr = NULL;
	for (int i = 0; i < strings-1; i++)
		for (int j = 1; j < strings-i; j++)
			if (strCompare(str[j-1], str[j]) == 1)
			{
				strPtr = str[j];
				str[j] = str[j-1];
				str[j-1] = strPtr;
				strPtr = NULL;
			}
	
	for (int i = 0; i < strings; i++)
		fileOut.write(reinterpret_cast<const char*>(str[i]), sizeof(char)*stringSize);
		
	for (int i = 0; i < strings; i++)
		delete [] str[i];
	delete [] str;
	str = NULL;
	fileIn.close();
	fileOut.close();
	
	return 0;
}

int strCompare(const char* str1, const char* str2)
{
	int length1 = strlen(str1);
	int length2 = strlen(str2);
	for (int i = 0; (i < length1) && (i < length2); i++)
	{
		if (str1[i] < str2[i])
			return -1;
		if (str1[i] > str2[i])
			return 1;
	}
	
	return length1 - length2;
}