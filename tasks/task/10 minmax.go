package task

type mminmax struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewMinmaxTask() Task {
	m := &mminmax{
		count: 31,
		name:  "Minmax",
	}
	m.makeEn()
	m.makeTj()
	m.makeRu()
	return m
}

func (m *mminmax) Count() int   { return m.count }
func (m *mminmax) Name() string { return m.name }
func (m *mminmax) En() []string { return m.en }
func (m *mminmax) Tj() []string { return m.tj }
func (m *mminmax) Ru() []string { return m.ru }

func (m *mminmax) makeEn() {
	m.en = make([]string, m.count)
	m.en[1] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. Find the minimal element and the maximal element of the sequence (that is, elements with the minimal value and the maximal value respectively).`
	m.en[2] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;rectangles are given. Each rectangle is defined by a pair of its sides (<i>a</i>,&nbsp;<i>b</i>). Find the rectangle with the minimal area and output the value of its area.`
	m.en[3] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;rectangles are given. Each rectangle is defined by a pair of its sides (<i>a</i>,&nbsp;<i>b</i>). Find the rectangle with the maximal perimeter and output the value of its perimeter.`
	m.en[4] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. Find the order number of the minimal element of the sequence.`
	m.en[5] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;pairs of real numbers (<i>m</i>,&nbsp;<i>v</i>) are given. Each pair of given numbers contains the weight&nbsp;<i>m</i> and the volume&nbsp;<i>v</i> of a detail that is made of some homogeneous material. Output the order number of a detail that is made of the material with maximal density. Also output the corresponding density. Note that the density&nbsp;<i>P</i> can be found by formula <i>P</i>&nbsp;=&nbsp;<i>m</i>/<i>v</i>.`
	m.en[6] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order numbers of the first minimal element and the last maximal element of the sequence.`
	m.en[7] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order numbers of the first maximal element and the last minimal element of the sequence.`
	m.en[8] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order numbers of the first and the last minimal elements of the sequence.`
	m.en[9] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order numbers of the first and the last maximal elements of the sequence.`
	m.en[10] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order number of the first <i>extremal</i> (that is, minimal or maximal) element of the sequence.`
	m.en[11] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the order number of the last <i>extremal</i> (that is, minimal or maximal) element of the sequence.`
	m.en[12] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. Output the minimal positive number contained in the sequence. If the sequence does not contain positive numbers then output&nbsp;0.`
	m.en[13] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output the order number of the first maximal odd number contained in the sequence. If the sequence does not contain odd numbers then output&nbsp;0.`
	m.en[14] = `A positive real number&nbsp;<i>B</i> and a sequence of 10&nbsp;real numbers are given. Find the minimum among elements that are greater than&nbsp;<i>B</i> and output its value and its order number. If the sequence does not contain elements greater than&nbsp;<i>B</i> then output&nbsp;0 twice (the first zero as a real number, the second zero as an integer).`
	m.en[15] = `Two real numbers&nbsp;<i>B</i>, <i>C</i> (0&nbsp;&lt;&nbsp;<i>B</i>&nbsp;&lt;&nbsp;<i>C</i>) and a sequence of 10&nbsp;real numbers are given. Find the maximum among elements that are in the interval (<i>B</i>,&nbsp;<i>C</i>) and output its value and its order number. If the sequence does not contain elements in the interval (<i>B</i>,&nbsp;<i>C</i>) then output&nbsp;0 twice (the first zero as a real number, the second zero as an integer).`
	m.en[16] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the amount of the elements that are located before the first minimal element.`
	m.en[17] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the amount of the elements that are located after the last maximal element.`
	m.en[18] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the amount of the elements that are located between the first and the last maximal element. If the sequence contains the unique maximal element then output&nbsp;0.`
	m.en[19] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the amount of the minimal elements of the sequence.`
	m.en[20] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the total amount of all <i>extremal</i> (that is, minimal or maximal) elements of the sequence.`
	m.en[21] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;2) and a sequence of <i>N</i>&nbsp;real numbers are given. The elements of the sequence represent some experimental data. Find the average of the experimental data provided that the minimal and maximal values must be ignored.`
	m.en[22] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;2) and a sequence of <i>N</i>&nbsp;real numbers are given. Find two smallest elements of the sequence and output their values in ascending order.`
	m.en[23] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;3) and a sequence of <i>N</i>&nbsp;real numbers are given. Find three greatest elements of the sequence and output their values in descending order.`
	m.en[24] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of <i>N</i>&nbsp;real numbers are given. Find the maximal sum of two adjacent elements of the sequence.`
	m.en[25] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of <i>N</i>&nbsp;real numbers are given. Find two adjacent elements that have the minimal product of their values and output their order numbers in ascending order.`
	m.en[26] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output the maximal amount of successive elements whose values are even numbers. If the sequence does not contain even numbers then output&nbsp;0.`
	m.en[27] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. The sequence contains elements of values&nbsp;0 and&nbsp;1 only. Find the longest subsequence of the successive elements with equal values, and output the order number of its initial element and the amount of its elements. If there are several such subsequences then output the order number of the first one.`
	m.en[28] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. The sequence contains elements of values&nbsp;0 and&nbsp;1 only. Find the longest subsequence of the successive elements of value&nbsp;1, and output the order number of its initial element and the amount of its elements. If there are several such subsequences then output the order number of the first one. If the sequence does not contain elements of value&nbsp;1 then output&nbsp;0 twice.`
	m.en[29] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the maximal amount of the successive elements with the minimal value.`
	m.en[30] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Find the minimal amount of the successive elements with the maximal value.`
}

func (m *mminmax) makeTj() {
	m.tj = make([]string, m.count)
	m.tj[1] = `Адади бутуни N ва маҷмӯъ аз N ададҳо дода шудаанд. Элементҳои хурдтарин ва калонтарини маҷмӯи мазкурро ёбед ва бо тартиби нишондодашуда хориҷ кунед.`
	m.tj[2] = `Адади бутуни N ва маҷмӯъ аз N росткунҷаҳо бо тарафҳои додашуда — ҷуфтҳои ададҳо (a, b) дода шудаанд. Масоҳати хурдтарини росткунҷаро аз маҷмӯи мазкур ёбед.`
	m.tj[3] = `Адади бутуни N ва маҷмӯъ аз N росткунҷаҳо бо тарафҳои додашуда - ҷуфтҳои ададҳо (a, b) дода шудаанд. Периметри калонтарини росткунҷаро аз маҷмӯи мазкур ёбед.`
	m.tj[4] = `Адади бутуни N ва маҷмӯъ аз N ададҳо дода шудаанд. Рақами элементи хурдтаринро аз маҷмӯи мазкур ёбед.`
	m.tj[5] = `Адади бутуни N ва маҷмӯъ аз N ҷуфти ададҳо (m, v) — маълумот дар бораи вазн m ва ҳаҷми v ҷузъиёт, ки аз материалҳои гуногун тайёр карда шудаанд, дода шудаанд. Рақами ҷузъиётеро, ки аз материали дорои зичии калонтарин тайёр карда шудааст, ҳамчунин бузургии ин зичии калонтаринро хориҷ кунед. Зичӣ P аз рӯи формулаи P = m/v ҳисоб карда мешавад.`
	m.tj[6] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақами элементи якӯми хурдтарин ва рақами элементи охирини калонтаринро аз маҷмӯи мазкур ёбед ва онҳоро бо тартиби нишондодашуда хориҷ кунед.`
	m.tj[7] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақами элементи якӯми калонтарин ва рақами элементи охирини хурдтаринро аз маҷмӯи мазкур ёбед ва онҳоро бо тартиби нишондодашуда хориҷ кунед.`
	m.tj[8] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақамҳои элементҳои хурдтарини якӯм ва охиринро аз маҷмӯи мазкур ёбед ва онҳоро бо тартиби нишондодашуда хориҷ кунед.`
	m.tj[9] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақамҳои элементҳои калонтарини якӯм ва охиринро аз маҷмӯи мазкур ёбед ва онҳоро бо тартиби нишондодашуда хориҷ кунед.`
	m.tj[10] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақами элементи якӯми экстремалӣ (яъне калонтарин ё хурдтарин)-ро аз маҷмӯи мазкур ёбед.`
	m.tj[11] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Рақами элементи охирини экстремалӣ (яъне калонтарин ё хурдтарин)-ро аз маҷмӯи мазкур ёбед.`
	m.tj[12] = `Адади бутуни N ва маҷмӯъ аз N ададҳо дода шудаанд. Адади мусбии хурдтаринро аз маҷмӯи мазкур ёбед. Агар дар маҷмӯъ адади мусбӣ набошад, пас 0(нул)-ро хориҷ кунед.`
	m.tj[13] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Рақами адади тоқи калонтарини аввалинро аз маҷмӯи мазкур ёбед. Агар адади тоқ дар маҷмӯъ набошад, пас 0(нул)-ро хориҷ кунед.`
	m.tj[14] = `Адади B(>0) ва маҷмӯъ аз даҳ ададҳо дода шудаанд. Аз дохили он элементҳои маҷмӯъ, ки аз B  калон ҳастанд, элементи хурдтаринашро, ҳамчунин рақами тартибии онро хориҷ кунед. Агар дар маҷмӯъ адади аз B калон набошад, пас ду маротиба 0(нул)-ро хориҷ кунед.`
	m.tj[15] = `Ададҳои B, C (В>0, C>В) ва маҷмӯъ аз даҳ ададҳо дода шудаанд. Аз дохили он элементҳои маҷмӯъ, ки дар фосилаи (B,C) мехобанд, калонтаринашро, ҳамчунин рақами тартибии онро хориҷ кунед. Агар дар маҷмӯъ адади талабкардашуда набошад, пас ду маротиба 0(нул)-ро хориҷ кунед.`
	m.tj[16] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори элементҳоеро, ки пеш аз элементи хурдтарини аввалин ҷойгир шудаанд, ёбед.`
	m.tj[17] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори элементҳоеро, ки пас аз элементи калонтарини охирин ҷойгир шудаанд, ёбед.`
	m.tj[18] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори элементҳоеро, ки дар байни элементҳои калонтарини аввалин ва охирин ҷойгир шудаанд, ёбед. Агар дар маҷмӯъ як то элементи калонтарин мавҷуд бошад, пас 0(нул)-ро хориҷ кунед.`
	m.tj[19] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори элементҳои хурдтаринро аз маҷмӯи мазкур ёбед.`
	m.tj[20] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори умумии элементҳои экстремалӣ (яъне калонтарин ва хурдтарин)-ро аз маҷмӯи мазкур ёбед.`
	m.tj[21] = `Адади бутуни N(>2) ва маҷмӯъ аз N ададҳо - қиматҳои баъзе бузургиҳо, ки дар N таҷрибаҳо ҳосил гардидаанд, дода шудаанд. Қимати миёнаи ин бузургиро ёбед. Ҳангоми ҳисобкунии қимати миёна қиматҳои калонтарин ва хурдтарини дар маҷмӯъ мавҷударо ба назар нагиред.`
	m.tj[22] = `Адади бутуни N(>2) ва маҷмӯъ аз N ададҳо дода шудаанд. Ду то элементҳои хурдтаринро аз маҷмӯи мазкур ёбед ва ин элементҳоро бо тартиби афзуншавии қиматҳояшон хориҷ кунед.`
	m.tj[23] = `Адади бутуни N(>3) ва маҷмӯъ аз N ададҳо дода шудаанд. Се то элементҳои калонтаринро аз маҷмӯи мазкур ёбед ва ин элементҳоро бо тартиби камшавии қиматҳояшон хориҷ кунед.`
	m.tj[24] = `Адади бутуни N(>1) ва маҷмӯъ аз N ададҳо дода шудаанд. Суммаи калонтарини ду ададҳои ҳамсояро аз маҷмӯи мазкур ёбед.`
	m.tj[25] = `Адади бутуни N(>1) ва маҷмӯъ аз N ададҳо дода шудаанд. Рақамҳои ду ададҳои ҳамсояро аз маҷмӯи мазкур ёбед, ки ҳосилизарби онҳо хурдтарин аст ва дар аввал рақами хурд ва пасон рақами калонро хориҷ кунед.`
	m.tj[26] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори калонтарини ададҳои ҷуфти пайиҳамояндаро дар маҷмӯъ ёбед. Агар дар маҷмӯъ адади ҷуфт набошад, пас 0(нул)-ро хориҷ кунед.`
	m.tj[27] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун, ки дорои фақат нулҳо ва воҳидҳо аст, дода шудаанд. Рақами элементеро, ки аз он сар карда пайдарпайии дарозтарини ададҳои якхела сар мешавад ва миқдори элементҳоро дар ин пайдарпайӣ ёбед. Агар чунин пайдарпайиҳо бисёр бошанд, пас рақами якӯмини онҳоро хориҷ кунед.`
	m.tj[28] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун, ки дорои фақат нулҳо ва воҳидҳо ҳастанд, дода шудаанд. Рақами элементеро, ки аз он сар карда пайдарпайии дарозтарини воҳидҳо сар мешавад, ва миқдори элементҳоро дар ин пайдарпайӣ ёбед. Агар чунин пайдарпайиҳо бисёр бошанд, пас рақами охирини онҳоро хориҷ кунед. Агар дар маҷмӯи додашуда воҳид(рақами як) набошад, пас ду маротиба 0(нул)-ро хориҷ кунед.`
	m.tj[29] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори калонтарини элементҳои хурдтарини пайиҳамояндаро аз маҷмӯи мазкур ёбед.`
	m.tj[30] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Миқдори хурдтарини элементҳои калонтарини пайиҳамояндаро аз маҷмӯи мазкур ёбед.`
}

func (m *mminmax) makeRu() {
	m.ru = make([]string, m.count)
	m.ru[1] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;чисел. Найти минимальный и максимальный из элементов данного набора и вывести их в указанном порядке.`
	m.ru[2] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;прямоугольников, заданных своими сторонами &#8212; парами чисел (<i>a</i>,&nbsp;<i>b</i>). Найти минимальную площадь прямоугольника из данного набора.`
	m.ru[3] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;прямоугольников, заданных своими сторонами &#8212; парами чисел (<i>a</i>,&nbsp;<i>b</i>). Найти максимальный периметр прямоугольника из данного набора.`
	m.ru[4] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;чисел. Найти номер минимального элемента из данного набора.`
	m.ru[5] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;пар чисел (<i>m</i>,&nbsp;<i>v</i>) &#8212; данные о массе&nbsp;<i>m</i> и объеме&nbsp;<i>v</i> деталей, изготовленных из различных материалов. Вывести номер детали, изготовленной из материала максимальной плотности, а также величину этой максимальной плотности. Плотность&nbsp;<i>P</i> вычисляется по формуле <i>P</i>&nbsp;=&nbsp;<i>m</i>/<i>v</i>.`
	m.ru[6] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номера первого минимального и последнего максимального элемента из данного набора и вывести их в указанном порядке.`
	m.ru[7] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номера первого максимального и последнего минимального элемента из данного набора и вывести их в указанном порядке.`
	m.ru[8] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номера первого и последнего минимального элемента из данного набора и вывести их в указанном порядке.`
	m.ru[9] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номера первого и последнего максимального элемента из данного набора и вывести их в указанном порядке.`
	m.ru[10] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номер первого <i>экстремального</i> (т.&nbsp;е. минимального или максимального) элемента из данного набора.`
	m.ru[11] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номер последнего <i>экстремального</i> (т.&nbsp;е. минимального или максимального) элемента из данного набора.`
	m.ru[12] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;чисел. Найти минимальное положительное число из данного набора. Если положительные числа в наборе отсутствуют, то вывести&nbsp;0.`
	m.ru[13] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти номер первого максимального нечетного числа из данного набора. Если нечетные числа в наборе отсутствуют, то вывести&nbsp;0.`
	m.ru[14] = `Дано число&nbsp;<i>B</i> (&gt;&nbsp;0) и набор из десяти чисел. Вывести минимальный из тех элементов набора, которые больше&nbsp;<i>B</i>, а также его номер. Если чисел, больших&nbsp;<i>B</i>, в наборе нет, то дважды вывести&nbsp;0.`
	m.ru[15] = `Даны числа&nbsp;<i>B</i>, <i>C</i> (0&nbsp;&lt;&nbsp;<i>B</i>&nbsp;&lt;&nbsp;<i>C</i>) и набор из десяти чисел. Вывести максимальный из элементов набора, содержащихся в интервале (<i>B</i>,&nbsp;<i>C</i>), и его номер. Если требуемые числа в наборе отсутствуют, то дважды вывести&nbsp;0.`
	m.ru[16] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти количество элементов, расположенных перед первым минимальным элементом.`
	m.ru[17] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти количество элементов, расположенных после последнего максимального элемента.`
	m.ru[18] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти количество элементов, содержащихся между первым и последним максимальным элементом. Если в наборе имеется единственный максимальный элемент, то вывести&nbsp;0.`
	m.ru[19] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти количество минимальных элементов из данного набора.`
	m.ru[20] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти общее количество <i>экстремальных</i> (т.&nbsp;е. минимальных и максимальных) элементов из данного набора.`
	m.ru[21] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;2) и набор из <i>N</i>&nbsp;чисел &#8212; значений некоторой величины, полученных в <i>N</i>&nbsp;опытах. Найти среднее значение этой величины. При вычислении среднего значения не учитывать минимальное и максимальное из имеющихся в наборе значений.`
	m.ru[22] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;2) и набор из <i>N</i>&nbsp;чисел. Найти два наименьших элемента из данного набора и вывести эти элементы в порядке возрастания их значений.`
	m.ru[23] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;3) и набор из <i>N</i>&nbsp;чисел. Найти три наибольших элемента из данного набора и вывести эти элементы в порядке убывания их значений.`
	m.ru[24] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;чисел. Найти максимальную сумму двух соседних чисел из данного набора.`
	m.ru[25] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;чисел. Найти номера двух соседних чисел из данного набора, произведение которых является минимальным, и вывести вначале меньший, а затем больший номер.`
	m.ru[26] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти максимальное количество четных чисел в наборе, идущих подряд. Если четные числа в наборе отсутствуют, то вывести&nbsp;0.`
	m.ru[27] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел, содержащий только нули и единицы. Найти номер элемента, с которого начинается самая длинная последовательность одинаковых чисел, и количество элементов в этой последовательности. Если таких последовательностей несколько, то вывести номер первой из них.`
	m.ru[28] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел, содержащий только нули и единицы. Найти номер элемента, с которого начинается самая длинная последовательность единиц, и количество элементов в этой последовательности. Если таких последовательностей несколько, то вывести номер последней из них. Если единицы в исходном наборе отсутствуют, то дважды вывести&nbsp;0.`
	m.ru[29] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти максимальное количество подряд идущих минимальных элементов из данного набора.`
	m.ru[30] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Найти минимальное количество подряд идущих максимальных элементов из данного набора.`
}
