#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("Text54");
	const int stringSize = 100;

	char textFileName[100], charFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(charFileName, 100);

	ifstream textFile(textFileName);
	ofstream charFile(charFileName, ios::binary);

	char str[stringSize+1];
	int chars[256] = {0};
	int len, code;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		len = strlen(str);
		for (int i = 0; i < len; i++)
		{
			code = int(str[i]);
			code += code < 0 ? 256 : 0;
			if (chars[code] == 0)
			{
				chars[code]++;
				charFile.write(reinterpret_cast<const char*>(&str[i]), sizeof(str[i]));
			}
		}
	}
	textFile.close();
	charFile.close();

	return 0;
}
