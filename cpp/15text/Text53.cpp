#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text53");
	const int stringSize = 100;

	char textFileName[100], charFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(charFileName, 100);

	ifstream textFile(textFileName);
	ofstream charFile(charFileName, ios::binary);

	char str[stringSize + 1];
	int len;
	char symbols[] = "!?(),.:;-\"";
	int arraySize = sizeof(symbols) / sizeof(char);
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize + 1);
		len = strlen(str);
		for (int i = 0; i < len; i++)
		{
			for (int j = 0; j < arraySize; j++)
				if (str[i] == symbols[j])
				{
					charFile.write(reinterpret_cast<const char*>(&str[i]), sizeof(str[i]));
					break;
				}
		}
	}
	textFile.close();
	charFile.close();

	return 0;
}
