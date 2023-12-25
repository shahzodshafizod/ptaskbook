#include <iostream>
#include <fstream>
using namespace std;

void Swap(int& a, int& b);
void CtrlJ(char* str, int len, int stringSize);

int main()
{
	//Task("Text37");
	const int stringSize = 50;
	char fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);

	char str[2][stringSize+1];
	int curr = 0, prev = 1;
	bool isFirst = true;
	while (file.peek() != -1)
	{
		if (isFirst)
		{
			file.getline(str[curr], stringSize+1);
			isFirst = false;
			Swap(curr, prev);
			continue;
		}
		
		file.getline(str[curr], stringSize+1);
		if (strlen(str[curr]) == 0)
		{
			tempFile << str[prev] << "\n\n";
			isFirst = true;
			Swap(curr, prev);
			continue;
		}
		
		CtrlJ(str[prev], strlen(str[prev]), stringSize);
		tempFile << str[prev] << endl;
		Swap(curr, prev);
	}
	tempFile << str[prev];

	file.close();
	tempFile.close();

	remove(fileName);
	rename(tempFileName, fileName);

	return 0;
}

void Swap(int& a, int& b)
{
	int temp = a;
	a = b;
	b = temp;
}

void CtrlJ(char* str, int len, int stringSize)
{
	for (int i = len-1; i >= 0; i--)
	{
		if (str[i] == ' ')
		{
			for (int j = len; j > i; j--)
				str[j] = str[j-1];
			len++;
			while ( (i >= 0) && (str[i] == ' ') )
				i--;
		}
		if (len == stringSize)
			break;
		if (i == 0)
			i = len;
	}
	str[len] = '\0';
}