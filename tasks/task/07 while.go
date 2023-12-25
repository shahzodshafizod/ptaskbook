package task

type wwhile struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewWhileTask() Task {
	w := &wwhile{
		count: 31,
		name:  "While",
	}
	w.makeEn()
	w.makeTj()
	w.makeRu()
	return w
}

func (w *wwhile) Count() int   { return w.count }
func (w *wwhile) Name() string { return w.name }
func (w *wwhile) En() []string { return w.en }
func (w *wwhile) Tj() []string { return w.tj }
func (w *wwhile) Ru() []string { return w.ru }

func (w *wwhile) makeEn() {
	w.en = make([]string, w.count)
	w.en[1] = `Two positive real numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>) are given. A segment of length&nbsp;<i>A</i> contains the greatest possible amount of segments of length&nbsp;<i>B</i> (without overlaps). Not using multiplication and division, find the length of unused part of the segment&nbsp;<i>A</i>.`
	w.en[2] = `Two positive real numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>) are given. A segment of length&nbsp;<i>A</i> contains the greatest possible amount of segments of length&nbsp;<i>B</i> (without overlaps). Not using multiplication and division, find the amount of segments&nbsp;<i>B</i>, which are placed on the segment&nbsp;<i>A</i>.`
	w.en[3] = `Two positive integers&nbsp;<i>N</i> and&nbsp;<i>K</i> are given. Using addition and subtraction only, find a quotient of the integer division&nbsp;<i>N</i> on&nbsp;<i>K</i> and also a remainder after this division.`
	w.en[4] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. If it equals&nbsp;3 raised to some integer power then output true, otherwise output false.`
	w.en[5] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0) that equals&nbsp;2 raised to some integer power: <i>N</i>&nbsp;=&nbsp;2<sup><i>K</i></sup>, find the exponent&nbsp;<i>K</i> of the power.`
	w.en[6] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0), compute the <i>double factorial of&nbsp;<i>N</i></i>: <i>N</i>!!&nbsp;=&nbsp;<i>N</i>&#183;(<i>N</i>&#8722;2)&#183;(<i>N</i>&#8722;4)&#183;&#8230;,where the last factor equals&nbsp;2 if&nbsp;<i>N</i> is an even number, and&nbsp;1 otherwise. To avoid the integer overflow, compute the double factorial using a real variable and output the result as a real number.`
	w.en[7] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0), find the smallest positive integer&nbsp;<i>K</i> such that its square is greater than&nbsp;<i>N</i>: <i>K</i><sup>2</sup>&nbsp;&gt;&nbsp;<i>N</i>. Do not use the operation of extracting a root.`
	w.en[8] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0), find the largest integer&nbsp;<i>K</i> such that its square is not greater than&nbsp;<i>N</i>: <i>K</i><sup>2</sup>&nbsp;&#8804;&nbsp;<i>N</i>. Do not use the operation of extracting a root.`
	w.en[9] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;1), find the smallest integer&nbsp;<i>K</i> such that the inequality 3<sup><i>K</i></sup>&nbsp;&gt;&nbsp;<i>N</i> is fulfilled.`
	w.en[10] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;1), find the largest integer&nbsp;<i>K</i> such that the inequality 3<sup><i>K</i></sup>&nbsp;&lt;&nbsp;<i>N</i> is fulfilled.`
	w.en[11] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) is given. Find the smallest integer&nbsp;<i>K</i> such that the sum 1&nbsp;+&nbsp;2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;<i>K</i> is greater than or equal to&nbsp;<i>N</i>. Output&nbsp;<i>K</i> and the corresponding sum.`
	w.en[12] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) is given. Find the largest integer&nbsp;<i>K</i> such that the sum 1&nbsp;+&nbsp;2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;<i>K</i> is less than or equal to&nbsp;<i>N</i>. Output&nbsp;<i>K</i> and the corresponding sum.`
	w.en[13] = `A real number&nbsp;<i>A</i> (&gt;&nbsp;1) is given. Find the smallest integer&nbsp;<i>K</i> such that the sum 1&nbsp;+&nbsp;1/2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;1/<i>K</i> is greater than&nbsp;<i>A</i>. Output&nbsp;<i>K</i> and the corresponding sum.`
	w.en[14] = `A real number&nbsp;<i>A</i> (&gt;&nbsp;1) is given. Find the largest integer&nbsp;<i>K</i> such that the sum 1&nbsp;+&nbsp;1/2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;1/<i>K</i> is less than&nbsp;<i>A</i>. Output&nbsp;<i>K</i> and the corresponding sum.`
	w.en[15] = `A principal of 1000&nbsp;euro is invested at a rate of <i>P</i>&nbsp;percent compounded annually. A real number&nbsp;<i>P</i> is given, 0&nbsp;&lt;&nbsp;<i>P</i>&nbsp;&lt;&nbsp;25. Find, how many years&nbsp;<i>K</i> it will take for an investment to exceed 1100&nbsp;euro. Output&nbsp;<i>K</i> (as an integer) and the compound amount&nbsp;<i>S</i> of the principal at the end of <i>K</i>&nbsp;years (as a real number).`
	w.en[16] = `The skier began trainings having run 10&nbsp;km. Each next day he increased the run distance by <i>P</i>&nbsp;percent from the distance of the last day. A real number&nbsp;<i>P</i> is given, 0&nbsp;&lt;&nbsp;<i>P</i>&nbsp;&lt;&nbsp;50). Find, how many days&nbsp;<i>K</i> it will take for a total run to exceed 200&nbsp;km. Output&nbsp;<i>K</i> (as an integer) and the total run&nbsp;<i>S</i> (as a real number).`
	w.en[17] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0), output all digits of the number&nbsp;<i>N</i> starting from the right digit (a ones digit). Use the operators of integer division and taking the remainder after integer division.`
	w.en[18] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;0), find the amount and the sum of its digits. Use the operators of integer division and taking the remainder after integer division.`
	w.en[19] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. Output an integer obtained from the given one by reading it from right to left. Use the operators of integer division and taking the remainder after integer division.`
	w.en[20] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. Determine whether its decimal representation contains a digit&nbsp;&#34;2&#34; or not, and output true or false respectively. Use the operators of integer division  and taking the remainder after integer division.`
	w.en[21] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;0) is given. Determine whether its decimal representation contains odd digits or not, and output true or false respectively. Use the operators of integer division  and taking the remainder after integer division.`
	w.en[22] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) is given. If it is a <i>prime number</i>, i.&nbsp;e., it has not positive divisors except&nbsp;1 and itself, then output true, otherwise output false.`
	w.en[23] = `Two positive integers&nbsp;<i>A</i> and&nbsp;<i>B</i> are given. Find their <i>greatest common divisor</i> (GCD) using the <i>Euclidean algorithm</i>: GCD(<i>A</i>, <i>B</i>)&nbsp;=&nbsp;GCD(<i>B</i>, <i>A</i>&nbsp;mod&nbsp;<i>B</i>), &nbsp; &nbsp;if <i>B</i>&nbsp;&#8800;&nbsp;0; &nbsp; &nbsp; &nbsp; &nbsp;GCD(<i>A</i>, 0)&nbsp;=&nbsp;<i>A</i>,where &#34;mod&#34; denotes the operator of taking the remainder after integer division.`
	w.en[24] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) is given. An integer-valued sequence of the <i>Fibonacci numbers</i>&nbsp;<i>F</i><sub><i>K</i></sub> is defined as: <i>F</i><sub>1</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>F</i><sub>2</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>F</i><sub><i>K</i></sub>&nbsp;=&nbsp;<i>F</i><sub><i>K</i>&#8722;2</sub>&nbsp;+&nbsp;<i>F</i><sub><i>K</i>&#8722;1</sub>, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;3,&nbsp;4,&nbsp;&#8230;&nbsp;.Determine whether&nbsp;<i>N</i> is a Fibonacci number or not, and output true or false respectively.`
	w.en[25] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;1), find the first Fibonacci number greater than&nbsp;<i>N</i> (see the <i>Fibonacci numbers</i> definition in While24).`
	w.en[26] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;1) that is a Fibonacci number: <i>N</i>&nbsp;=&nbsp;<i>F</i><sub><i>K</i></sub>, output previous and next Fibonacci numbers:&nbsp;<i>F</i><sub><i>K</i>&#8722;1</sub> and&nbsp;<i>F</i><sub><i>K</i>+1</sub> (see the <i>Fibonacci numbers</i> definition in While24).`
	w.en[27] = `Given an integer&nbsp;<i>N</i> (&gt;&nbsp;1) that is a Fibonacci number: <i>N</i>&nbsp;=&nbsp;<i>F</i><sub><i>K</i></sub>, find its order number&nbsp;<i>K</i> (see the <i>Fibonacci numbers</i> definition in While24).`
	w.en[28] = `A real number&nbsp;&#949; (&gt;&nbsp;0) is given. A sequence of real numbers&nbsp;<i>A</i><sub><i>K</i></sub> is defined as: <i>A</i><sub>1</sub>&nbsp;=&nbsp;2, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub><i>K</i></sub>&nbsp;=&nbsp;2&nbsp;+&nbsp;1/<i>A</i><sub><i>K</i>&#8722;1</sub>, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;2,&nbsp;3,&nbsp;&#8230;&nbsp;.Find the first index&nbsp;<i>K</i> such that the inequality |<i>A</i><sub><i>K</i></sub>&nbsp;&#8722;&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub>|&nbsp;&lt;&nbsp;&#949;  is fulfilled. Output the index&nbsp;<i>K</i> and the terms&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub> and&nbsp;<i>A</i><sub><i>K</i></sub>.`
	w.en[29] = `A real number&nbsp;&#949; (&gt;&nbsp;0) is given. A sequence of real numbers&nbsp;<i>A</i><sub><i>K</i></sub> is defined as: <i>A</i><sub>1</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub>2</sub>&nbsp;=&nbsp;2, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub><i>K</i></sub>&nbsp;=&nbsp;(<i>A</i><sub><i>K</i>&#8722;2</sub>&nbsp;+&nbsp;2&#183;<i>A</i><sub><i>K</i>&#8722;1</sub>)/3, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;3,&nbsp;4,&nbsp;&#8230;&nbsp;.Find the first index&nbsp;<i>K</i> such that the inequality |<i>A</i><sub><i>K</i></sub>&nbsp;&#8722;&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub>|&nbsp;&lt;&nbsp;&#949;  is fulfilled. Output the index&nbsp;<i>K</i> and the terms&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub> and&nbsp;<i>A</i><sub><i>K</i></sub>.`
	w.en[30] = `Three positive real numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are given. A rectangle of size&nbsp;<i>A</i>&nbsp;&#215;&nbsp;<i>B</i> contains the greatest possible amount of squares with side&nbsp;<i>C</i> (without overlaps). Find the amount of squares placed on the rectangle. Do not use the operators of multiplication and division.`
}

func (w *wwhile) makeTj() {
	w.tj = make([]string, w.count)
	w.tj[1] = `Ададҳои мусбии A ва B (A>B) дода шудаанд. Дар порчаи дарозиаш A миқдори калонтарини имконпазири порчаҳои дарозиашон B ҷой дода шудаанд. Амалҳои зарб ва тақсимро истифода набурда, дарозии қисми банднабудаи порчаи A-ро ёбед.`
	w.tj[2] = `Ададҳои мусбии A ва B (A>B) дода шудаанд. Дар порчаи дарозиаш A миқдори калонтарини имконпазири порчаҳои дарозиашон B ҷой дода шудаанд. Амалҳои зарб ва тақсимро истифода набурда, миқдори порчаҳои B-ро, ки дар порчаи A ҷойгиранд, ёбед.`
	w.tj[3] = `Ададҳои мусбии бутуни N ва K дода шудаанд. Фақат амалҳои ҷамъ ва тарҳро истифода бурда, қимати тақсими бутуни N-ро ба K, ҳамчунин боқимондаи ин тақсимро ёбед.`
	w.tj[4] = `Адади бутуни N(>0) дода шудааст. Агар он дараҷаи адади 3 бошад, пас True-ро хориҷ кунед, агар набошад — False-ро хориҷ кунед.`
	w.tj[5] = `Адади бутуни N(>0), ки якчанд дараҷаи адади 2 аст, дода шудааст: N = 2<sup>K</sup>. Адади бутуни K — нишондиҳандаи ин дараҷаро ёбед.`
	w.tj[6] = `Адади бутуни N(>0) дода шудааст. Факториали дукаратаи N-ро ёбед: N!!=N·(N–2)·(N–4)·… (зарбшавандаи охирин ба 2 баробар аст, агар N — ҷуфт бошад ва ба 1 баробар аст, агар N — тоқ бошад). Барои роҳ надодан ба пуршавии қатори ададҳои бутун ин зарбкуниро бо ёрии тағйирёбандаи ҳақиқӣ иҷро кунед ва онро чун адади ҳақиқӣ хориҷ кунед.`
	w.tj[7] = `Адади бутуни N(>0) дода шудааст. Адади бутуни мусбии хурдтарини K-ро ёбед, ки квадрати он аз N калон аст: K<sup>2</sup>>N. Функсияи азрешабарории квадратиро истифода набаред.`
	w.tj[8] = `Адади бутуни N(>0) дода шудааст. Адади бутуни калонтарини K-ро ёбед, ки квадрати он аз N калон нест: K<sup>2</sup>≤N. Функсияи азрешабарории квадратиро истифода набаред.`
	w.tj[9] = `Адади бутуни N(>1) дода шудааст. Адади бутуни хурдтарини K-ро ёбед, ки нобаробарии зеринро иҷро мекунад 3<sup>K</sup>>N.`
	w.tj[10] = `Адади бутуни N(>1) дода шудааст. Адади бутуни калонтарини K-ро ёбед, ки нобаробарии зеринро иҷро мекунад 3<sup>K</sup>&lt;N.`
	w.tj[11] = `Адади бутуни N (>1) дода шудааст. Аз ададҳои бутуни K хурдтаринашро хориҷ кунед, ки барои он суммаи 1 + 2 + … + K калон ё баробари N мешавад ва ҳамчунин худи ин суммаро низ хориҷ кунед.`
	w.tj[12] = `Адади бутуни N (>1) дода шудааст. Аз ададҳои бутуни K калонтаринашро хориҷ кунед, ки барои он суммаи 1+2+…+K хурд ё баробари N мешавад ва ҳамчунин ин суммаро низ хориҷ кунед.`
	w.tj[13] = `Адади A(>1) дода шудааст. Аз ададҳои бутуни K хурдтаринашро хориҷ кунед, ки барои он суммаи 1+1/2+…+1/K калони A мешавад ва ҳамчунин худи ин суммаро низ хориҷ кунед.`
	w.tj[14] = `Адади A(>1) дода шудааст. Аз ададҳои бутуни K калонтаринашро хориҷ кунед, ки барои он суммаи 1+1/2+..+1/K хурди A мешавад ва ҳамчунин худи ин суммаро низ хориҷ кунед.`
	w.tj[15] = `Пасандози аввалин дар бонк ба 1000 сомонӣ баробар аст. Пас аз ҳар моҳ андозаи пасандоз ба P фоиз аз суммаи мавҷуда зиёд мешавад (P — адади ҳақиқӣ, 0&lt;P&lt;25). Аз рӯи P-и додашуда муайян кунед, ки пас аз чанд моҳ андозаи пасандоз аз 1100 сомонӣ зиёд мешавад. Миқдори моҳҳои ёфташуда K (адади бутун) ва андозаи ниҳоии пасандоз S (адади ҳақиқӣ)-ро ёбед.`
	w.tj[16] = `Варзишгар-лижарон машқро бо давидани 10 км дар рӯзи аввал сар кард. Ҳар як рӯзи оянда он дарозии давишро ба P фоизи рӯзи пешина зиёд кард (P - адади ҳақиқӣ, 0&lt;P&lt;50). Аз рӯи P-и додашуда муайян кунед, ки пас аз чанд рӯз давиши натиҷавии лижарон барои ҳамаи рӯзҳо аз 200 км зиёд мешавад. Миқдори ёфташудаи рӯзҳо K (адади бутун) ва давиши натиҷавӣ S (адади ҳақиқӣ)-ро ёбед.`
	w.tj[17] = `Адади бутуни N(>0) дода шудааст. Бо истифодабарии амалҳои тақсими бутун ва гирифтани бақия аз ҳосили тақсим ҳамаи рақамҳои онро сар карда аз тарафи рост (қатори воҳидӣ) хориҷ кунед.`
	w.tj[18] = `Адади бутуни N(>0) дода шудааст. Бо истифодабарии амалҳои тақсими бутун ва гирифтани бақия аз ҳосили тақсим миқдор ва суммаи рақамҳои онро ёбед.`
	w.tj[19] = `Адади бутуни N(>0) дода шудааст. Бо истифодабарии амалҳои тақсими бутун ва гирифтани бақия аз ҳосили тақсим ададеро ёбед, ки ҳангоми хондани адади N аз рост ба чап ҳосил гардидааст.`
	w.tj[20] = `Адади бутуни N(>0) дода шудааст. Бо ёрии амалҳои тақсими бутун ва гирифтани бақия аз ҳосили тақсим муайян кунед, ки дар навишти адади N рақами «2» ҳаст ё не. Агар бошад, пас True-ро хориҷ кунед, вагарна — False-ро хориҷ кунед.`
	w.tj[21] = `Адади бутуни N(>0) дода шудааст. Бо ёрии амалҳои тақсими бутун ва гирифтани бақия аз ҳосили тақсим муайян кунед, ки дар навишти адади N рақами тоқ ҳаст ё не. Агар бошад, пас True-ро хориҷ кунед, вагарна - False-ро хориҷ кунед.`
	w.tj[22] = `Адади бутуни N(>1) дода шудааст. Агар он адади оддӣ бошад, яъне ба ғайр аз 1 ва худаш тақсимкунандаи бутун надошта бошад, пас True-ро хориҷ кунед, вагарна False-ро хориҷ кунед.`
	w.tj[23] = `Ададҳои бутуни мусбии A ва B дода шудаанд. Бо истифодабарии алгоритми Евклид калонтарин тақсимкунандаи умумӣ (КТУ)-и онҳоро ёбед: КТУ(A,B)=КТУ(B,A mod B), агар B<>0; КТУ(A,0)=A, агар B=0 бошад.`
	w.tj[24] = `Адади бутуни N(>1) дода шудааст. Пайдарпайии ададҳои Фибоначчи F<sub>k</sub> ба тариқи зайл муайян карда мешавад: F<sub>1</sub>=1, F<sub>2</sub>=1, F<sub>k</sub>= F<sub>k-2</sub>+F<sub>k-1</sub>, K=3,4, ... . Санҷед, ки адади N адади Фибоначчи аст ё не. Агар бошад, пас True-ро хориҷ кунед, вагарна - False-ро хориҷ кунед.`
	w.tj[25] = `Адади бутуни N(>1) дода шудааст. Аввалин адади Фибоначчиро ёбед, ки аз N калон аст. (муайянкунии ададҳои Фибоначчи дар масъалаи While24 оварда шудааст).`
	w.tj[26] = `Адади бутуни N(>1), ки адади Фибоначчи мебошад: N=F<sub>k</sub> (муайянкунии ададҳои Фибоначчи дар масъалаи While24 оварда шудааст) дода шудааст. Ададҳои бутуни F<sub>k-1</sub> ва F<sub>k+1</sub> - ададҳои пешина ва пасинаи Фибоначчиро ёбед.`
	w.tj[27] = `Адади бутуни N(>1), ки адади Фибоначчи мебошад: N=F<sub>k</sub> (муайянкунии ададҳои Фибоначчи дар масъалаи While24 оварда шудааст) дода шудааст. Адади бутуни K — рақами тартибии адади Фибоначчи N-ро ёбед.`
	w.tj[28] = `Адади ҳақиқии eps(>0) дода шудааст. Пайдарпайии ададҳои ҳақиқии A<sub>k</sub> ба тариқи зайл муайян карда мешавад: A<sub>1</sub>=2, A<sub>k</sub>=2+1/A<sub>k-1</sub>, K=2, 3, ... . Якӯмин рақами K-ро ёбед, ки барои он шарти |A<sub>k</sub>–A<sub>k-1</sub>|&lt;eps иҷро мегардад ва ин рақамро, ҳамчунин ададҳои A<sub>k-1</sub> ва A<sub>k</sub>-ро хориҷ кунед.`
	w.tj[29] = `Адади ҳақиқии eps(>0) дода шудааст. Пайдарпайии ададҳои ҳақиқии A<sub>k</sub> ба тариқи зайл муайян карда мешавад:A<sub>1</sub>=1, A<sub>2</sub>=2, A<sub>k</sub>=(A<sub>k-2</sub>+2*A<sub>k-1</sub>)/3, K=3, 4, ... . Якӯмин рақами K-ро ёбед, ки барои он шарти |A<sub>k</sub>–A<sub>k-1</sub>|&lt;eps мегардад ва ин рақамро, ҳамчунин ададҳои A<sub>k-1</sub> ва A<sub>k</sub>-ро хориҷ кунед.`
	w.tj[30] = `Ададҳои мусбии A, B, C дода шудаанд. Дар росткунҷаи андозааш AxB миқдори калонтарини имконпазири квадратҳо бо тарафи андозааш C ҷойгир аст. Миқдори квадратҳои дар росткунҷа ҷойгирифтаро ёбед. Амалҳои зарб ва тақсимро истифода набаред.`
}

func (w *wwhile) makeRu() {
	w.ru = make([]string, w.count)
	w.ru[1] = `Даны положительные числа&nbsp;<i>A</i> и&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). На отрезке длины&nbsp;<i>A</i> размещено максимально возможное количество отрезков длины&nbsp;<i>B</i> (без наложений). Не используя операции умножения и деления, найти длину незанятой части отрезка&nbsp;<i>A</i>.`
	w.ru[2] = `Даны положительные числа&nbsp;<i>A</i> и&nbsp;<i>B</i> (<i>A</i>&nbsp;&gt;&nbsp;<i>B</i>). На отрезке длины&nbsp;<i>A</i> размещено максимально возможное количество отрезков длины&nbsp;<i>B</i> (без наложений). Не используя операции умножения и деления, найти количество отрезков&nbsp;<i>B</i>, размещенных на отрезке&nbsp;<i>A</i>.`
	w.ru[3] = `Даны целые положительные числа&nbsp;<i>N</i> и&nbsp;<i>K</i>. Используя только операции сложения и вычитания, найти частное от деления нацело&nbsp;<i>N</i> на&nbsp;<i>K</i>, а также остаток от этого деления.`
	w.ru[4] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Если оно является степенью числа&nbsp;3, то вывести true, если не является &#8212; вывести false.`
	w.ru[5] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0), являющееся некоторой степенью числа&nbsp;2: <i>N</i>&nbsp;=&nbsp;2<sup><i>K</i></sup>. Найти целое число&nbsp;<i>K</i> &#8212; показатель этой степени.`
	w.ru[6] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Найти <i>двойной факториал&nbsp;<i>N</i></i>:<i>N</i>!!&nbsp;=&nbsp;<i>N</i>&#183;(<i>N</i>&#8722;2)&#183;(<i>N</i>&#8722;4)&#183;&#8230;(последний сомножитель равен&nbsp;2, если <i>N</i> &#8212; четное, и&nbsp;1, если <i>N</i> &#8212; нечетное). Чтобы избежать целочисленного переполнения, вычислять это произведение с помощью вещественной переменной и вывести его как вещественное число.`
	w.ru[7] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Найти наименьшее целое положительное число&nbsp;<i>K</i>, квадрат которого превосходит&nbsp;<i>N</i>: <i>K</i><sup>2</sup>&nbsp;&gt;&nbsp;<i>N</i>. Функцию извлечения квадратного корня не использовать.`
	w.ru[8] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Найти наибольшее целое число&nbsp;<i>K</i>, квадрат которого не превосходит <i>N</i>: <i>K</i><sup>2</sup>&nbsp;&#8804;&nbsp;<i>N</i>. Функцию извлечения квадратного корня не использовать.`
	w.ru[9] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Найти наименьшее целое число&nbsp;<i>K</i>, при котором выполняется неравенство 3<sup><i>K</i></sup>&nbsp;&gt;&nbsp;<i>N</i>.`
	w.ru[10] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Найти наибольшее целое число&nbsp;<i>K</i>, при котором выполняется неравенство 3<sup><i>K</i></sup>&nbsp;&lt;&nbsp;<i>N</i>.`
	w.ru[11] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Вывести наименьшее из целых чисел&nbsp;<i>K</i>, для которых сумма 1&nbsp;+&nbsp;2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;<i>K</i> будет больше или равна&nbsp;<i>N</i>, и саму эту сумму.`
	w.ru[12] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Вывести наибольшее из целых чисел&nbsp;<i>K</i>, для которых сумма 1&nbsp;+&nbsp;2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;<i>K</i> будет меньше или равна&nbsp;<i>N</i>, и саму эту сумму.`
	w.ru[13] = `Дано число&nbsp;<i>A</i> (&gt;&nbsp;1). Вывести наименьшее из целых чисел&nbsp;<i>K</i>, для которых сумма 1&nbsp;+&nbsp;1/2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;1/<i>K</i> будет больше&nbsp;<i>A</i>, и саму эту сумму.`
	w.ru[14] = `Дано число&nbsp;<i>A</i> (&gt;&nbsp;1). Вывести наибольшее из целых чисел&nbsp;<i>K</i>, для которых сумма 1&nbsp;+&nbsp;1/2&nbsp;+&nbsp;&#8230;&nbsp;+&nbsp;1/<i>K</i> будет меньше&nbsp;<i>A</i>, и саму эту сумму.`
	w.ru[15] = `Начальный вклад в банке равен 1000&nbsp;руб. В конце каждого месяца размер вклада увеличивается на <i>P</i>&nbsp;процентов от имеющейся суммы (<i>P</i>&nbsp;&#8212; вещественное число, 0&nbsp;&lt;&nbsp;<i>P</i>&nbsp;&lt;&nbsp;25). По данному&nbsp;<i>P</i> определить, через сколько месяцев размер вклада превысит 1100&nbsp;руб., и вывести найденное количество месяцев&nbsp;<i>K</i> (целое число) и итоговый размер вклада&nbsp;<i>S</i> (вещественное число).`
	w.ru[16] = `Спортсмен-лыжник начал тренировки, пробежав в первый день 10&nbsp;км. Каждый следующий день он увеличивал длину пробега на <i>P</i>&nbsp;процентов от пробега предыдущего дня (<i>P</i>&nbsp;&#8212; вещественное, 0&nbsp;&lt;&nbsp;<i>P</i>&nbsp;&lt;&nbsp;50). По данному&nbsp;<i>P</i> определить, после какого дня суммарный пробег лыжника за все дни превысит 200&nbsp;км, и вывести найденное количество дней&nbsp;<i>K</i> (целое) и суммарный пробег&nbsp;<i>S</i> (вещественное число).`
	w.ru[17] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Используя операции деления нацело и взятия остатка от деления, вывести все его цифры, начиная с самой правой (разряда единиц).`
	w.ru[18] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Используя операции деления нацело и взятия остатка от деления, найти количество и сумму его цифр.`
	w.ru[19] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). Используя операции деления нацело и взятия остатка от деления, найти число, полученное при прочтении числа&nbsp;<i>N</i> справа налево.`
	w.ru[20] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). С помощью операций деления нацело и взятия остатка от деления определить, имеется ли в записи числа&nbsp;<i>N</i> цифра&nbsp;&#171;2&#187;. Если имеется, то вывести true, если нет &#8212; вывести false.`
	w.ru[21] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;0). С помощью операций деления нацело и взятия остатка от деления определить, имеются ли в записи числа&nbsp;<i>N</i> нечетные цифры. Если имеются, то вывести true, если нет &#8212; вывести false.`
	w.ru[22] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Если оно является <i>простым</i>, т.&nbsp;е. не имеет положительных делителей, кроме&nbsp;1 и самого себя, то вывести true, иначе вывести false.`
	w.ru[23] = `Даны целые положительные числа&nbsp;<i>A</i> и&nbsp;<i>B</i>. Найти их <i>наибольший общий делитель</i> (НОД), используя <i>алгоритм Евклида</i>: НОД(<i>A</i>, <i>B</i>)&nbsp;=&nbsp;НОД(<i>B</i>, <i>A</i>&nbsp;mod&nbsp;<i>B</i>), &nbsp; &nbsp;если <i>B</i>&nbsp;&#8800;&nbsp;0; &nbsp; &nbsp; &nbsp; &nbsp;НОД(<i>A</i>, 0)&nbsp;=&nbsp;<i>A</i>,где &#171;mod&#187; обозначает операцию взятия остатка от деления.`
	w.ru[24] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Последовательность <i>чисел Фибоначчи</i>&nbsp;<i>F</i><sub><i>K</i></sub> определяется следующим образом: <i>F</i><sub>1</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>F</i><sub>2</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>F</i><sub><i>K</i></sub>&nbsp;=&nbsp;<i>F</i><sub><i>K</i>&#8722;2</sub>&nbsp;+&nbsp;<i>F</i><sub><i>K</i>&#8722;1</sub>, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;3,&nbsp;4,&nbsp;&#8230;&nbsp;.Проверить, является ли число&nbsp;<i>N</i> числом Фибоначчи. Если является, то вывести true, если нет &#8212; вывести false.`
	w.ru[25] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1). Найти первое число Фибоначчи, большее&nbsp;<i>N</i> (определение <i>чисел Фибоначчи</i> дано в задании While24).`
	w.ru[26] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1), являющееся числом Фибоначчи: <i>N</i>&nbsp;=&nbsp;<i>F</i><sub><i>K</i></sub> (определение <i>чисел Фибоначчи</i> дано в задании While24). Найти целые числа&nbsp;<i>F</i><sub><i>K</i>&#8722;1</sub> и&nbsp;<i>F</i><sub><i>K</i>+1</sub> &#8212; предыдущее и последующее числа Фибоначчи.`
	w.ru[27] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1), являющееся числом Фибоначчи: <i>N</i>&nbsp;=&nbsp;<i>F</i><sub><i>K</i></sub> (определение <i>чисел Фибоначчи</i> дано в задании While24). Найти целое число&nbsp;<i>K</i> &#8212; порядковый номер числа Фибоначчи&nbsp;<i>N</i>.`
	w.ru[28] = `Дано вещественное число&nbsp;&#949; (&gt;&nbsp;0). Последовательность вещественных чисел <i>A</i><sub><i>K</i></sub> определяется следующим образом: <i>A</i><sub>1</sub>&nbsp;=&nbsp;2, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub><i>K</i></sub>&nbsp;=&nbsp;2&nbsp;+&nbsp;1/<i>A</i><sub><i>K</i>&#8722;1</sub>, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;2,&nbsp;3,&nbsp;&#8230;&nbsp;.Найти первый из номеров&nbsp;<i>K</i>, для которых выполняется условие |<i>A</i><sub><i>K</i></sub>&nbsp;&#8722;&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub>|&nbsp;&lt;&nbsp;&#949;, и вывести этот номер, а также числа&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub> и&nbsp;<i>A</i><sub><i>K</i></sub>.`
	w.ru[29] = `Дано вещественное число&nbsp;&#949; (&gt;&nbsp;0). Последовательность вещественных чисел <i>A</i><sub><i>K</i></sub> определяется следующим образом: <i>A</i><sub>1</sub>&nbsp;=&nbsp;1, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub>2</sub>&nbsp;=&nbsp;2, &nbsp; &nbsp; &nbsp; &nbsp;<i>A</i><sub><i>K</i></sub>&nbsp;=&nbsp;(<i>A</i><sub><i>K</i>&#8722;2</sub>&nbsp;+&nbsp;2&#183;<i>A</i><sub><i>K</i>&#8722;1</sub>)/3, &nbsp; &nbsp;<i>K</i>&nbsp;=&nbsp;3,&nbsp;4,&nbsp;&#8230;&nbsp;.Найти первый из номеров&nbsp;<i>K</i>, для которых выполняется условие |<i>A</i><sub><i>K</i></sub>&nbsp;&#8722;&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub>|&nbsp;&lt;&nbsp;&#949;, и вывести этот номер, а также числа&nbsp;<i>A</i><sub><i>K</i>&#8722;1</sub> и&nbsp;<i>A</i><sub><i>K</i></sub>.`
	w.ru[30] = `Даны положительные числа&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>. На прямоугольнике размера&nbsp;<i>A</i>&nbsp;&#215;&nbsp;<i>B</i> размещено максимально возможное количество квадратов со стороной&nbsp;<i>C</i> (без наложений). Найти количество квадратов, размещенных на прямоугольнике. Операции умножения и деления не использовать.`
}
