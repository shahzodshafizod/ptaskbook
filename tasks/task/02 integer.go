package task

type iinteger struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewIntegerTask() Task {
	i := &iinteger{
		count: 31,
		name:  "Integer",
	}
	i.makeEn()
	i.makeTj()
	i.makeRu()
	return i
}

func (i *iinteger) Count() int   { return i.count }
func (b *iinteger) Name() string { return b.name }
func (i *iinteger) En() []string { return i.en }
func (i *iinteger) Tj() []string { return i.tj }
func (i *iinteger) Ru() []string { return i.ru }

func (i *iinteger) makeEn() {
	i.en = make([]string, i.count)
	i.en[1] = `A distance&nbsp;<i>L</i> is given in centimeters. Find the amount of full meters of this distance (1&nbsp;m&nbsp;=&nbsp;1000&nbsp;cm). Use the operator of integer division.`
	i.en[2] = `A weight&nbsp;<i>M</i> is given in kilograms. Find the amount of full tons of this weight (1&nbsp;ton&nbsp;=&nbsp;1000&nbsp;kg). Use the operator of integer division.`
	i.en[3] = `A file size is given in bytes. Find the amount of full Kbytes of this size (1&nbsp;K&nbsp;=&nbsp;1024&nbsp;bytes). Use the operator of integer division.`
	i.en[4] = `Two positive integers&nbsp;<i>A</i> and&nbsp;<i>B</i> are given (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). Segment of length&nbsp;<i>A</i> contains the greatest possible amount of inside segments of length&nbsp;<i>B</i> (without overlaps). Find the amount of segments&nbsp;<i>B</i> placed on the segment&nbsp;<i>A</i>. Use the operator of integer division.`
	i.en[5] = `Two positive integers&nbsp;<i>A</i> and&nbsp;<i>B</i> are given (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). Segment of length&nbsp;<i>A</i> contains the greatest possible amount of inside segments of length&nbsp;<i>B</i> (without overlaps). Find the length of unused part of the segment&nbsp;<i>A</i>. Use the operator of taking the remainder after integer division.`
	i.en[6] = `A two-digit integer is given. Output its left digit (a tens digit) and then its right digit (a ones digit). Use the operator of integer division for obtaining the tens digit  and the operator of taking remainder for obtaining the ones digit.`
	i.en[7] = `A two-digit integer is given. Find the sum and the product of its digits.`
	i.en[8] = `A two-digit integer is given. Output an integer obtained from the given one by exchange of its digits.`
	i.en[9] = `A three-digit integer is given. Using one operator of integer division find first digit of the given integer (a hundreds digit).`
	i.en[10] = `A three-digit integer is given. Output its last digit (a ones digit) and then its middle digit (a tens digit).`
	i.en[11] = `A three-digit integer is given. Find the sum and the product of its digits.`
	i.en[12] = `A three-digit integer is given. Output an integer obtained from the given one by reading it from right to left.`
	i.en[13] = `A three-digit integer is given. Output an integer obtained from the given one by moving its left digit to the right side.`
	i.en[14] = `A three-digit integer is given. Output an integer obtained from the given one by moving its right digit to the left side.`
	i.en[15] = `A three-digit integer is given. Output an integer obtained from the given one by exchange a tens digit and a hundreds digit (for example, 123 will be changed to&nbsp;213).`
	i.en[16] = `A three-digit integer is given. Output an integer obtained from the given one by exchange a ones digit and a tens digit (for example, 123 will be changed to 132).`
	i.en[17] = `An integer greater than 999 is given. Using one operator of integer division and one operator of taking the remainder find a hundreds digit of the given integer.`
	i.en[18] = `An integer greater than 999 is given. Using one operator of integer division and one operator of taking the remainder find a thousands digit of the given integer.`
	i.en[19] = `From the beginning of the day <i>N</i>&nbsp;seconds have passed (<i>N</i>&nbsp;is integer).  Find an amount of full minutes passed from the beginning of the day.`
	i.en[20] = `From the beginning of the day <i>N</i>&nbsp;seconds have passed (<i>N</i>&nbsp;is integer).  Find an amount of full hours passed from the beginning of the day.`
	i.en[21] = `From the beginning of the day <i>N</i>&nbsp;seconds have passed (<i>N</i>&nbsp;is integer).  Find an amount of seconds passed from the beginning of the last minute.`
	i.en[22] = `From the beginning of the day <i>N</i>&nbsp;seconds have passed (<i>N</i>&nbsp;is integer).  Find an amount of seconds passed from the beginning of the last hour.`
	i.en[23] = `From the beginning of the day <i>N</i>&nbsp;seconds have passed (<i>N</i>&nbsp;is integer).  Find an amount of full minutes passed from the beginning of the last hour.`
	i.en[24] = `Days of week are numbered as: 0 &#8212; Sunday, 1 &#8212; Monday, 2 &#8212; Tuesday,&nbsp;&#8230;, 6 &#8212; Saturday. An integer&nbsp;<i>K</i> in the range&nbsp;1 to&nbsp;365 is given. Find the number of day of week for <i>K</i>-th day of year provided that in this year January&nbsp;1 was Monday.`
	i.en[25] = `Days of week are numbered as: 0 &#8212; Sunday, 1 &#8212; Monday, 2 &#8212; Tuesday,&nbsp;&#8230;, 6 &#8212; Saturday. An integer&nbsp;<i>K</i> in the range&nbsp;1 to&nbsp;365 is given. Find the number of day of week for <i>K</i>-th day of year provided that in this year January&nbsp;1 was Thursday.`
	i.en[26] = `Days of week are numbered as: 1 &#8212; Monday, 2 &#8212; Tuesday,&nbsp;&#8230;, 6 &#8212; Saturday, 7 &#8212; Sunday. An integer&nbsp;<i>K</i> in the range&nbsp;1 to&nbsp;365 is given. Find the number of day of week for <i>K</i>-th day of year provided that in this year January&nbsp;1 was Tuesday.`
	i.en[27] = `Days of week are numbered as: 1 &#8212; Monday, 2 &#8212; Tuesday,&nbsp;&#8230;, 6 &#8212; Saturday, 7 &#8212; Sunday. An integer&nbsp;<i>K</i> in the range&nbsp;1 to&nbsp;365 is given. Find the number of day of week for <i>K</i>-th day of year provided that in this year January&nbsp;1 was Saturday.`
	i.en[28] = `Days of week are numbered as: 1 &#8212; Monday, 2 &#8212; Tuesday,&nbsp;&#8230;, 6 &#8212; Saturday, 7 &#8212; Sunday. An integer&nbsp;<i>K</i> in the range&nbsp;1 to&nbsp;365 and an integer&nbsp;<i>N</i> in the range&nbsp;1 to&nbsp;7 are given. Find the number of day of week for <i>K</i>-th day of year provided that in this year January&nbsp;1 was <i>N</i>-th day of week.`
	i.en[29] = `Three positive integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are given. A&nbsp;rectangle of the size&nbsp;<i>A</i>&nbsp;&#215;&nbsp;<i>B</i> contains the greatest possible amount of inside squares with the side length&nbsp;<i>C</i> (without overlaps). Find the amount of squares placed on the rectangle and the area of unused part of the rectangle.`
	i.en[30] = `Given a year (as a positive integer), find the respective number of the century. Note that, for example, 20th century began with the year 1901.`
}

func (i *iinteger) makeTj() {
	i.tj = make([]string, i.count)
	i.tj[1] = `Масофа L бо сантиметр дода шудааст. Амали тақсими бутунро истифода бурда, миқдори метрҳои пурраро дар он ёбед (1 метр = 100 см).`
	i.tj[2] = `Вазн М бо килограм дода шудааст. Амали тақсими бутунро истифода бурда, миқдори тоннаҳои пурраро дар он ёбед (1 тонна = 1000 кг).`
	i.tj[3] = `Андозаи файл бо байт дода шудааст. Амали тақсими бутунро истифода бурда, миқдори килобайтҳои пурраро, ки файли мазкур банд мекунад, ёбед (1 килобайт = 1024 байт).`
	i.tj[4] = `Ададҳои бутуни мусбат A ва B (A>B) дода шудаанд. Дар порчаи дарозиаш A миқдори калонтарини имконпазири порчаҳои дарозиашон B ҷойгиранд. Бо истифодабарии амали тақсими бутун миқдори порчаҳои B-ро, ки дар порчаи A ҷойгиранд, ёбед.`
	i.tj[5] = `Ададҳои бутуни мусбат A ва B (A>B) дода шудаанд. Дар порчаи дарозиаш A миқдори калонтарини имконпазири порчаҳои дарозиашон B ҷойгиранд. Бо истифодабарии амали гирифтани бақия аз тақсими бутун дарозии қисми бандНАбудаи порчаи A-ро ёбед.`
	i.tj[6] = `Адади дурақама дода шудааст. Дар аввал рақами чапи он (даҳӣ)-ро, сонӣ рақами рости он (воҳид)-ро хориҷ кунед. Барои ёфтани даҳӣ амали тақсими бутунро, барои ёфтани воҳид бошад - амали гирифтани бақия аз тақсимро истифода буред.`
	i.tj[7] = `Адади дурақама дода шудааст.  Сумма ва ҳосилизарби рақамҳои онро ёбед.`
	i.tj[8] = `Адади дурақама дода шудааст. Ададеро хориҷ кунед, ки дар натиҷаи ҷойивазкунии рақамҳои адади ибтидоӣ пайдо гардидааст.`
	i.tj[9] = `Адади серақама дода шудааст. Бо истифодабарии як амали тақсими бутун рақами аввалаи адади мазкур (садӣ)-ро хориҷ кунед.`
	i.tj[10] = `Адади серақама дода шудааст. Дар аввал рақами охирини он (воҳид)-ро, сонӣ рақами мобайнии он(даҳӣ)-ро хориҷ кунед.`
	i.tj[11] = `Адади серақама дода шудааст. Сумма ва ҳосилизарби рақамҳои онро ёбед.`
	i.tj[12] = `Адади серақама дода шудааст. Ададеро хориҷ кунед, ки дар натиҷаи хондани адади ибтидоӣ аз рост ба чап ҳосил шудааст.`
	i.tj[13] = `Адади серақама дода шудааст. Дар он рақами аз чап якӯмро хат зада, онро аз рост нависед. Адади ҳосилшударо хориҷ кунед.`
	i.tj[14] = `Адади серақама дода шудааст. Дар он рақами аз рост якӯмро хат зада, онро аз чап нависед. Адади ҳосилшударо хориҷ кунед.`
	i.tj[15] = `Адади серақама дода шудааст. Ададеро хориҷ кунед, ки дар натиҷаи ҷойивазкунии рақамҳои садӣ ва даҳии адади ибтидоӣ ҳосил шудааст (масалан, 123 мешавад 213).`
	i.tj[16] = `Адади серақама дода шудааст. Ададеро хориҷ кунед, ки дар натиҷаи ҷойивазкунии рақамҳои даҳӣ ва воҳидии адади ибтидоӣ ҳосил шудааст.`
	i.tj[17] = `Адади бутуни аз 999 калон дода шудааст. Бо истифодабарии як амали тақсими бутун ва як амали гирифтани бақия аз тақсим рақамеро ёбед, ки ба қатори садии ин адад мувофиқ меояд.`
	i.tj[18] = `Адади бутуни аз 999 калон дода шудааст. Бо истифодабарии як амали тақсими бутун ва як амали гирифтани бақия аз тақсим рақамеро ёбед, ки ба қатори ҳазории ин адад мувофиқ меояд.`
	i.tj[19] = `Аз аввали шабонарӯз N сония (N-бутун) гузаштааст. Миқдори дақиқаҳои пурраеро, ки аз аввали шабонарӯз гузаштааст, ёбед.`
	i.tj[20] = `Аз аввали шабонарӯз N сония (N-бутун) гузаштааст. Миқдори соатҳои пурраеро, ки аз аввали шабонарӯз гузаштааст, ёбед.`
	i.tj[21] = `Аз аввали шабонарӯз N сония (N-бутун) гузаштааст. Миқдори сонияҳоеро, ки аз аввали дақиқаи охирин гузаштааст, ёбед.`
	i.tj[22] = `Аз аввали шабонарӯз N сония (N-бутун) гузаштааст. Миқдори сонияҳоеро, ки аз аввали соати охирин гузаштааст, ёбед.`
	i.tj[23] = `Аз аввали шабонарӯз N сония (N-бутун) гузаштааст. Миқдори дақиқаҳои пурраро, ки аз аввали соати охирин гузаштааст, ёбед.`
	i.tj[24] = `Рӯзҳои ҳафта ба тариқи зайл рақамгузорӣ карда шудаанд: 0-якшанбе, 1-душанбе, 2-сешанбе, ..., 6-шанбе. Адади бутуни K дар фосилаи 1-365 хобанда дода шудааст. Рақами рӯзи ҳафтаро барои рӯзи K-ӯми сол муайян кунед, агар маълум бошад, ки дар ин сол 1 январ душанбе буд.`
	i.tj[25] = `Рӯзҳои ҳафта ба тариқи зайл рақамгузорӣ карда шудаанд: 0-якшанбе, 1-душанбе, 2-сешанбе, ..., 6-шанбе. Адади бутуни K дар фосилаи 1-365 хобанда дода шудааст. Рақами рӯзи ҳафтаро барои рӯзи K-ӯми сол муайян кунед, агар маълум бошад, ки дар ин сол 1 январ панҷшанбе буд.`
	i.tj[26] = `Рӯзҳои ҳафта ба тариқи зайл рақамгузорӣ карда шудаанд: 1-душанбе, 2-сешанбе, ..., 6-шанбе, 7-якшанбе. Адади бутуни K дар фосилаи 1-365 хобанда дода шудааст. Рақами рӯзи ҳафтаро барои рӯзи K-ӯми сол муайян кунед, агар маълум бошад, ки дар ин сол 1 январ сешанбе буд.`
	i.tj[27] = `Рӯзҳои ҳафта ба тариқи зайл рақамгузорӣ карда шудаанд: 1-душанбе, 2-сешанбе, ..., 6-шанбе, 7-якшанбе. Адади бутуни K дар фосилаи 1-365 хобанда дода шудааст. Рақами рӯзи ҳафтаро барои рӯзи K-ӯми сол муайян кунед, агар маълум бошад, ки дар ин сол 1 январ шанбе буд.`
	i.tj[28] = `Рӯзҳои ҳафта ба тариқи зайл рақамгузорӣ карда шудаанд: 1-душанбе, 2-сешанбе, ..., 6-шанбе, 7-якшанбе. Адади бутуни K дар фосилаи 1-365 хобанда ва адади бутуни N дар фосилаи 1-7 хобанда дода шудаанд. Рақами рӯзи ҳафтаро барои рӯзи K-ӯми сол муайян кунед, агар маълум бошад, ки дар ин сол 1 январ рӯзи ҳафта бо рақами N буд.`
	i.tj[29] = `Ададҳои бутуни мусбии A, B, C дода шудаанд. Дар росткунҷаи андозааш AxB миқдори калонтарини имконпазири квадратҳо бо андозаи C ҷой дода шудааст. Миқдори квадратҳоеро, ки дар росткунҷа ҷойгиранд, ҳамчунин масоҳати қисми бандНАбудаи росткунҷаро ёбед.`
	i.tj[30] = `Рақами соли дилхоҳ дода шудааст (адади бутуни мусбӣ). Рақами садсолаи ба он мувофиқро муайян кунед, бо назардошти он, ки масалан, оғози садсолаи 20 соли 1901 буд.`
}

func (i *iinteger) makeRu() {
	i.ru = make([]string, i.count)
	i.ru[1] = `Дано расстояние&nbsp;<i>L</i> в сантиметрах. Используя операцию деления нацело, найти количество полных метров в нем (1&nbsp;метр&nbsp;=&nbsp;100&nbsp;см).`
	i.ru[2] = `Дана масса&nbsp;<i>M</i> в килограммах. Используя операцию деления нацело, найти количество полных тонн в ней (1&nbsp;тонна&nbsp;=&nbsp;1000&nbsp;кг).`
	i.ru[3] = `Дан размер файла в байтах. Используя операцию деления нацело, найти количество полных килобайтов, которые занимает данный файл (1&nbsp;килобайт&nbsp;=&nbsp;1024&nbsp;байта).`
	i.ru[4] = `Даны целые положительные числа&nbsp;<i>A</i> и&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). На отрезке длины&nbsp;<i>A</i> размещено максимально возможное количество отрезков длины&nbsp;<i>B</i> (без наложений). Используя операцию деления нацело, найти количество отрезков&nbsp;<i>B</i>, размещенных на отрезке&nbsp;<i>A</i>.`
	i.ru[5] = `Даны целые положительные числа&nbsp;<i>A</i> и&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). На отрезке длины&nbsp;<i>A</i> размещено максимально возможное количество отрезков длины&nbsp;<i>B</i> (без наложений). Используя операцию взятия остатка от деления нацело, найти длину незанятой части отрезка&nbsp;<i>A</i>.`
	i.ru[6] = `Дано двузначное число. Вывести вначале его левую цифру (десятки), а затем &#8212; его правую цифру (единицы). Для нахождения десятков использовать операцию деления нацело, для нахождения единиц &#8212; операцию взятия остатка от деления.`
	i.ru[7] = `Дано двузначное число. Найти сумму и произведение его цифр.`
	i.ru[8] = `Дано двузначное число. Вывести число, полученное при перестановке цифр исходного числа.`
	i.ru[9] = `Дано трехзначное число. Используя одну операцию деления нацело, вывести первую цифру данного числа (сотни).`
	i.ru[10] = `Дано трехзначное число. Вывести вначале его последнюю цифру (единицы), а затем &#8212; его среднюю цифру (десятки).`
	i.ru[11] = `Дано трехзначное число. Найти сумму и произведение его цифр.`
	i.ru[12] = `Дано трехзначное число. Вывести число, полученное при прочтении исходного числа справа налево.`
	i.ru[13] = `Дано трехзначное число. В нем зачеркнули первую слева цифру и приписали ее справа. Вывести полученное число.`
	i.ru[14] = `Дано трехзначное число. В нем зачеркнули первую справа цифру и приписали ее слева. Вывести полученное число.`
	i.ru[15] = `Дано трехзначное число. Вывести число, полученное при перестановке цифр сотен и десятков исходного числа (например, 123 перейдет в 213).`
	i.ru[16] = `Дано трехзначное число. Вывести число, полученное при перестановке цифр десятков и единиц исходного числа (например, 123 перейдет в 132).`
	i.ru[17] = `Дано целое число, большее 999. Используя одну операцию деления нацело и одну операцию взятия остатка от деления, найти цифру, соответствующую разряду сотен в записи этого числа.`
	i.ru[18] = `Дано целое число, большее 999. Используя одну операцию деления нацело и одну операцию взятия остатка от деления, найти цифру, соответствующую разряду тысяч в записи этого числа.`
	i.ru[19] = `С начала суток прошло <i>N</i>&nbsp;секунд (<i>N</i> &#8212; целое). Найти количество полных минут, прошедших с начала суток.`
	i.ru[20] = `С начала суток прошло <i>N</i>&nbsp;секунд (<i>N</i> &#8212; целое). Найти количество полных часов, прошедших с начала суток.`
	i.ru[21] = `С начала суток прошло <i>N</i>&nbsp;секунд (<i>N</i> &#8212; целое). Найти количество секунд, прошедших с начала последней минуты.`
	i.ru[22] = `С начала суток прошло <i>N</i>&nbsp;секунд (<i>N</i> &#8212; целое). Найти количество секунд, прошедших с начала последнего часа.`
	i.ru[23] = `С начала суток прошло <i>N</i>&nbsp;секунд (<i>N</i> &#8212; целое). Найти количество полных минут, прошедших с начала последнего часа.`
	i.ru[24] = `Дни недели пронумерованы следующим образом: 0 &#8212; воскресенье, 1 &#8212; понедельник, 2 &#8212; вторник,&nbsp;&#8230;, 6 &#8212; суббота. Дано целое число&nbsp;<i>K</i>, лежащее в диапазоне 1&#8211;365. Определить номер дня недели для <i>K</i>-го дня года, если известно, что в этом году 1&nbsp;января было понедельником.`
	i.ru[25] = `Дни недели пронумерованы следующим образом: 0 &#8212; воскресенье, 1 &#8212; понедельник, 2 &#8212; вторник,&nbsp;&#8230;, 6 &#8212; суббота. Дано целое число&nbsp;<i>K</i>, лежащее в диапазоне 1&#8211;365. Определить номер дня недели для <i>K</i>-го дня года, если известно, что в этом году 1&nbsp;января было четвергом.`
	i.ru[26] = `Дни недели пронумерованы следующим образом: 1 &#8212; понедельник, 2 &#8212; вторник,&nbsp;&#8230;, 6 &#8212; суббота, 7 &#8212; воскресенье. Дано целое число&nbsp;<i>K</i>, лежащее в диапазоне 1&#8211;365. Определить номер дня недели для <i>K</i>-го дня года, если известно, что в этом году 1&nbsp;января было вторником.`
	i.ru[27] = `Дни недели пронумерованы следующим образом: 1 &#8212; понедельник, 2 &#8212; вторник,&nbsp;&#8230;, 6 &#8212; суббота, 7 &#8212; воскресенье. Дано целое число&nbsp;<i>K</i>, лежащее в диапазоне 1&#8211;365. Определить номер дня недели для <i>K</i>-го дня года, если известно, что в этом году 1&nbsp;января было субботой.`
	i.ru[28] = `Дни недели пронумерованы следующим образом: 1 &#8212; понедельник, 2 &#8212; вторник,&nbsp;&#8230;, 6 &#8212; суббота, 7 &#8212; воскресенье. Дано целое число&nbsp;<i>K</i>, лежащее в диапазоне 1&#8211;365, и целое число&nbsp;<i>N</i>, лежащее в диапазоне 1&#8211;7. Определить номер дня недели для <i>K</i>-го дня года, если известно, что в этом году 1&nbsp;января было днем недели с номером&nbsp;<i>N</i>.`
	i.ru[29] = `Даны целые положительные числа&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>. На прямоугольнике размера&nbsp;<i>A</i>&nbsp;&#215;&nbsp;<i>B</i> размещено максимально возможное количество квадратов со стороной&nbsp;<i>C</i> (без наложений). Найти количество квадратов, размещенных на прямоугольнике, а также площадь незанятой части прямоугольника.`
	i.ru[30] = `Дан номер некоторого года (целое положительное число). Определить соответствующий ему номер столетия, учитывая, что, к примеру, началом 20&nbsp;столетия был 1901 год.`
}
