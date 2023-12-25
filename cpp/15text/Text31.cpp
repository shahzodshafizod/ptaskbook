#include <iostream>
#include <fstream>
using namespace std;

bool isKitobat(char);
void makeStr(const char* str, char** newStr, int index1, int index2);
int main()
{
	//Task("Text31");
	const int stringSize = 80;
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100], stringFileName[100], str[stringSize];
	cin.getline(fileName, 100);
	cin.getline(stringFileName, 100);
	ifstream file(fileName);
	ofstream stringFile(stringFileName, ios::binary);
	int length, index1, index2, strLen;
	char* newStr = NULL;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof()) { file.clear(); break; }
		strLen = strlen(str);
		length = index1 = index2 = 0;
		for (int i = 0; i < strLen;)
		{
			while ( (i < strLen) && (str[i] == ' ') )
				i++;
			if (i == strLen)
				break;
			
			length = 0;
			index1 = i;
			while ( (i < strLen) && (str[i] != ' ') && !isKitobat(str[i]) )
			{
				length++;
				i++;
			}
			index2 = i-1;
			
			if (length == K)
			{
				makeStr(str, &newStr, index1, index2);
				stringFile.write(reinterpret_cast<const char*>(newStr), stringSize*sizeof(char));
			}
			if (i == strLen)
				break;
			
			i += (int)isKitobat(str[i]);
		}
	}
	file.close();
	stringFile.close();

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}

void makeStr(const char* str, char** newStr, int index1, int index2)
{
	int len = index2 - index1 + 1;
	delete [] (*newStr);
	(*newStr) = new char [len+1];
	for (int i = index1, j = 0; i <= index2; i++, j++)
		(*newStr)[j] = str[i];
	(*newStr)[len] = '\0';
}

bool isKitobat(char C)
{
	bool result = (C == '.') || (C == ',') || (C == '?') || (C == '!') || (C == ':') ||
				  (C == ';') || (C == '-') || (C == '(') || (C == ')') || (C == '\"');
	return result;
}