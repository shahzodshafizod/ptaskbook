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
	//Task("Text38");
	const int stringSize = 100;

	int K;
	cin >> K;
	cin.ignore();

	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);

	ifstream fin(fileInName);
	ofstream fout(fileOutName);

	char str[stringSize+1];
	int posInFileStr = 0;
	Word word;
	int spaces;
	bool isNewStr = true;
	while (fin.peek() != -1)
	{
		fin.getline(str, stringSize+1);
		
		int len = strlen(str);
		if (len == 0)
		{
			fout << "\n\n";
			posInFileStr = 0;
			continue;
		}

		for (int i = 0; i < len; )
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
			if ( (word.length != 0) && (posInFileStr + word.length + spaces <= K) )
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
				i = word.index1;
				posInFileStr = 0;
				fout << endl;
			}
			if (i == len)
				break;
		}
	}

	fin.close();
	fout.close();

	return 0;
}
