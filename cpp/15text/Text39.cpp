#include <iostream>
#include <fstream>
using namespace std;

struct Word
{
	int index1;
	int index2;
	int length;
};

int main()
{
	//Task("Text39");
	const int stringSize = 100;

	int K;
	cin >> K;
	cin.ignore();

	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);

	ifstream fin(fileInName);
	ofstream fout(fileOutName);

	Word word;
	int spaces, posInFileStr = 0;
	char str[stringSize+1];
	while (true)
	{
		fin.getline(str, stringSize+1);
		if (fin.eof())
			break;
		
		int len = strlen(str);
		bool isSarxat = true;
		int i = 0;
		for (; i < 5; i++)
			if (str[i] != ' ')
			{
				isSarxat = false;
				break;
			}
		
		i = 0;
		if (isSarxat)
		{
			if (posInFileStr != 0)
			{
				posInFileStr = 0;
				fout << endl;
			}
			for ( ; i < 5; i++)
			{
				fout << ' ';
				posInFileStr++;
			}
		}
		while (i < len)
		{
			spaces = 0;
			while ( (i < len) && (str[i] == ' ') )
			{
				i++;
				spaces++;
			}

			word.index1 = i;
			word.length = 0;
			while ( (i < len) && (str[i] != ' ') )
			{
				i++;
				word.length++;
			}
			word.index2 = i-1;

			if (word.index1 == 0)
				spaces++;

			if ( (word.length != 0) && (posInFileStr + spaces + word.length <= K) )
			{
				if (posInFileStr != 0)
					for (int j = 0; j < spaces; j++)
					{
						fout << ' ';
						posInFileStr++;
					}
				for (int j = word.index1; j <= word.index2; j++)
				{
					fout << str[j];
					posInFileStr++;
				}
			}
			else
			{
				posInFileStr = 0;
				fout << endl;
				i = word.index1;
			}
			if (i == len)
				break;
		}
	}

	fin.close();
	fout.close();
	
	return 0;
}
