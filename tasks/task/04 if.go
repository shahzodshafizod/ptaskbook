package task

type iif struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewIfTask() Task {
	i := &iif{
		count: 31,
		name:  "If",
	}
	i.makeEn()
	i.makeTj()
	i.makeRu()
	return i
}

func (i *iif) Count() int   { return i.count }
func (i *iif) Name() string { return i.name }
func (i *iif) En() []string { return i.en }
func (i *iif) Tj() []string { return i.tj }
func (i *iif) Ru() []string { return i.ru }

func (i *iif) makeEn() {
	i.en = make([]string, i.count)
	i.en[1] = `An integer is given. If the integer is positive then decrease it by&nbsp;8, otherwise do not change it. Output the obtained integer.`
	i.en[2] = `An integer is given. If the integer is positive then decrease it by&nbsp;8, otherwise increase it by&nbsp;6. Output the obtained integer.`
	i.en[3] = `An integer is given. If the integer is positive then decrease it by&nbsp;8, if the integer is negative then increase it by&nbsp;6, if the integer equals&nbsp;0 then change it to&nbsp;10. Output the obtained integer.`
	i.en[4] = `Three integers are given. Find the amount of positive integers in the input data.`
	i.en[5] = `Three integers are given. Find the amount of positive and amount of negative integers in the input data.`
	i.en[6] = `Given two real numbers, output the larger value of them.`
	i.en[7] = `Given two real numbers, output the order number of the smaller of them.`
	i.en[8] = `Given two real numbers, output the larger value and then the smaller value of them.`
	i.en[9] = `The values of two real variables&nbsp;<i>A</i> and&nbsp;<i>B</i> are given. Redistribute the values so that&nbsp;<i>A</i> and&nbsp;<i>B</i> have the smaller and the larger value respectively. Output the new values of the variables&nbsp;<i>A</i> and&nbsp;<i>B</i>.`
	i.en[10] = `The values of two integer variables&nbsp;<i>A</i> and&nbsp;<i>B</i> are given. If the values are not equal then assign the sum of given values to each variable, otherwise assign zero value to each variable. Output the new values of the variables&nbsp;<i>A</i> and&nbsp;<i>B</i>.`
	i.en[11] = `The values of two integer variables&nbsp;<i>A</i> and&nbsp;<i>B</i> are given. If the values are not equal then assign the larger value to each variable, otherwise assign zero value to each variable. Output the new values of the variables&nbsp;<i>A</i> and&nbsp;<i>B</i>.`
	i.en[12] = `Given three real numbers, output the minimal value of them.`
	i.en[13] = `Given three real numbers, output the value between the minimum and the maximum.`
	i.en[14] = `Given three real numbers, output the minimal value and then the maximal value.`
	i.en[15] = `Given three real numbers, output the sum of two largest values.`
	i.en[16] = `The values of three real variables&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are given. If the values are in ascending order then double them, otherwise replace the value of each variable by its opposite value. Output the new values of the variables&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>.`
	i.en[17] = `The values of three real variables&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are given. If the values are in ascending or descending order then double them, otherwise replace the value of each variable by its opposite value. Output the new values of the variables&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>.`
	i.en[18] = `Three integers are given. One of them differs from two other equal integers. Output the order number of the integer that differs from the others.`
	i.en[19] = `Four integers are given. One of them differs from three other equal integers. Output the order number of the integer that differs from the others.`
	i.en[20] = `Three points&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> on the real axis are given. Determine whether&nbsp;<i>B</i> or&nbsp;<i>C</i> is closer to&nbsp;<i>A</i>. Output the nearest point and its distance from&nbsp;<i>A</i>.`
	i.en[21] = `Integer coordinates of a point in the coordinate plane are given. If the point coincides with the origin of coordinates then output&nbsp;0, otherwise if the point lies on the <i>x</i>-axis or <i>y</i>-axis then output&nbsp;1 or&nbsp;2 respectively. If the point does not lie on the coordinate axes then output&nbsp;3.`
	i.en[22] = `Given coordinates of a point that does not lie on the coordinate axes, find the number of a coordinate quarter containing the point.`
	i.en[23] = `Given integer coordinates of three vertices of a rectangle whose sides are parallel to coordinate axes, find the coordinates of the fourth vertex of the rectangle.`
	i.en[24] = `Given a real independent variable&nbsp;<i>x</i>, find the value of a real function&nbsp;<i>f</i> defined as: <table align="center"><tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>2&#183;sin(<i>x</i>),<td align=left>if <i>x</i>&nbsp;&gt;&nbsp;0,         <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>6&nbsp;&#8722;&nbsp;<i>x</i>,<td align=left>if <i>x</i>&nbsp;&#8804;&nbsp;0.</table>`
	i.en[25] = `Given an integer independent variable&nbsp;<i>x</i>, find the value of an integer function&nbsp;<i>f</i> defined as: <table align="center"><tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>2&#183;<i>x</i>,<td align=left>if <i>x</i>&nbsp;&lt;&nbsp;&#8722;2 or <i>x</i>&nbsp;&gt;&nbsp;2, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;3&#183;<i>x</i><td align=left>otherwise.</table>`
	i.en[26] = `Given a real independent variable&nbsp;<i>x</i>, find the value of a real function&nbsp;<i>f</i> defined as: <table align="center"><tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;<i>x</i>,<td align=left>if <i>x</i>&nbsp;&#8804;&nbsp;0, <tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right><i>x</i><sup>2</sup>,<td align=left>if 0&nbsp;&lt;&nbsp;<i>x</i>&nbsp;&lt;&nbsp;2, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>4,<td align=left>if <i>x</i>&nbsp;&#8805;&nbsp;2.</table>`
	i.en[27] = `Given a real independent variable&nbsp;<i>x</i>, find the value of an integer function&nbsp;<i>f</i> defined as: <table align="center"><tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>0,<td align=left>if <i>x</i>&nbsp;&lt;&nbsp;0, <tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>1,<td align=left>if <i>x</i> belongs to [0,&nbsp;1), [2,&nbsp;3),&nbsp;&#8230;, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;1,<td align=left>if <i>x</i> belongs to [1,&nbsp;2), [3,&nbsp;4),&nbsp;&#8230; .</table>`
	i.en[28] = `Given a number of year (as a positive integer), find the amount of days in the year. Note that the length of year is 365&nbsp;days for an ordinary year and 366&nbsp;days for a leap year. A <i>leap year</i> is every year whose number is divisible by&nbsp;4, as&nbsp;1980, except centenary years that are not divisible by&nbsp;400 (for example, 1300 and&nbsp;1900 are ordinary years, 1200 and&nbsp;2000 are leap years).`
	i.en[29] = `Given an integer, output its description string as: &#34;negative even number&#34;, &#34;zero number&#34;, &#34;positive odd number&#34;, etc.`
	i.en[30] = `An integer in the range&nbsp;1 to&nbsp;999 is given. Output its description string as: &#34;two-digit even number&#34;, &#34;three-digit odd number&#34;, etc.`
}

func (i *iif) makeTj() {
	i.tj = make([]string, i.count)
	i.tj[1] = `Адади бутун дода шудааст. Агар он мусбӣ бошад, аз он 8-ро тарҳ кунед; дар ҳолати акс онро тағйир надиҳед. Адади ҳосилшударо хориҷ кунед.`
	i.tj[2] = `Адади бутун дода шудааст. Агар он мусбӣ бошад, аз он 8-ро тарҳ кунед; дар ҳолати акс ба он 8-ро ҳамроҳ кунед. Адади ҳосилшударо хориҷ кунед.`
	i.tj[3] = `Адади бутун дода шудааст. Агар он мусбӣ бошад, аз он 8-ро тарҳ кунед; агар манфӣ бошад, пас ба он 8-ро ҳамроҳ кунед; агар нулӣ бошад, пас онро ба 10 иваз кунед. Адади ҳосилшударо хориҷ кунед.`
	i.tj[4] = `Се ададҳои бутун дода шудаанд. Миқдори ададҳои мусбиро дар маҷмӯаи ибтидоӣ ёбед.`
	i.tj[5] = `Се ададҳои бутун дода шудаанд. Миқдори ададҳои мусбӣ ва миқдори ададҳои манфиро дар маҷмӯаи ибтидоӣ ёбед.`
	i.tj[6] = `Ду ададҳо дода шудаанд. Калонтарини онҳоро ёбед.`
	i.tj[7] = `Ду ададҳо дода шудаанд. Рақами тартибии хурдтарини онҳоро ёбед.`
	i.tj[8] = `Ду ададҳо дода шудаанд. Дар аввал калонтарин ва сонӣ хурдтарини онҳоро ёбед.`
	i.tj[9] = `Ду тағйирёбандаҳои типи ҳақиқӣ дода шудаанд: A, B. Қиматҳои тағйирёбандаҳои мазкурро чунон ҷобаҷо кунед, ки дар A қимати хурдтарин ва дар B бошад — қимати калонтарин ҷоргир шавад. Қиматҳои нави тағйирёбандаҳои A ва B-ро хориҷ кунед.`
	i.tj[10] = `Ду тағйирёбандаҳои типи бутун дода шудаанд: A ва B. Агар қиматҳои онҳо нобаробар бошанд, пас ба ҳар як тағйирёбанда суммаи ин қиматҳоро бахшед, аммо агар баробар бошанд, пас ба тағйирёбандаҳо қиматҳои нулиро бахшед. Қиматҳои нави тағйирёбандаҳои A ва B-ро хориҷ кунед.`
	i.tj[11] = `Ду тағйирёбандаҳои типи бутун дода шудаанд: A ва B. Агар қиматҳои онҳо нобаробар бошанд, пас ба ҳар як тағйирёбанда калонтарини ин қиматҳоро бахшед, аммо агар баробар бошанд, пас ба тағйирёбандаҳо қиматҳои нулиро бахшед. Қиматҳои нави тағйирёбандаҳои A ва B-ро хориҷ кунед.`
	i.tj[12] = `Се ададҳо дода шудаанд. Хурдтарини онҳоро хориҷ кунед.`
	i.tj[13] = `Се ададҳо дода шудаанд. Қимати мобайниро аз байни онҳо ёбед (яъне ададеро, ки дар байни қиматҳои калонтарин ва хурдтарин ҷойгир аст).`
	i.tj[14] = `Се ададҳо дода шудаанд. Дар аввал адади хурдтарин ва сонӣ адади калонтаринро хориҷ кунед.`
	i.tj[15] = `Се ададҳо дода шудаанд. Суммаи ду ададҳои калонтаринро аз байни онҳо ёбед.`
	i.tj[16] = `Се тағйирёбандаҳои типи ҳақиқӣ дода шудаанд: A, B, C. Агар қиматҳои онҳо аз рӯи афзуншавӣ ҷобаҷо карда шуда бошанд, пас онҳоро ба ду зарб кунед; дар ҳолати акс қимати ҳар як тағйирёбандаро ба муқобилаломаташ иваз кунед. Қиматҳои нави тағйирёбандаҳои A, B, C-ро хориҷ кунед.`
	i.tj[17] = `Се тағйирёбандаҳои типи ҳақиқӣ дода шудаанд: A, B, C. Агар қиматҳои онҳо аз рӯи афзуншавӣ ва ё камшавӣ ҷобаҷо карда шуда бошанд, пас онҳоро ба ду зарб кунед; дар холати акс қимати ҳар як тағйирёбандаро ба муқобилаломаташ иваз кунед. Қиматҳои нави тағйирёбандаҳои A, B, C-ро хориҷ кунед.`
	i.tj[18] = `Се ададҳои бутун дода шудаанд, ки яке аз онҳо аз ду ададҳои дигарии байни ҳам баробар фарқ мекунад. Рақами тартибии адади фарқкунандаро муайян кунед.`
	i.tj[19] = `Чор ададҳои бутун дода шудаанд, ки яке аз онҳо аз се дадҳои дигарии байни ҳам баробар фарқ мекунад. Рақами тартибии адади фарқкунандаро муайян кунед.`
	i.tj[20] = `Дар тири ададӣ се нуқтаҳои: A, B, C ҷойгир шудаанд. Муайян кунед, ки кадоме аз ду нуқтаҳои охирӣ (B ё C) ба нуқтаи А наздиктар ҷой гирифтааст. Ин нуқта ва масофаи онро аз нуқтаи А хориҷ кунед.`
	i.tj[21] = `Координатаҳои типи бутуни нуқта дар ҳамворӣ дода шудаанд. Агар нуқта бо оғози координата мувофиқ ояд, пас 0(нул)-ро хориҷ кунед. Агар нуқта бо оғози координата мувофиқат накунаду, аммо дар тирҳои OX ё OY хобад, пас мувофиқан 1 ва ё 2-ро хориҷ кунед. Агар нуқта дар тирҳои координатавӣ нахобад, пас 3-ро хориҷ кунед.`
	i.tj[22] = `Координатаҳои нуқта, ки дар тирҳои координатавии OX ва OY намехобанд, дода шудаанд. Рақами чоряки координатавиеро муайян кунед, ки дар он ин нуқта ҷойгир аст.`
	i.tj[23] = `Координатаҳои типи бутуни се қуллаҳои росткунҷа дода шудаанд, ки тарафҳои он ба тирҳои координатавӣ параллел мебошанд. Координатаҳои қуллаи чорӯми он ёфта шавад.`
	i.tj[24] = `Барои адади ҳақиқии додашуда x қимати функсияи зерин f ёфта шавад, ки қиматҳои ҳақиқӣ қабул мекунад:<br>f(x)=2*sin(x), агар x>0<br>f(x)=6-x, агар x<=0.`
	i.tj[25] = `Барои адади бутуни додашуда x қимати функсияи зерин f ёфта шавад, ки қиматҳои типи бутунро қабул мекунад:<br>f(x)=2*x, агар x<-2 ё x>2<br>f(x)=-3*x дар ҳолати акс.`
	i.tj[26] = `Барои адади ҳақиқии додашуда x қимати функсияи зерин f ёфта шавад, ки қиматҳои ҳақиқӣ қабул мекунад:<br>f(x)=-x, агар x≤0;<br>f(x)=x<sup>2</sup>, агар 0&lt;x&lt;2;<br>f(x)=4, агар x≥2.`
	i.tj[27] = `Барои адади ҳақиқии додашуда x қимати функсияи зерин f ёфта шавад, ки қиматҳои типи бутунро қабул мекунад:<br>f(x)=0, агар x<0;<br>f(x)=1, агар x дар [0,1), [2,3), ...;<br>f(x)=-1, агар x дар [1,2), [3,4), ... .`
	i.tj[28] = `Рақами сол дода шудааст (адади бутуни мусбӣ). Миқдори рӯзҳоро дар ин сол муайян кунед, бо назардошти он, ки соли муқаррарӣ дорои 365 рӯз аст ва соли қабисавӣ - дорои 366 рӯз. Сол қабисавӣ ҳисобида мешавад, агар бебақия ба 4 тақсим шавад, ба истиснои он солҳое, ки онҳо бебақия ба 100 тақсим шуда, ба 400 тақсим намешаванд (масалан, солҳои 300, 1300 ва 1900 солҳои қабисавӣ нестанд, аммо 1200 ва 2000 бошанд - солҳои қабисавӣ ҳастанд).`
	i.tj[29] = `Адади бутун дода шудааст. Сатр-эзоҳи онро дар намуди "адади ҷуфти манфӣ", "адади нулӣ", "адади тоқи мусбӣ" ва ғ. хориҷ кунед.`
	i.tj[30] = `Адади бутуни дар фосилаи 1-999 хобанда дода шудааст. Сатр-эзоҳи онро дар намуди "адади дурақамаи ҷуфт", "адади серақамаи тоқ" ва ғ. хориҷ кунед.`
}

func (i *iif) makeRu() {
	i.ru = make([]string, i.count)
	i.ru[1] = `Дано целое число. Если оно является положительным, то вычесть из него&nbsp;8; в противном случае не изменять его. Вывести полученное число.`
	i.ru[2] = `Дано целое число. Если оно является положительным, то вычесть из него&nbsp;8; в противном случае прибавить к нему&nbsp;6. Вывести полученное число.`
	i.ru[3] = `Дано целое число. Если оно является положительным, то вычесть из него&nbsp;8; если отрицательным, то прибавить к нему&nbsp;6; если нулевым, то заменить его на&nbsp;10. Вывести полученное число.`
	i.ru[4] = `Даны три целых числа. Найти количество положительных чисел в исходном наборе.`
	i.ru[5] = `Даны три целых числа. Найти количество положительных и количество отрицательных чисел в исходном наборе.`
	i.ru[6] = `Даны два числа. Вывести большее из них.`
	i.ru[7] = `Даны два числа. Вывести порядковый номер меньшего из них.`
	i.ru[8] = `Даны два числа. Вывести вначале большее, а затем меньшее из них.`
	i.ru[9] = `Даны две переменные вещественного типа:&nbsp;<i>A</i>,&nbsp;<i>B</i>. Перераспределить значения данных переменных так, чтобы в&nbsp;<i>A</i> оказалось меньшее из значений, а в&nbsp;<i>B</i> &#8212; большее. Вывести новые значения переменных&nbsp;<i>A</i> и&nbsp;<i>B</i>.`
	i.ru[10] = `Даны две переменные целого типа:&nbsp;<i>A</i> и&nbsp;<i>B</i>. Если их значения не равны, то присвоить каждой переменной сумму этих значений, а если равны, то присвоить переменным нулевые значения. Вывести новые значения переменных&nbsp;<i>A</i> и&nbsp;<i>B</i>.`
	i.ru[11] = `Даны две переменные целого типа:&nbsp;<i>A</i> и&nbsp;<i>B</i>. Если их значения не равны, то присвоить каждой переменной большее из этих значений, а если равны, то присвоить переменным нулевые значения. Вывести новые значения переменных&nbsp;<i>A</i> и&nbsp;<i>B</i>.`
	i.ru[12] = `Даны три числа. Найти наименьшее из них.`
	i.ru[13] = `Даны три числа. Найти среднее из них (т.&nbsp;е. число, расположенное между наименьшим и наибольшим).`
	i.ru[14] = `Даны три числа. Вывести вначале наименьшее, а затем наибольшее из данных чисел.`
	i.ru[15] = `Даны три числа. Найти сумму двух наибольших из них.`
	i.ru[16] = `Даны три переменные вещественного типа:&nbsp;<i>A</i>, <i>B</i>, <i>C</i>. Если их значения упорядочены по возрастанию, то удвоить их; в противном случае заменить значение каждой переменной на противоположное. Вывести новые значения переменных&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>.`
	i.ru[17] = `Даны три переменные вещественного типа:&nbsp;<i>A</i>, <i>B</i>, <i>C</i>. Если их значения упорядочены по возрастанию или убыванию, то удвоить их; в противном случае заменить значение каждой переменной на противоположное. Вывести новые значения переменных&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>.`
	i.ru[18] = `Даны три целых числа, одно из которых отлично от двух других, равных между собой. Определить порядковый номер числа, отличного от остальных.`
	i.ru[19] = `Даны четыре целых числа, одно из которых отлично от трех других, равных между собой. Определить порядковый номер числа, отличного от остальных.`
	i.ru[20] = `На числовой оси расположены три точки:&nbsp;<i>A</i>, <i>B</i>, <i>C</i>. Определить, какая из двух последних точек (<i>B</i> или&nbsp;<i>C</i>) расположена ближе к&nbsp;<i>A</i>, и вывести эту точку и ее расстояние от точки&nbsp;<i>A</i>.`
	i.ru[21] = `Даны целочисленные координаты точки на плоскости. Если точка совпадает с началом координат, то вывести&nbsp;0. Если точка не совпадает с началом координат, но лежит на оси&nbsp;<i>OX</i> или&nbsp;<i>OY</i>, то вывести соответственно&nbsp;1 или&nbsp;2. Если точка не лежит на координатных осях, то вывести&nbsp;3.`
	i.ru[22] = `Даны координаты точки, не лежащей на координатных осях&nbsp;<i>OX</i> и&nbsp;<i>OY</i>. Определить номер координатной четверти, в которой находится данная точка.`
	i.ru[23] = `Даны целочисленные координаты трех вершин прямоугольника, стороны которого параллельны координатным осям. Найти координаты его четвертой вершины.`
	i.ru[24] = `Для данного вещественного&nbsp;<i>x</i> найти значение следующей функции&nbsp;<i>f</i>, принимающей вещественные значения: <table align="center"><tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>2&#183;sin(<i>x</i>),<td align=left>если <i>x</i>&nbsp;&gt;&nbsp;0,         <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>6&nbsp;&#8722;&nbsp;<i>x</i>,<td align=left>если <i>x</i>&nbsp;&#8804;&nbsp;0.</table>`
	i.ru[25] = `Для данного целого&nbsp;<i>x</i> найти значение следующей функции&nbsp;<i>f</i>, принимающей значения целого типа: <table align="center"><tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>2&#183;<i>x</i>,<td align=left>если <i>x</i>&nbsp;&lt;&nbsp;&#8722;2 или <i>x</i>&nbsp;&gt;&nbsp;2, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;3&#183;<i>x</i><td align=left>в противном случае.</table>`
	i.ru[26] = `Для данного вещественного&nbsp;<i>x</i> найти значение следующей функции&nbsp;<i>f</i>, принимающей вещественные значения: <table align="center"><tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;<i>x</i>,<td align=left>если <i>x</i>&nbsp;&#8804;&nbsp;0, <tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right><i>x</i><sup>2</sup>,<td align=left>если 0&nbsp;&lt;&nbsp;<i>x</i>&nbsp;&lt;&nbsp;2, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>4,<td align=left>если <i>x</i>&nbsp;&#8805;&nbsp;2.</table>`
	i.ru[27] = `Для данного вещественного&nbsp;<i>x</i> найти значение следующей функции&nbsp;<i>f</i>, принимающей значения целого типа: <table align="center"><tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>0,<td align=left>если <i>x</i>&nbsp;&lt;&nbsp;0, <tr><td align=right><i>f</i>(<i>x</i>)<td align=center>&nbsp;=&nbsp;<td align=right>1,<td align=left>если <i>x</i> принадлежит [0,&nbsp;1), [2,&nbsp;3),&nbsp;&#8230;, <tr><td align=right>&nbsp;<td align=center>&nbsp;<td align=right>&#8722;1,<td align=left>если <i>x</i> принадлежит [1,&nbsp;2), [3,&nbsp;4),&nbsp;&#8230; .</table>`
	i.ru[28] = `Дан номер года (положительное целое число). Определить количество дней в этом году, учитывая, что обычный год насчитывает 365&nbsp;дней, а високосный &#8212; 366&nbsp;дней. Високосным считается год, делящийся на&nbsp;4, за исключением тех годов, которые делятся на&nbsp;100 и не делятся на&nbsp;400 (например, годы&nbsp;300, 1300 и&nbsp;1900 не являются високосными, а&nbsp;1200 и&nbsp;2000 &#8212; являются).`
	i.ru[29] = `Дано целое число. Вывести его строку-описание вида &#171;отрицательное четное число&#187;, &#171;нулевое число&#187;, &#171;положительное нечетное число&#187; и&nbsp;т.&nbsp;д.`
	i.ru[30] = `Дано целое число, лежащее в диапазоне 1&#8211;999. Вывести его строку-описание вида &#171;четное двузначное число&#187;, &#171;нечетное трехзначное число&#187; и&nbsp;т.&nbsp;д.`
}
