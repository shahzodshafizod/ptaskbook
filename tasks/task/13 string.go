package task

type sstring struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewStringTask() Task {
	s := &sstring{
		count: 71,
		name:  "String",
	}
	s.makeEn()
	s.makeTj()
	s.makeRu()
	return s
}

func (s *sstring) Count() int   { return s.count }
func (s *sstring) Name() string { return s.name }
func (s *sstring) En() []string { return s.en }
func (s *sstring) Tj() []string { return s.tj }
func (s *sstring) Ru() []string { return s.ru }

func (s *sstring) makeEn() {
	s.en = make([]string, s.count)
	s.en[1] = `Given a character&nbsp;<i>C</i>, output its numeric value in the character set.`
	s.en[2] = `Given an integer&nbsp;<i>N</i> (32&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;126), output a character with the numeric value&nbsp;<i>N</i> in the character set.`
	s.en[3] = `Given a character&nbsp;<i>C</i>, output two characters: the first character precedes&nbsp;<i>C</i> in the character set, the second one follows&nbsp;<i>C</i> in the character set.`
	s.en[4] = `Given an integer&nbsp;<i>N</i> (1&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;26), output <i>N</i>&nbsp;first <i>capital</i> (that is, uppercase) letters of the English alphabet (&#34;A&#34;, &#34;B&#34;, &#34;C&#34;, and so on).`
	s.en[5] = `Given an integer&nbsp;<i>N</i> (1&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;26), output <i>N</i>&nbsp;last <i>small</i> (that is, lowercase) letters of the English alphabet in inverse order (&#34;z&#34;, &#34;y&#34;, &#34;x&#34;, and so on).`
	s.en[6] = `A character&nbsp;<i>C</i> representing a digit or a letter of the Latin alphabet is given. If <i>C</i>&nbsp;is a digit then output the string &#34;digit&#34;, if <i>C</i>&nbsp;is a capital letter then output the string &#34;capital&#34;, otherwise output the string &#34;small&#34;.`
	s.en[7] = `Given a nonempty string, output numeric values of its first and last character in the character set.`
	s.en[8] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a character&nbsp;<i>C</i>, output a string that is of length&nbsp;<i>N</i> and contains characters&nbsp;<i>C</i>.`
	s.en[9] = `Given an even integer&nbsp;<i>N</i> (&gt;&nbsp;0) and two characters&nbsp;<i>C</i><sub>1</sub>, <i>C</i><sub>2</sub>, output a string that is of length&nbsp;<i>N</i>, begins with&nbsp;<i>C</i><sub>1</sub>, and contains alternating characters&nbsp;<i>C</i><sub>1</sub> and&nbsp;<i>C</i><sub>2</sub>.`
	s.en[10] = `Given a string, output a new string that contains  the given string characters in inverse order.`
	s.en[11] = `Given a nonempty string, output a new string that contains the given string characters separated by a blank character.`
	s.en[12] = `Given a nonempty string and an integer&nbsp;<i>N</i> (&gt;&nbsp;0), output a new string that contains the given string characters separated by <i>N</i>&nbsp;characters &#34;*&#34;.`
	s.en[13] = `Given a string, find the amount of digits in the string.`
	s.en[14] = `Given a string, find the amount of Latin capital letters in the string.`
	s.en[15] = `Given a string, find the amount of Latin letters in the string.`
	s.en[16] = `Given a string, convert all Latin capital letters of the string to lowercase.`
	s.en[17] = `Given a string, convert all Latin small letters of the string to uppercase.`
	s.en[18] = `Given a string, convert all Latin capital letters of the string to lowercase  and all Latin small letters of the string to uppercase.`
	s.en[19] = `A string is given. If the string represents an integer then output&nbsp;1, if the string represents a real number (with nonzero fractional part) then output&nbsp;2, otherwise output&nbsp;0. A fractional part of a real number is preceded by the <i>decimal point</i> &#34;.&#34;.`
	s.en[20] = `Given a positive integer, output all digit characters in the decimal representation of the integer (from left to right).`
	s.en[21] = `Given a positive integer, output all digit characters in the decimal representation of the integer (from right to left).`
	s.en[22] = `Given a string that represents a positive integer, output the sum of digits of this integer.`
	s.en[23] = `Given a string that represents an arithmetic expression of the form &#34;&lt;digit&gt;&#177;&lt; digit&gt;&#177;&#8230;&#177;&lt;digit&gt;&#34; with operators &#34;+&#34; and &#34;&#8722;&#34; only (for example, &#34;4+7&#8722;2&#8722;8&#34;), output the value of given expression as an integer.`
	s.en[24] = `Given a string with the binary representation of a positive integer, output a new string with the decimal representation of this integer.`
	s.en[25] = `Given a string with the decimal representation of a positive integer, output a new string with the binary representation of this integer.`
	s.en[26] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;0) and a string&nbsp;<i>S</i> are given. Transform the string&nbsp;<i>S</i> to a new string of length&nbsp;<i>N</i> as follows: if the length of&nbsp;<i>S</i> is greater than&nbsp;<i>N</i> then remove its first characters, if the length of&nbsp;<i>S</i> is less than&nbsp;<i>N</i> then add characters &#34;.&#34; to the beginning of&nbsp;<i>S</i>.`
	s.en[27] = `Two positive integers&nbsp;<i>N</i><sub>1</sub>, <i>N</i><sub>2</sub> and two strings&nbsp;<i>S</i><sub>1</sub>, <i>S</i><sub>2</sub> are given. Output new string that contains <i>N</i><sub>1</sub>&nbsp;first characters of the string&nbsp;<i>S</i><sub>1</sub> and <i>N</i><sub>2</sub>&nbsp;last characters of the string&nbsp;<i>S</i><sub>2</sub> (in that order).`
	s.en[28] = `Given a character&nbsp;<i>C</i> and a string&nbsp;<i>S</i>, double each occurrence of the character&nbsp;<i>C</i> in the string&nbsp;<i>S</i>.`
	s.en[29] = `Given a character&nbsp;<i>C</i> and two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub>, insert the string&nbsp;<i>S</i><sub>0</sub> into the string&nbsp;<i>S</i> before each occurrence of the character&nbsp;<i>C</i>.`
	s.en[30] = `Given a character&nbsp;<i>C</i> and two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub>, insert the string&nbsp;<i>S</i><sub>0</sub> into the string&nbsp;<i>S</i> after each occurrence of the character&nbsp;<i>C</i>.`
	s.en[31] = `Two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub> are given. If the string&nbsp;<i>S</i><sub>0</sub> is a substring of&nbsp;<i>S</i> then output true, otherwise output false.`
	s.en[32] = `Two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub> are given. Find the amount of occurrences of&nbsp;<i>S</i><sub>0</sub> in the string&nbsp;<i>S</i>.`
	s.en[33] = `Two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub> are given. Remove the first occurrence of&nbsp;<i>S</i><sub>0</sub> from the string&nbsp;<i>S</i>. If the string&nbsp;<i>S</i> does not contain required substrings then do not change it.`
	s.en[34] = `Two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub> are given. Remove the last occurrence of&nbsp;<i>S</i><sub>0</sub> from the string&nbsp;<i>S</i>. If the string&nbsp;<i>S</i> does not contain required substrings then do not change it.`
	s.en[35] = `Two strings&nbsp;<i>S</i>, <i>S</i><sub>0</sub> are given. Remove all occurrences of&nbsp;<i>S</i><sub>0</sub> from the string&nbsp;<i>S</i>. If the string&nbsp;<i>S</i> does not contain required substrings then do not change it.`
	s.en[36] = `Three strings&nbsp;<i>S</i>, <i>S</i><sub>1</sub>, <i>S</i><sub>2</sub> are given. Replace the first occurrence of&nbsp;<i>S</i><sub>1</sub> in the string&nbsp;<i>S</i> by the string&nbsp;<i>S</i><sub>2</sub>.`
	s.en[37] = `Three strings&nbsp;<i>S</i>, <i>S</i><sub>1</sub>, <i>S</i><sub>2</sub> are given. Replace the last occurrence of&nbsp;<i>S</i><sub>1</sub> in the string&nbsp;<i>S</i> by the string&nbsp;<i>S</i><sub>2</sub>.`
	s.en[38] = `Three strings&nbsp;<i>S</i>, <i>S</i><sub>1</sub>, <i>S</i><sub>2</sub> are given. Replace all occurrences of&nbsp;<i>S</i><sub>1</sub> in the string&nbsp;<i>S</i> by the string&nbsp;<i>S</i><sub>2</sub>.`
	s.en[39] = `A string with at least one blank character is given. Output the substring of&nbsp;<i>S</i> that contains all characters between the first and the second blank character. If the string&nbsp;<i>S</i> contains only one blank character then output an empty string.`
	s.en[40] = `A string with at least one blank character is given. Output the substring of&nbsp;<i>S</i> that contains all characters between the first and the last blank character. If the string&nbsp;<i>S</i> contains only one blank character then output an empty string.`
	s.en[41] = `A string that contains English words separated by one or more blank characters is given. Find the amount of words in the string.`
	s.en[42] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Find the amount of words whose first letter is coincides with the last one.`
	s.en[43] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Find the amount of words containing at least one letter &#34;E&#34;.`
	s.en[44] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Find the amount of words containing exactly three letters &#34;E&#34;.`
	s.en[45] = `A string that contains English words separated by one or more blank characters is given. Find the length of the shortest word.`
	s.en[46] = `A string that contains English words separated by one or more blank characters is given. Find the length of the longest word.`
	s.en[47] = `A string that contains English words separated by one or more blank characters is given. Output a new string that contains the given words (in the same order) separated by one character &#34;.&#34;.`
	s.en[48] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Process each word as follows: replace all next occurrences of its first letter by the character &#34;.&#34; (for example, the word &#34;MINIMUM&#34; must be transformed into &#34;MINI.U.&#34;). Do not change blank characters in the string.`
	s.en[49] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Process each word as follows: replace all previous occurrences of its last letter by the character &#34;.&#34; (for example, the word &#34;MINIMUM&#34; must be transformed into &#34;.INI.UM&#34;). Do not change blank characters in the string.`
	s.en[50] = `A string that contains English words separated by one or more blank characters is given. Output a new string that contains the given words in inverse order. The words must be separated by one blank character.`
	s.en[51] = `A string that contains English words separated by one or more blank characters is given. All string letters are in uppercase. Output a new string that contains the given words in alphabetic order. The words must be separated by one blank character.`
	s.en[52] = `A string with an English sentence is given. Convert the first letter of each word to uppercase. A <i>word</i> is defined as a character sequence that does not contain blank characters and is bounded by blank characters or the string beginning/end. If the first word character is not a letter then do not change this word.`
	s.en[53] = `A string with an English sentence is given. Find the amount of punctuation marks in the string.`
	s.en[54] = `A string with an English sentence is given. Find the amount of vowels (&#34;a&#34;, &#34;i&#34;, &#34;e&#34;, &#34;o&#34;, &#34;u&#34;) in the string.`
	s.en[55] = `A string with an English sentence is given. Output the longest word in the string. If there are several words of the maximal length then output the first one. A <i>word</i> is defined as a character sequence that does not contain blank characters, punctuation marks and is bounded by blank characters, punctuation marks or the string beginning/end.`
	s.en[56] = `A string with an English sentence is given. Output the shortest word in the string. If there are several words of the maximal length then output the last one. A <i>word</i> is defined as a character sequence that does not contain blank characters, punctuation marks and is bounded by blank characters, punctuation marks or the string beginning/end.`
	s.en[57] = `A string with an English sentence is given. Remove all superfluous blank characters in the string, so that its words were separated by exactly one blank character.`
	s.en[58] = `A string that contains a <i>fully qualified path name</i> (that is, the drive and directory parts, the file name and extension) is given. Extract the file name (without the path and extension) from the string.`
	s.en[59] = `A string that contains a <i>fully qualified path name</i> (that is, the drive and directory parts, the file name and extension) is given. Extract the extension (without the preceding dot character) from the string.`
	s.en[60] = `A string that contains a <i>fully qualified path name</i> (that is, the drive and directory parts, the file name and extension) is given. Extract the first directory name (without backslashes &#34;\&#34;) from the string. If the file with the given name is located in the root directory then output a backslash.`
	s.en[61] = `A string that contains a <i>fully qualified path name</i> (that is, the drive and directory parts, the file name and extension) is given. Extract the last directory name (without backslashes &#34;\&#34;) from the string. If the file with the given name is located in the root directory then output a backslash.`
	s.en[62] = `A string with an English sentence is given. Encrypt the string using the right cyclic shift of any letter by one position of the English alphabet (for instance, the letter &#34;A&#34; is encoded by the letter &#34;B&#34;, &#34;a&#34; is encoded by &#34;b&#34;, &#34;B&#34; is encoded by &#34;C&#34;, &#34;z&#34; is encoded by &#34;a&#34;, and so on). Do not change blank characters and punctuation marks.`
	s.en[63] = `A string with an English sentence and an integer&nbsp;<i>K</i> (0&nbsp;&lt;&nbsp;<i>K</i>&nbsp;&lt;&nbsp;10) are given. Encrypt the string using the right cyclic shift of any letter by <i>K</i>&nbsp;positions of the English alphabet (for instance, if K&nbsp;=&nbsp;2 then the letter &#34;A&#34; is encoded by the letter &#34;C&#34;, &#34;a&#34; is encoded by &#34;c&#34;, &#34;B&#34; is encoded by &#34;D&#34;, &#34;z&#34; is encoded by &#34;b&#34;, and so on). Do not change blank characters and punctuation marks.`
	s.en[64] = `A string with an encrypted English sentence and an integer&nbsp;<i>K</i> (0&nbsp;&lt;&nbsp;<i>K</i>&nbsp;&lt;&nbsp;10) are given. The string is encrypted by means of the right cyclic shift of any letter by <i>K</i>&nbsp;positions of the English alphabet (see String63). Decrypt the given string.`
	s.en[65] = `A string with an encrypted English sentence and its decrypted first character&nbsp;<i>C</i> are given (the character&nbsp;<i>C</i> is always an English letter). The string is encrypted by means of the right cyclic shift of any letter by <i>K</i>&nbsp;positions of the English alphabet (see String63). Find the number&nbsp;<i>K</i> and decrypt the given string.`
	s.en[66] = `A string with an English sentence is given. Encrypt the string by moving all characters that are at the string positions with even numbers (2, 4,&nbsp;&#8230;) to the beginning of the string (in the same order) and moving all characters that are at the string positions with odd numbers (1, 3,&nbsp;&#8230;) to the end of the string (in inverse order). For instance, the string &#34;Program&#34; must be encrypted to &#34;rgamroP&#34;.`
	s.en[67] = `A string with an encrypted English sentence is given (the method of encryption is described in String66). Decrypt the string.`
	s.en[68] = `A string that contains digits and Latin small letters is given. If letters of the string are in alphabetic order then output&nbsp;0, otherwise output the order number of the first string character that is a letter and breaks the required order.`
	s.en[69] = `A string that contains Latin letters and parentheses &#34;(&#34;, &#34;)&#34; is given. If parentheses are in correct order (that is, each closing parenthesis&nbsp;&#34;)&#34; corresponds to an opening one &#34;(&#34;) then output&nbsp;0. If the string contains illegal parentheses &#34;)&#34; then output the order number of the first string character that is an illegal &#34;)&#34;. If the amount of closing parentheses is less than the amount of opening ones then output&nbsp;&#8722;1.`
	s.en[70] = `A string that contains Latin letters and brackets of three types (parentheses&nbsp;&#34;()&#34;, square brackets&nbsp;&#34;[]&#34;, braces &#34;{}&#34;) is given. If brackets are in correct order (that is, each closing bracket corresponds to an opening one of the same type) then output&nbsp;0. If the string contains illegal closing brackets then output the order number of the first string character that is an illegal closing bracket. If the amount of closing brackets is less than the amount of opening ones then output&nbsp;&#8722;1.`
}

func (s *sstring) makeTj() {
	s.tj = make([]string, s.count)
	s.tj[1] = `Рамзи C дода шудааст. Коди онро (яъне рақами тартибиро дар ҷадвали кодҳо) хориҷ кунед.`
	s.tj[2] = `Адади бутуни N (32 &lt;= N &lt;= 126) дода шудааст. Рамзеро, ки кодаш ба N баробар аст, хориҷ кунед.`
	s.tj[3] = `Рамзи C дода шудааст. Ду рамзеро хориҷ кунед, ки якӯми онҳо дар ҷадвали кодҳо пеш аз рамзи C аст, дуюмаш бошад, пас аз рамзи C аст.`
	s.tj[4] = `Адади бутуни N (1 &lt;= N &lt;= 26) дода шудааст. N-то ҳарфҳои калони аввалини алифбои лотиниро хориҷ кунед.`
	s.tj[5] = `Адади бутуни N (1 &lt;= N &lt;= 26) дода шудааст. N-то ҳарфҳои хурди охирини алифбои лотиниро бо тартиби баръакс (аз ҳарфи «z» сар карда) хориҷ кунед.`
	s.tj[6] = `Рамзи C дода шудааст, ки рақам ё ҳарф (лотинӣ ё тоҷикӣ)-ро тасвир мекунад. Агар C рақам бошад, пас сатри «digit»-ро хориҷ кунед, агар ҳарфи лотинӣ бошад — сарти «lat»-ро хориҷ кунед, агар тоҷикӣ бошад — сатри «toj»-ро хориҷ кунед.`
	s.tj[7] = `Сатре, ки холӣ нест, дода шудааст. Кодҳои рамзҳои аввалин ва охирини онро хориҷ кунед.`
	s.tj[8] = `Адади бутуни N (> 0) ва рамзи C дода шудаанд. Сатри дарозии N-ро хориҷ кунед, ки он аз рамзҳои C таркиб ёфтааст.`
	s.tj[9] = `Адади ҷуфти N (> 0) ва рамзҳои C<sub>1</sub> ва C<sub>2</sub> дода шудаанд. Сатри дарозии N-ро хориҷ кунед, ки он аз рамзҳои пайиҳамояндаи C<sub>1</sub> ва C<sub>2</sub>, аз C<sub>1</sub> сар карда, таркиб ёфтааст.`
	s.tj[10] = `Сатр дода шудааст. Сатреро хориҷ кунед, ки худи ҳамон рамзҳоро дорост, аммо бо тартиби баръакс ҷойгир шудаанд.`
	s.tj[11] = `Сатри S, ки холӣ нест, дода шудааст. Сатреро хориҷ кунед, ки рамзҳои сатри S-ро дорост, ки дар байни онҳо яктогӣ фосила гузошта шудаст.`
	s.tj[12] = `Сатри нохолии (холӣ нест) S ва адади бутуни N (> 0) дода шудаанд. Сатреро хориҷ кунед, ки рамзҳои сатри S-ро дорост, ки дар байни онҳо N-тогӣ рамзҳои «*» (ситорача) гузошта шудаанд.`
	s.tj[13] = `Сатр дода шудааст. Миқдори рақамҳои дар он мавҷударо ҳисоб кунед.`
	s.tj[14] = `Сатр дода шудааст. Миқдори ҳарфҳои калони лотинии дар он мавҷударо ҳисоб кунед.`
	s.tj[15] = `Сатр дода шудааст. Миқдори умумии ҳарфҳои хурди лотинӣ ва тоҷикии дар он мавҷударо ҳисоб кунед.`
	s.tj[16] = `Сатр дода шудааст. Ҳамаи ҳарфҳои калони лотинии дар он мавҷударо ба хурд табдил диҳед.`
	s.tj[17] = `Сатр дода шудааст. Ҳамаи ҳарфҳои хурд (ҳам лотинӣ ва ҳам тоҷикӣ)-и дар он мавҷударо ба калон табдил диҳед.`
	s.tj[18] = `Сатр дода шудааст. Ҳамаи ҳарфҳои хурд (ҳам лотинӣ ва ҳам тоҷикӣ)-и дар он мавҷударо ба калон табдил диҳед, ҳарфҳои калонро бошад — ба хурд табдил диҳед.`
	s.tj[19] = `Сатр дода шудааст. Агар он навишти адади бутунро нишон диҳад, пас 1(як)-ро хориҷ кунед, агар адади ҳақиқӣ (бо қисми касрӣ)-ро нишон диҳад — 2(ду)-ро хориҷ кунед; агар сатрро ба адад табдил додан имконнопазир бошад, пас 0(нул)-ро хориҷ кунед. Ба назар гиред, ки қисми касрии адади ҳақиқӣ аз қисми бутунаш бо нуқтаи даҳӣ «.» ҷудо карда мешавад.`
	s.tj[20] = `Адади бутуни мусбӣ дода шудааст. Рамзҳоеро хориҷ кунед, ки рақамҳои ин ададро тасвир мекунанд (бо тартиби аз чап ба рост).`
	s.tj[21] = `Адади бутуни мусбӣ дода шудааст. Рамзҳоеро хориҷ кунед, ки рақамҳои ин ададро тасвир мекунанд (бо тартиби аз рост ба чап).`
	s.tj[22] = `Сатре, ки адади бутуни мусбиро тасвир мекунад, дода шудааст. Суммаи рақамҳои ин ададро хориҷ кунед.`
	s.tj[23] = `Сатре, ки ифодаи арифметикии намуди «<рақам>±<рақам>±…±<рақам>»-ро тасвир мекунад, ки дар ин ҷо дар ҷои аломати амалҳои «±» рамзи «+» ё «–» ҷойгир аст (масалан, «4+7–2–8»), дода шудааст. Қимати ифодаи мазкур (адади бутун)-ро хориҷ кунед.`
	s.tj[24] = `Сатре, ки навишти дуии адади бутуни мусбиро тасвир мекунад, дода шудааст. Сатреро хориҷ кунед, ки навишти даҳии худи ҳамин ададро тасвир мекунад.`
	s.tj[25] = `Сатре, ки навишти даҳии адади бутуни мусбиро тасвир мекунад, дода шудааст. Сатреро хориҷ кунед, ки навишти дуии худи ҳамин ададро тасвир мекунад.`
	s.tj[26] = `Адади бутуни N (> 0) ва сатри S дода шудаанд. Сатри S-ро ба сатри дарозии N ба тариқи зайл табдил диҳед: агар дарозии сатри S аз N калон бошад, пас рамзҳои аввалинро партоед, агар дарозии сатри S аз N хурд бошад, пас ба аввали он рамзҳои «.» (нуқта)-ро ҳамроҳ кунед.`
	s.tj[27] = `Ададҳои бутуни мусбии N<sub>1</sub> ва N<sub>2</sub> ва сатрҳои S<sub>1</sub> ва S<sub>2</sub> дода шудаанд. Аз ин сатрҳо сатри наверо ҳосил кунед, ки N<sub>1</sub>-то рамзҳои аввалини сатри S<sub>1</sub>-ро доро буда, N<sub>2</sub>-то рамзҳои охирини сатри S<sub>2</sub>-ро доро бошад (бо тартиби нишондодашуда).`
	s.tj[28] = `Рамзи C ва сатри S дода шудаанд. Ҳар як мавқеи рамзи C-ро дар сатри S дуто кунед.`
	s.tj[29] = `Рамзи C ва сатрҳои S, S<sub>0</sub> дода шудаанд. Пеш аз ҳар як мавқеи рамзи C дар сатри S сатри S<sub>0</sub>-ро гузоред.`
	s.tj[30] = `Рамзи C ва сатрҳои S, S<sub>0</sub> дода шудаанд. Пас аз ҳар як мавқеи рамзи C дар сатри S сатри S<sub>0</sub>-ро гузоред.`
	s.tj[31] = `Сатрҳои S ва S<sub>0</sub> дода шудаанд. Мавҷудияти сатри S<sub>0</sub>-ро дар сатри S санҷед. Агар мавҷуд бошад, пас True-ро хориҷ кунед, агар мавҷуд набошад, пас False-ро хориҷ кунед.`
	s.tj[32] = `Сатрҳои S ва S<sub>0</sub> дода шудаанд. Миқдори мавқеъҳои сатри S<sub>0</sub>-ро дар сатри S ёбед.`
	s.tj[33] = `Сатрҳои S ва S<sub>0</sub> дода шудаанд. Аз сатри S зерсатри аввалинеро, ки бо сатри S<sub>0</sub> мувофиқ меояд, нест кунед. Агар зерсатри мувофиқоянда набошад, пас сатри S-ро бе тағйирот хориҷ кунед.`
	s.tj[34] = `Сатрҳои S ва S<sub>0</sub> дода шудаанд. Аз сатри S зерсатри охиринеро, ки бо сатри S<sub>0</sub> мувофиқ меояд, нест кунед. Агар зерсатри мувофиқоянда набошад, пас сатри S-ро бе тағйирот хориҷ кунед.`
	s.tj[35] = `Сатрҳои S ва S<sub>0</sub> дода шудаанд. Аз сатри S ҳамаи зерсатрҳои бо сатри S<sub>0</sub> мувофиқояндаро нест кунед. Агар зерсатри мувофиқоянда набошад, пас сатри S-ро бе тағйирот хориҷ кунед.`
	s.tj[36] = `Сатрҳои S, S<sub>1</sub> ва S<sub>2</sub> дода шудаанд. Дар сатри S мавқеи якӯми сатри S<sub>1</sub>-ро ба сатри S<sub>2</sub> иваз кунед.`
	s.tj[37] = `Сатрҳои S, S<sub>1</sub> ва S<sub>2</sub> дода шудаанд. Дар сатри S мавқеи охирини сатри S<sub>1</sub>-ро ба сатри S<sub>2</sub> иваз кунед.`
	s.tj[38] = `Сатрҳои S, S<sub>1</sub> ва S<sub>2</sub> дода шудаанд. Дар сатри S ҳамаи мавқеъҳои сатри S<sub>1</sub>-ро ба сатри S<sub>2</sub> иваз кунед.`
	s.tj[39] = `Сатре, ки ба ҳадди имкон як рамзи фосила дорад, дода шудааст. Зерсатреро хориҷ кунед, ки дар байни фосилаи якӯм ва дуюми сатри ибтидоӣ ҷойгир аст. Агар сатр фақат якто фосила дошта бошад, пас сатри холиро хориҷ кунед.`
	s.tj[40] = `Сатре, ки ба ҳадди имкон як рамзи фосила дорад, дода шудааст. Зерсатреро хориҷ кунед, ки дар байни фосилаи якӯм ва охирини сатри ибтидоӣ ҷойгир аст. Агар сатр фақат якто фосила дошта бошад, пас сатри холиро хориҷ кунед.`
	s.tj[41] = `Сатре, ки аз калимаҳои тоҷикии бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Миқдори калимаҳоро дар сатр ёбед.`
	s.tj[42] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Миқдори калимаҳоеро, ки онҳо бо як ҳарф сар шуда бо худи ҳамон ҳарф тамом мешаванд, ёбед.`
	s.tj[43] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Миқдори калимаҳоеро, ки онҳо аққалан якто ҳарфи «А»-ро доростанд, ёбед.`
	s.tj[44] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Миқдори калимаҳоеро, ки онҳо расо сето ҳарфи «А»-ро доростанд, ёбед.`
	s.tj[45] = `Сатре, ки аз калимаҳои тоҷикии бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Дарозии калимаи кӯтоҳтаринро ёбед.`
	s.tj[46] = `Сатре, ки аз калимаҳои тоҷикии бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Дарозии калимаи дарозтаринро ёбед.`
	s.tj[47] = `Сатре, ки аз калимаҳои тоҷикии бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Сатреро хориҷ кунед, ки худи ҳамин калимаҳоро дорост, ки онҳо бо якто рамзи «.» (нуқта) ҷудо карда шудаанд. Дар охири сатр нуқта нагузоред.`
	s.tj[48] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Барои ҳар як калимаи сатр ҳамаи мавқеъҳои ояндаи ҳарфи якӯми онро ба рамзи «.» (нуқта) иваз карда, ҳар як калимаи сатрро табдил диҳед. Масалан, калимаи «МИНИМУМ»-ро бояд ба «МИНИ.У.» табдил дод. Миқдори фосилаҳои байни калимаҳоро тағйир надиҳед.`
	s.tj[49] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Барои ҳар як калимаи сатр ҳамаи мавқеъҳои пешинаи ҳарфи охирини онро ба рамзи «.» (нуқта) иваз карда, ҳар як калимаи сатрро табдил диҳед. Масалан, калимаи «МИНИМУМ»-ро бояд ба «.ИНИ.УМ» табдил дод. Миқдори фосилаҳои байни калимаҳоро тағйир надиҳед.`
	s.tj[50] = `Сатре, ки аз калимаҳои тоҷикии бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Сатреро хориҷ кунед, ки худи ҳамин калимаҳоро дорост, ки онҳо бо як фосила ҷудо карда шуда, бо тартиби баръакс ҷойгир шудаанд.`
	s.tj[51] = `Сатре, ки аз калимаҳои тоҷикии бо ҳарфҳои калон навишташуда ва бо фосилаҳо ҷудокардашуда (бо якто ё якчандто фосилаҳо) иборат аст, дода шудааст. Сатреро хориҷ кунед, ки худи ҳамин калимаҳоро дорост, ки онҳо бо як фосила ҷудо карда шуда, бо тартиби алифбоӣ ҷойгир шудаанд.`
	s.tj[52] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Сатрро чунон тағйир диҳед, ки ҳар як калима бо ҳарфи калон сар шавад. Калима гуфта, маҷмӯи рамзҳоеро мегӯянд, ки фосила надорад ва бо фосилаҳо ё оғоз/анҷоми сатр маҳдуд аст. Калимаеро, ки бо ҳарф сар намешавад, тағйир надиҳед.`
	s.tj[53] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Миқдори аломатҳои китобатии дар сатр мавҷударо ҳисоб кунед.`
	s.tj[54] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Миқдори ҳарфҳои садоноки дар сатр мавҷударо ҳисоб кунед.`
	s.tj[55] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Калимаи дарозтарини ҷумларо хориҷ кунед. Агар чунин калимаҳо якчандто бошанд, пас якӯмини онҳоро хориҷ кунед. Калима гуфта, маҷмӯи рамзҳоеро мегӯянд, ки фосилаҳо ва аломатҳои китобатӣ надорад ва бо фосилаҳо, аломатҳои китобатӣ ё оғоз/анҷоми сатр маҳдуд аст.`
	s.tj[56] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Калимаи кӯтоҳтарини ҷумларо хориҷ кунед. Агар чунин калимаҳо якчандто бошанд, пас охирини онҳоро хориҷ кунед. Калима гуфта, маҷмӯи рамзҳоеро мегӯянд, ки фосилаҳо ва аломатҳои китобатӣ надорад ва бо фосилаҳо, аломатҳои китобатӣ ё оғоз/анҷоми сатр маҳдуд аст.`
	s.tj[57] = `Сатр-ҷумла бо фосилаҳои барзиёди байни калимаҳо дода шудааст. Онро чунон табдил диҳед, ки дар байни калимаҳо расо якто фосила бошад.`
	s.tj[58] = `Сатре, ки номи пурраи файлро, яъне номи диск, рӯйхати пӯшаҳо (роҳ), худи ном ва номи иловагиро дорост, дода шудааст. Аз ин сатр номи файлро (бе номи иловагӣ) ҷудо кунед.`
	s.tj[59] = `Сатре, ки номи пурраи файлро, яъне номи диск, рӯйхати пӯшаҳо (роҳ), худи ном ва номи иловагиро дорост, дода шудааст. Аз ин сатр номи иловагии файлро (бе нуқтаи пешина) ҷудо кунед.`
	s.tj[60] = `Сатре, ки номи пурраи файлро дорост, дода шудааст. Аз ин сатр номи пӯшаи аввалинро (бе рамзи «\») ҷудо кунед. Агар файл дар пӯшаи асосӣ ҷойгир бошад, пас рамзи «\»-ро хориҷ кунед.`
	s.tj[61] = `Сатре, ки номи пурраи файлро дорост, дода шудааст. Аз ин сатр номи пӯшаи охиринро (бе рамзи «\») ҷудо кунед. Агар файл дар пӯшаи асосӣ ҷойгир бошад, пас рамзи «\»-ро хориҷ кунед.`
	s.tj[62] = `Сатр-ҷумла бо забони тоҷикӣ дода шудааст. Онро бо роҳи ивази даврии ҳар як ҳарф ба ҳарфи пасинаи он дар алифбо рамзгузорӣ кунед ва дар ин ҳангом калонию хурдӣ (андоза)-и ҳарфҳоро нигоҳ доред («А» ба «Б» мекӯчад, «а» — ба «б», «Б» — ба «В», «я» — ба «а» ва ғ.). Ҳарфи «ё»-ро дар алифбо ба инобат нагиред («е» бояд ба «ж» кӯчад). Аломатҳои китобатӣ ва фосилаҳоро тағйир надиҳед.`
	s.tj[63] = `Сатр-ҷумла бо забони тоҷикӣ ва адади K (0 < K < 10) дода шудаанд. Сатрро бо роҳи ивази даврии ҳар як ҳарф ба ҳарфи ҳамон андоза (калонию хурдӣ), ки дар алифбо дар мавқеи K-ӯм пас аз ҳарфи рамзгузоришаванда меистад (масалан, барои K = 2 «А» ба «В» мекӯчад, «а» — ба «в», «Б» — ба «Г», «я» — ба «б» ва ғ.), рамзгузорӣ кунед. Ҳарфи «ё»-ро дар алифбо ба инобат нагиред, аломатҳои китобатӣ ва фосилаҳоро тағйир надиҳед.`
	s.tj[64] = `Ҷумлаи рамзгузоришуда бо забони тоҷикӣ (тарзи рамзгузорӣ дар масъалаи String63 оварда шудааст) ва ҷойивазкунии кодии K (0 < K < 10) дода шудаанд. Ҷумларо аз рамз бароред.`
	s.tj[65] = `Ҷумлаи рамзгузоришуда бо забони тоҷикӣ (тарзи рамзгузорӣ дар масъалаи String63 оварда шудааст) ва рамзи аввалини азрамзбаровардашудаи он C дода шудаанд. Ҷойивазкунии кодӣ K-ро ёбед ва ҷумларо аз рамз бароред.`
	s.tj[66] = `Сатр-ҷумла дода шудааст. Онро бо чунин тарз ба рамз дароред: дар аввал ҳамаи рамзҳои дар мақомҳои ҷуфти сатр ҷойгирбударо ҷойгир кунед, сонӣ бо тартиби баръакс ҳамаи рамзҳои дар мақомҳои тоқ ҷойгирбударо ҷойгир кунед (масалан, сатри «Программа» ба «ргамамроП» табдил меёбад).`
	s.tj[67] = `Ҷумлае, ки аз рӯи қоидаи дар масъалаи String66 овардашуда рамзгузорӣ карда шудааст, дода шудааст. Ин ҷумларо аз рамз бароред.`
	s.tj[68] = `Сатре, ки рақамҳо ва ҳарфҳои хурди лотиниро дорост, дода шудааст. Агар ҳарфҳо дар сатр аз рӯи алифбо ба тартиб гузошта шуда бошанд, пас 0(нул)-ро хориҷ кунед; дар ҳолати акс рақами тартибии рамзи аввалини сатрро, ки тартиби алифборо вайрон мекунад, хориҷ кунед.`
	s.tj[69] = `Сатре, ки ҳарфҳои лотинӣ ва қавсҳои давриро дорост, дода шудааст. Агар қавсҳо дуруст гузошта шуда бошанд (яъне ба ҳар як қавси кушодашуда якто қавси пӯшидашаванда мувофиқ меояд), пас 0(нул)-ро хориҷ кунед. Дар ҳолати акс ё рақами тартибии мавқееро хориҷ кунед, ки дар он қавси пӯшидашавандаи нодурусти аввалин гузошта шудааст, ё агар қавсҳои пӯшидашаванда кам бошанд, адади -1(яки манфӣ)-ро хориҷ кунед.`
	s.tj[70] = `Сатре, ки ҳарфҳои лотинӣ ва қавсҳои се намудро дорост: «()», «[]», «{}», дода шудааст. Агар қавсҳо дуруст гузошта шуда бошанд (яъне ба ҳар як қавси кушодашуда якто қавси пӯшидашавандаи ҳамон намуд мувофиқ меояд), пас 0(нул)-ро хориҷ кунед. Дар ҳолати акс ё рақами мавқееро хориҷ кунед, ки дар он қавси нодурусти аввалин гузошта шудааст, ё агар қавсҳои пӯшидашаванда кам бошанд, адади –1(яки манфӣ)-ро хориҷ кунед.`
}

func (s *sstring) makeRu() {
	s.ru = make([]string, s.count)
	s.ru[1] = `Дан символ&nbsp;<i>C</i>. Вывести его <i>код</i> (т.&nbsp;е. номер в кодовой таблице).`
	s.ru[2] = `Дано целое число&nbsp;<i>N</i> (32&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;126). Вывести символ с кодом, равным&nbsp;<i>N</i>.`
	s.ru[3] = `Дан символ&nbsp;<i>C</i>. Вывести два символа, первый из которых предшествует символу&nbsp;<i>C</i> в кодовой таблице, а второй следует за символом&nbsp;<i>C</i>.`
	s.ru[4] = `Дано целое число&nbsp;<i>N</i> (1&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;26). Вывести <i>N</i>&nbsp;первых <i>прописных</i> (т.&nbsp;е. заглавных) букв латинского алфавита.`
	s.ru[5] = `Дано целое число&nbsp;<i>N</i> (1&nbsp;&#8804;&nbsp;<i>N</i>&nbsp;&#8804;&nbsp;26). Вывести <i>N</i>&nbsp;последних <i>строчных</i> (т.&nbsp;е. маленьких) букв латинского алфавита в обратном порядке (начиная с буквы &#171;z&#187;).`
	s.ru[6] = `Дан символ&nbsp;<i>C</i>, изображающий цифру или букву (латинскую или русскую). Если&nbsp;<i>C</i> изображает цифру, то вывести строку &#171;digit&#187;, если латинскую букву &#8212; вывести строку &#171;lat&#187;, если русскую &#8212; вывести строку &#171;rus&#187;.`
	s.ru[7] = `Дана непустая строка. Вывести коды ее первого и последнего символа.`
	s.ru[8] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0) и символ&nbsp;<i>C</i>. Вывести строку длины&nbsp;<i>N</i>, которая состоит из символов&nbsp;<i>C</i>.`
	s.ru[9] = `Дано четное число&nbsp;<i>N</i> (&gt;&nbsp;0) и символы&nbsp;<i>C</i><sub>1</sub> и&nbsp;<i>C</i><sub>2</sub>. Вывести строку длины&nbsp;<i>N</i>, которая состоит из чередующихся символов&nbsp;<i>C</i><sub>1</sub> и&nbsp;<i>C</i><sub>2</sub>, начиная с&nbsp;<i>C</i><sub>1</sub>.`
	s.ru[10] = `Дана строка. Вывести строку, содержащую те же символы, но расположенные в обратном порядке.`
	s.ru[11] = `Дана непустая строка&nbsp;<i>S</i>. Вывести строку, содержащую символы строки&nbsp;<i>S</i>, между которыми вставлено по одному пробелу.`
	s.ru[12] = `Дана непустая строка&nbsp;<i>S</i> и целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Вывести строку, содержащую символы строки&nbsp;<i>S</i>, между которыми вставлено по <i>N</i>&nbsp;символов &#171;*&#187; (звездочка).`
	s.ru[13] = `Дана строка. Подсчитать количество содержащихся в ней цифр.`
	s.ru[14] = `Дана строка. Подсчитать количество содержащихся в ней прописных латинских букв.`
	s.ru[15] = `Дана строка. Подсчитать общее количество содержащихся в ней строчных латинских и русских букв.`
	s.ru[16] = `Дана строка. Преобразовать в ней все прописные латинские буквы в строчные.`
	s.ru[17] = `Дана строка. Преобразовать в ней все строчные буквы (как латинские, так и русские) в прописные.`
	s.ru[18] = `Дана строка. Преобразовать в ней все строчные буквы (как латинские, так и русские) в прописные, а прописные &#8212; в строчные.`
	s.ru[19] = `Дана строка. Если она представляет собой запись целого числа, то вывести&nbsp;1, если вещественного (с дробной частью) &#8212; вывести&nbsp;2; если строку нельзя преобразовать в число, то вывести&nbsp;0. Считать, что дробная часть вещественного числа отделяется от его целой части десятичной <i>точкой</i>&nbsp;&#171;.&#187;.`
	s.ru[20] = `Дано целое положительное число. Вывести символы, изображающие цифры этого числа (в порядке слева направо).`
	s.ru[21] = `Дано целое положительное число. Вывести символы, изображающие цифры этого числа (в порядке справа налево).`
	s.ru[22] = `Дана строка, изображающая целое положительное число. Вывести сумму цифр этого числа.`
	s.ru[23] = `Дана строка, изображающая арифметическое выражение вида &#171;&lt;цифра&gt;&#177;&lt;цифра&gt;&#177;&#8230;&#177;&lt;цифра&gt;&#187;, где на месте знака операции &#171;&#177;&#187; находится символ &#171;+&#187; или &#171;&#8722;&#187; (например, &#171;4+7&#8722;2&#8722;8&#187;). Вывести значение данного выражения (целое число).`
	s.ru[24] = `Дана строка, изображающая двоичную запись целого положительного числа. Вывести строку, изображающую десятичную запись этого же числа.`
	s.ru[25] = `Дана строка, изображающая десятичную запись целого положительного числа. Вывести строку, изображающую двоичную запись этого же числа.`
	s.ru[26] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0) и строка&nbsp;<i>S</i>. Преобразовать строку&nbsp;<i>S</i> в строку длины&nbsp;<i>N</i> следующим образом: если длина строки&nbsp;<i>S</i> больше&nbsp;<i>N</i>, то отбросить первые символы, если длина строки&nbsp;<i>S</i> меньше&nbsp;<i>N</i>, то в ее начало добавить символы&nbsp;&#171;.&#187; (точка).`
	s.ru[27] = `Даны целые положительные числа&nbsp;<i>N</i><sub>1</sub> и&nbsp;<i>N</i><sub>2</sub> и строки&nbsp;<i>S</i><sub>1</sub> и&nbsp;<i>S</i><sub>2</sub>. Получить из этих строк новую строку, содержащую первые <i>N</i><sub>1</sub>&nbsp;символов строки&nbsp;<i>S</i><sub>1</sub> и последние <i>N</i><sub>2</sub>&nbsp;символов строки&nbsp;<i>S</i><sub>2</sub> (в указанном порядке).`
	s.ru[28] = `Дан символ&nbsp;<i>C</i> и строка&nbsp;<i>S</i>. Удвоить каждое вхождение символа&nbsp;<i>C</i> в строку&nbsp;<i>S</i>.`
	s.ru[29] = `Дан символ&nbsp;<i>C</i> и строки&nbsp;<i>S</i>, <i>S</i><sub>0</sub>. Перед каждым вхождением символа&nbsp;<i>C</i> в строку&nbsp;<i>S</i> вставить строку&nbsp;<i>S</i><sub>0</sub>.`
	s.ru[30] = `Дан символ&nbsp;<i>C</i> и строки&nbsp;<i>S</i>, <i>S</i><sub>0</sub>. После каждого вхождения символа&nbsp;<i>C</i> в строку&nbsp;<i>S</i> вставить строку&nbsp;<i>S</i><sub>0</sub>.`
	s.ru[31] = `Даны строки&nbsp;<i>S</i> и&nbsp;<i>S</i><sub>0</sub>. Проверить, содержится ли строка&nbsp;<i>S</i><sub>0</sub> в строке&nbsp;<i>S</i>. Если содержится, то вывести true, если не содержится, то вывести false.`
	s.ru[32] = `Даны строки&nbsp;<i>S</i> и&nbsp;<i>S</i><sub>0</sub>. Найти количество вхождений строки&nbsp;<i>S</i><sub>0</sub> в строку&nbsp;<i>S</i>.`
	s.ru[33] = `Даны строки&nbsp;<i>S</i> и&nbsp;<i>S</i><sub>0</sub>. Удалить из строки&nbsp;<i>S</i> первую подстроку, совпадающую с&nbsp;<i>S</i><sub>0</sub>. Если совпадающих подстрок нет, то вывести строку&nbsp;<i>S</i> без изменений.`
	s.ru[34] = `Даны строки&nbsp;<i>S</i> и&nbsp;<i>S</i><sub>0</sub>. Удалить из строки&nbsp;<i>S</i> последнюю подстроку, совпадающую с&nbsp;<i>S</i><sub>0</sub>. Если совпадающих подстрок нет, то вывести строку&nbsp;<i>S</i> без изменений.`
	s.ru[35] = `Даны строки&nbsp;<i>S</i> и&nbsp;<i>S</i><sub>0</sub>. Удалить из строки&nbsp;<i>S</i> все подстроки, совпадающие с&nbsp;<i>S</i><sub>0</sub>. Если совпадающих подстрок нет, то вывести строку&nbsp;<i>S</i> без изменений.`
	s.ru[36] = `Даны строки&nbsp;<i>S</i>, <i>S</i><sub>1</sub> и&nbsp;<i>S</i><sub>2</sub>. Заменить в строке&nbsp;<i>S</i> первое вхождение строки&nbsp;<i>S</i><sub>1</sub> на строку&nbsp;<i>S</i><sub>2</sub>.`
	s.ru[37] = `Даны строки&nbsp;<i>S</i>, <i>S</i><sub>1</sub> и&nbsp;<i>S</i><sub>2</sub>. Заменить в строке&nbsp;<i>S</i> последнее вхождение строки&nbsp;<i>S</i><sub>1</sub> на строку&nbsp;<i>S</i><sub>2</sub>.`
	s.ru[38] = `Даны строки&nbsp;<i>S</i>, <i>S</i><sub>1</sub> и&nbsp;<i>S</i><sub>2</sub>. Заменить в строке&nbsp;<i>S</i> все вхождения строки&nbsp;<i>S</i><sub>1</sub> на строку&nbsp;<i>S</i><sub>2</sub>.`
	s.ru[39] = `Дана строка, содержащая по крайней мере один символ пробела. Вывести подстроку, расположенную между первым и вторым пробелом исходной строки. Если строка содержит только один пробел, то вывести пустую строку.`
	s.ru[40] = `Дана строка, содержащая по крайней мере один символ пробела. Вывести подстроку, расположенную между первым и последним пробелом исходной строки. Если строка содержит только один пробел, то вывести пустую строку.`
	s.ru[41] = `Дана строка, состоящая из русских слов, разделенных пробелами (одним или несколькими). Найти количество слов в строке.`
	s.ru[42] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Найти количество слов, которые начинаются и заканчиваются одной и той же буквой.`
	s.ru[43] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Найти количество слов, которые содержат хотя бы одну букву&nbsp;&#171;А&#187;.`
	s.ru[44] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Найти количество слов, которые содержат ровно три буквы&nbsp;&#171;А&#187;.`
	s.ru[45] = `Дана строка, состоящая из русских слов, разделенных пробелами (одним или несколькими). Найти длину самого короткого слова.`
	s.ru[46] = `Дана строка, состоящая из русских слов, разделенных пробелами (одним или несколькими). Найти длину самого длинного слова.`
	s.ru[47] = `Дана строка, состоящая из русских слов, разделенных пробелами (одним или несколькими). Вывести строку, содержащую эти же слова, разделенные одним символом &#171;.&#187; (точка). В конце строки точку не ставить.`
	s.ru[48] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Преобразовать каждое слово в строке, заменив в нем все последующие вхождения его первой буквы на символ &#171;.&#187; (точка). Например, слово &#171;МИНИМУМ&#187; надо преобразовать в &#171;МИНИ.У.&#187;. Количество пробелов между словами не изменять.`
	s.ru[49] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Преобразовать каждое слово в строке, заменив в нем все предыдущие вхождения его последней буквы на символ &#171;.&#187; (точка). Например, слово &#171;МИНИМУМ&#187; надо преобразовать в &#171;.ИНИ.УМ&#187;. Количество пробелов между словами не изменять.`
	s.ru[50] = `Дана строка, состоящая из русских слов, разделенных пробелами (одним или несколькими). Вывести строку, содержащую эти же слова, разделенные одним пробелом и расположенные в обратном порядке.`
	s.ru[51] = `Дана строка, состоящая из русских слов, набранных заглавными буквами и разделенных пробелами (одним или несколькими). Вывести строку, содержащую эти же слова, разделенные одним пробелом и расположенные в алфавитном порядке.`
	s.ru[52] = `Дана строка-предложение на русском языке. Преобразовать строку так, чтобы каждое слово начиналось с заглавной буквы. <i>Словом</i> считать набор символов, не содержащий пробелов и ограниченный пробелами или началом/концом строки. Слова, не начинающиеся с буквы, не изменять.`
	s.ru[53] = `Дана строка-предложение на русском языке. Подсчитать количество содержащихся в строке знаков препинания.`
	s.ru[54] = `Дана строка-предложение на русском языке. Подсчитать количество содержащихся в строке гласных букв.`
	s.ru[55] = `Дана строка-предложение на русском языке. Вывести самое длинное слово в предложении. Если таких слов несколько, то вывести первое из них. <i>Словом</i> считать набор символов, не содержащий пробелов, знаков препинания и ограниченный пробелами, знаками препинания или началом/концом строки.`
	s.ru[56] = `Дана строка-предложение на русском языке. Вывести самое короткое слово в предложении. Если таких слов несколько, то вывести последнее из них. <i>Словом</i> считать набор символов, не содержащий пробелов, знаков препинания и ограниченный пробелами, знаками препинания или началом/концом строки.`
	s.ru[57] = `Дана строка-предложение с избыточными пробелами между словами. Преобразовать ее так, чтобы между словами был ровно один пробел.`
	s.ru[58] = `Дана строка, содержащая <i>полное имя файла</i>, т.&nbsp;е. имя диска, список каталогов (путь), собственно имя и расширение. Выделить из этой строки имя файла (без расширения).`
	s.ru[59] = `Дана строка, содержащая <i>полное имя файла</i>, т.&nbsp;е. имя диска, список каталогов (путь), собственно имя и расширение. Выделить из этой строки расширение файла (без предшествующей точки).`
	s.ru[60] = `Дана строка, содержащая полное имя файла. Выделить из этой строки название первого каталога (без символов&nbsp;&#171;\&#187;). Если файл содержится в корневом каталоге, то вывести символ&nbsp;&#171;\&#187;.`
	s.ru[61] = `Дана строка, содержащая полное имя файла. Выделить из этой строки название последнего каталога (без символов&nbsp;&#171;\&#187;). Если файл содержится в корневом каталоге, то вывести символ&nbsp;&#171;\&#187;.`
	s.ru[62] = `Дана строка-предложение на русском языке. Зашифровать ее, выполнив циклическую замену каждой буквы на следующую за ней в алфавите и сохранив при этом регистр букв (&#171;А&#187; перейдет в&nbsp;&#171;Б&#187;, &#171;а&#187; &#8212; в&nbsp;&#171;б&#187;, &#171;Б&#187; &#8212; в&nbsp;&#171;В&#187;, &#171;я&#187; &#8212; в&nbsp;&#171;а&#187; и т.&nbsp;д.). Букву&nbsp;&#171;ё&#187; в алфавите не учитывать (&#171;е&#187; должна переходить в&nbsp;&#171;ж&#187;). Знаки препинания и пробелы не изменять.`
	s.ru[63] = `Дана строка-предложение на русском языке и число&nbsp;<i>K</i> (0&nbsp;&lt;&nbsp;<i>K</i>&nbsp;&lt;&nbsp;10). Зашифровать строку, выполнив циклическую замену каждой буквы на букву того же регистра, расположенную в алфавите на <i>K</i>-й позиции после шифруемой буквы (например, для <i>K</i>&nbsp;=&nbsp;2 &#171;А&#187; перейдет в&nbsp;&#171;В&#187;, &#171;а&#187; &#8212; в&nbsp;&#171;в&#187;, &#171;Б&#187; &#8212; в&nbsp;&#171;Г&#187;, &#171;я&#187; &#8212; в&nbsp;&#171;б&#187; и т.&nbsp;д.). Букву&nbsp;&#171;ё&#187; в алфавите не учитывать, знаки препинания и пробелы не изменять.`
	s.ru[64] = `Дано зашифрованное предложение на русском языке (способ шифрования описан в задании String63) и кодовое смещение&nbsp;<i>K</i> (0&nbsp;&lt;&nbsp;<i>K</i>&nbsp;&lt;&nbsp;10). Расшифровать предложение.`
	s.ru[65] = `Дано зашифрованное предложение на русском языке (способ шифрования описан в задании String63) и его расшифрованный первый символ&nbsp;<i>C</i>. Найти кодовое смещение&nbsp;<i>K</i> и расшифровать предложение.`
	s.ru[66] = `Дана строка-предложение. Зашифровать ее, поместив вначале все символы, расположенные на четных позициях строки, а затем, в обратном порядке, все символы, расположенные на нечетных позициях (например, строка &#171;Программа&#187; превратится в &#171;ргамамроП&#187;).`
	s.ru[67] = `Дано предложение, зашифрованное по правилу, описанному в задании String66. Расшифровать это предложение.`
	s.ru[68] = `Дана строка, содержащая цифры и строчные латинские буквы. Если буквы в строке упорядочены по алфавиту, то вывести&nbsp;0; в противном случае вывести номер первого символа строки, нарушающего алфавитный порядок.`
	s.ru[69] = `Дана строка, содержащая латинские буквы и круглые скобки. Если скобки расставлены правильно (т.&nbsp;е. каждой открывающей соответствует одна закрывающая), то вывести число&nbsp;0. В противном случае вывести или номер позиции, в которой расположена первая ошибочная закрывающая скобка, или, если закрывающих скобок не хватает, число&nbsp;&#8722;1.`
	s.ru[70] = `Дана строка, содержащая латинские буквы и скобки трех видов:&nbsp;&#171;()&#187;, &#171;[]&#187;,&nbsp;&#171;{}&#187;. Если скобки расставлены правильно (т.&nbsp;е. каждой открывающей соответствует закрывающая скобка того же вида), то вывести число&nbsp;0. В противном случае вывести или номер позиции, в которой расположена первая ошибочная скобка, или, если закрывающих скобок не хватает, число&nbsp;&#8722;1.`
}
