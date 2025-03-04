package task

type sseries struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewSeriesTask() Task {
	s := &sseries{
		count: 41,
		name:  "Series",
	}
	s.makeEn()
	s.makeTj()
	s.makeRu()
	return s
}

func (s *sseries) Count() int   { return s.count }
func (s *sseries) Name() string { return s.name }
func (s *sseries) En() []string { return s.en }
func (s *sseries) Tj() []string { return s.tj }
func (s *sseries) Ru() []string { return s.ru }

func (s *sseries) makeEn() {
	s.en = make([]string, s.count)
	s.en[1] = `Given ten real numbers, find their sum.`
	s.en[2] = `Given ten real numbers, find their product.`
	s.en[3] = `Given ten real numbers, find their average.`
	s.en[4] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. Output the sum and the product of all elements of this sequence.`
	s.en[5] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;positive real numbers are given. Output in the same order the integer parts of all elements of this sequence (as real numbers with zero fractional part). Also output the sum of all integer parts.`
	s.en[6] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;positive real numbers are given. Output in the same order the fractional parts of all elements of this sequence (as real numbers with zero integer part). Also output the product of all fractional parts.`
	s.en[7] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. Output in the same order the rounded values of all elements of this sequence to the nearest whole number (as integers). Also output the sum of all rounded values.`
	s.en[8] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output in the same order all even-valued elements of the sequence and also their amount&nbsp;<i>K</i>.`
	s.en[9] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output in the same order the order numbers of all odd-valued elements of the sequence and also their amount&nbsp;<i>K</i>.`
	s.en[10] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output the logical value true if the sequence contains positive-valued elements, otherwise output false.`
	s.en[11] = `Integers&nbsp;<i>K</i>,&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. Output the logical value false if the sequence contains elements of value less than&nbsp;<i>K</i>, otherwise output false.`
	s.en[12] = `A sequence of nonzero integers terminated by zero is given (the final zero is not an element of the sequence). Output the length of the sequence.`
	s.en[13] = `A sequence of nonzero integers terminated by zero is given. Output the sum of all positive-valued elements of the sequence. If the sequence does not contain the required elements then output&nbsp;0.`
	s.en[14] = `An integer&nbsp;<i>K</i> and a sequence of nonzero integers terminated by zero are given (the final zero is not an element of the sequence). Output the amount of elements whose value less than&nbsp;<i>K</i>.`
	s.en[15] = `An integer&nbsp;<i>K</i> and a sequence of nonzero integers terminated by zero are given (the final zero is not an element of the sequence). Output the order number of the first element whose value greater than&nbsp;<i>K</i>. If the sequence does not contain the required elements then output&nbsp;0.`
	s.en[16] = `An integer&nbsp;<i>K</i> and a sequence of nonzero integers terminated by zero are given (the final zero is not an element of the sequence). Output the order number of the last element whose value greater than&nbsp;<i>K</i>. If the sequence does not contain the required elements then output&nbsp;0.`
	s.en[17] = `A real number&nbsp;<i>B</i>, an integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers are given. The values of elements of the sequence are in ascending order. Output the number&nbsp;<i>B</i> jointly with the elements of the sequence so that all output numbers were in ascending order.`
	s.en[18] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. The values of elements of the sequence are in ascending order, the sequence may contain equal elements. Output in the same order all distinct elements of the sequence.`
	s.en[19] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of <i>N</i>&nbsp;integers are given.  Output the elements that are less than their left-side neighbor. Also output the amount&nbsp;<i>K</i> of such elements.`
	s.en[20] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of <i>N</i>&nbsp;integers are given.  Output the elements that are less than their right-side neighbor. Also output the amount&nbsp;<i>K</i> of such elements.`
	s.en[21] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of N real numbers are given. Output the logical value true if the values of elements are in ascending order, otherwise output false.`
	s.en[22] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;1) and a sequence of <i>N</i>&nbsp;real numbers are given. Output&nbsp;0 if the values of elements are in descending order, otherwise output the order number of the first element that breaks the required order.`
	s.en[23] = `An integer&nbsp;<i>N</i> (&gt;&nbsp;2) and a sequence of <i>N</i>&nbsp;real numbers are given. A sequence is called a <i>sawtooth</i> one if each inner element of the sequence is either greater or less than both of its neighbors (that is, each inner element is a <i>tooth</i>). Output&nbsp;0 if the given sequence is a sawtooth one, otherwise output the order number of the first element that is not a tooth.`
	s.en[24] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. The sequence contains at least two zero-valued elements. Output the sum of the values of elements placed between two last zero-valued elements. If two last zero-valued elements are placed side by side then output&nbsp;0.`
	s.en[25] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;integers are given. The sequence contains at least two zero-valued elements. Output the sum of the values of elements placed between the first and the last zero-valued elements. If the first and the last zero-valued elements are placed side by side then output&nbsp;0.`
	s.en[26] = `Positive integers&nbsp;<i>K</i>, <i>N</i> and a sequence of <i>N</i>&nbsp;real numbers&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub> are given. For each element of the sequence find its value raised to the power of&nbsp;<i>K</i>: (<i>A</i><sub>1</sub>)<sup><i>K</i></sup>, (<i>A</i><sub>2</sub>)<sup><i>K</i></sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i></sub>)<sup><i>K</i></sup>`
	s.en[27] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub> are given. Output the following numbers:<i>A</i><sub>1</sub>, (<i>A</i><sub>2</sub>)<sup>2</sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i>&#8722;1</sub>)<sup><i>N</i>&#8722;1</sup>, (<i>A</i><sub><i>N</i></sub>)<sup><i>N</i></sup>`
	s.en[28] = `An integer&nbsp;<i>N</i> and a sequence of <i>N</i>&nbsp;real numbers&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub> are given. Output the following numbers:(<i>A</i><sub>1</sub>)<sup><i>N</i></sup>, (<i>A</i><sub>2</sub>)<sup><i>N</i>&#8722;1</sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i>&#8722;1</sub>)<sup>2</sup>, <i>A</i><sub><i>N</i></sub>`
	s.en[29] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Find the total sum of the values of elements of all given sequences.`
	s.en[30] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Find the sum of the values of all elements for each given sequence.`
	s.en[31] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Find the amount of the sequences containing an element with the value&nbsp;2.`
	s.en[32] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Output the order number of the first element with the value&nbsp;2 for each given sequence (if a sequence does not contain elements with the required value then output&nbsp;0).`
	s.en[33] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Output the order number of the last element with the value&nbsp;2 for each given sequence (if a sequence does not contain elements with the required value then output&nbsp;0).`
	s.en[34] = `Integers&nbsp;<i>K</i>, <i>N</i> and <i>K</i>&nbsp;sequences of integers are given. Each given sequence contains <i>N</i>&nbsp;elements. Process each sequence as follows: output the sum of the values of all its elements if the sequence contains an element with the value&nbsp;2, otherwise output&nbsp;0.`
	s.en[35] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence is terminated by zero, which is not an element of the sequence. Output the length of each given sequence. Also output the total length of all given sequences.`
	s.en[36] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence contains at least two elements and is terminated by zero, which is not an element of the sequence. Output the amount of the sequences whose element values are in ascending order.`
	s.en[37] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence contains at least two elements and is terminated by zero, which is not an element of the sequence. Output the amount of sequences whose element values are in ascending or in descending order.`
	s.en[38] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence contains at least two elements and is terminated by zero, which is not an element of the sequence. Process each sequence as follows: output&nbsp;1 or&nbsp;&#8722;1 if its element values are in ascending or in descending order respectively, otherwise output&nbsp;0.`
	s.en[39] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence contains at least three elements and is terminated by zero, which is not an element of the sequence. Output the amount of the sawtooth sequences (see the definition of a <i>sawtooth sequence</i> in Series23).`
	s.en[40] = `An integer&nbsp;<i>K</i> and <i>K</i>&nbsp;sequences of nonzero integers are given. Each given sequence contains at least three elements and is terminated by zero, which is not an element of the sequence. Process each sequence as follows: output its length if the sequence is a sawtooth one (see Series23), otherwise output the order number of its first element that is not a tooth.`
}

func (s *sseries) makeTj() {
	s.tj = make([]string, s.count)
	s.tj[1] = `Даҳ то ададҳои ҳақиқӣ дода шудаанд. Суммаи онҳоро ёбед.`
	s.tj[2] = `Даҳ то ададҳои ҳақиқӣ дода шудаанд. Ҳосилизарби онҳоро ёбед.`
	s.tj[3] = `Даҳ то ададҳои ҳақиқӣ дода шудаанд. Қимати миёнаи арифметикии онҳоро ёбед.`
	s.tj[4] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Сумма ва ҳосилизарби ададҳоро аз маҷмӯи мазкур хориҷ кунед.`
	s.tj[5] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқии мусбӣ дода шудаанд. Қисмҳои бутуни ҳамаи ададҳоро (чун ададҳои ҳақиқӣ бо қисми даҳии нулӣ) аз маҷмӯи мазкур бо ҳамин тартиб хориҷ кунед, ҳамчунин суммаи ҳамаи қисмҳои бутунро низ хориҷ кунед.`
	s.tj[6] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқии мусбӣ дода шудаанд. Қисмҳои даҳии ҳамаи ададҳоро (чун ададҳои ҳақиқӣ бо қисми бутуни нулӣ), ҳамчунин  ҳосилизарби ҳамаи қисмҳои даҳиро хориҷ кунед.`
	s.tj[7] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Қиматҳои яклухткардашудаи ҳамаи ададҳоро аз маҷмӯи мазкур (чун ададҳои бутун), ҳамчунин суммаи қиматҳои яклухткардашударо низ хориҷ кунед.`
	s.tj[8] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Ҳамаи ададҳои ҷуфтро аз маҷмӯи мазкур бо ҳамон тартиб ва миқдори чунин ададҳо K-ро хориҷ кунед.`
	s.tj[9] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Рақамҳои ҳамаи ададҳои тоқро аз маҷмӯи мазкур бо ҳамон тартиб ва миқдори чунин ададҳо K-ро хориҷ кунед.`
	s.tj[10] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Агар дар маҷмӯъ адади мусбӣ бошад, пас True-ро хориҷ кунед; дар ҳолати акс бошад, False-ро хориҷ кунед.`
	s.tj[11] = `Ададҳои бутуни K, N ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Агар дар маҷмӯъ адади аз K хурд бошад, пас True-ро хориҷ кунед; дар ҳолати акс False-ро хориҷ кунед.`
	s.tj[12] = `Маҷмӯи ададҳои бутуни ғайринулӣ дода шудааст; нишонаи баанҷомрасии он - адади 0 аст. Миқдори ададҳоро дар маҷмӯъ хориҷ кунед.`
	s.tj[13] = `Маҷмӯи ададҳои бутуни ғайринулӣ дода шудааст; нишонаи баанҷомрасии он - адади 0 аст. Суммаи ҳамаи ададҳои ҷуфти мусбиро аз маҷмӯи мазкур хориҷ кунед. Агар адади талабкардашуда дар маҷмӯъ набошад, пас 0(нул)-ро хориҷ кунед.`
	s.tj[14] = `Адади бутуни K ва маҷмӯъ аз ададҳои бутуни ғайринулӣ дода шудаанд; нишонаи баанҷомрасии он - адади 0 аст. Миқдори ададҳои аз K хурдро дар маҷмӯъ хориҷ кунед.`
	s.tj[15] = `Адади бутуни K ва маҷмӯи ададҳои бутуни ғайринулӣ дода шудаанд; нишонаи баанҷомрасии он - адади 0 аст. Рақами адади аввалини аз K калонро дар маҷмӯъ хориҷ кунед. Агар чунин адад набошад, пас 0(нул)-ро хориҷ кунед.`
	s.tj[16] = `Адади бутуни K ва маҷмӯи ададҳои бутуни ғайринулӣ дода шудаанд; нишонаи баанҷомрасии он - адади 0 аст. Рақами адади охирини аз K калонро дар маҷмӯъ хориҷ кунед. Агар чунин адад набошад, пас 0(нул)-ро хориҷ кунед.`
	s.tj[17] = `Адади ҳақиқии B, адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқии афзуншавандаи батартибгузошташуда дода шудаанд. Элементҳои маҷмӯъро якҷоя бо адади B хориҷ кунед, дар ин ҳангом тартиби ададҳои хориҷшавандаро нигоҳ доред.`
	s.tj[18] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутуни афзуншавандаи батартибгузошташуда дода шудаанд. Маҷмӯи мазкур элементҳои якхеларо низ доро буда метавонад. Ҳамаи элементҳои гуногуни маҷмӯи мазкурро бо ҳамон тартиб хориҷ кунед.`
	s.tj[19] = `Адади бутуни N ( > 1) ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Он элементҳои маҷмӯъро, ки онҳо аз ҳамсояи чапи худ хурд ҳастанд ва ҳамчунин миқдори чунин элементҳо K-ро хориҷ кунед.`
	s.tj[20] = `Адади бутуни N ( > 1) ва маҷмӯъ аз N ададҳои бутун дода шудаанд. Он элементҳои маҷмӯъро, ки онҳо аз ҳамсояи рости худ хурд ҳастанд ва миқдори чунин элементҳо K-ро хориҷ кунед.`
	s.tj[21] = `Адади бутуни N (> 1) ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Санҷед, ки маҷмӯи мазкур пайдарпайии афзуншавандаро ташкил медиҳад ё не. Агар ташкил диҳад, пас True-ро хориҷ кунед, вагарна - False-ро хориҷ кунед.`
	s.tj[22] = `Адади бутуни N (> 1) ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Агар маҷмӯи мазкур пайдарпайии камшавандаро ташкил диҳад, пас 0(нул)-ро хориҷ кунед; дар ҳолати акс рақами адади аввалинеро, ки қонуниятро вайрон мекунад, хориҷ кунед.`
	s.tj[23] = `Адади бутуни N (> 2) ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд. Маҷмӯъ аррашакл номида мешавад, агар ҳар як элементи дохилии он аз ҳар ду ҳамсояҳои худ ё калон, ё хурд (яъне "дандона" аст) бошад. Агар маҷмӯи мазкур аррашакл бошад, пас 0(нул)-ро хориҷ кунед; дар ҳолати акс рақами элементи аввалинро, ки дандона ба шумор намеравад, хориҷ кунед.`
	s.tj[24] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун, ки ба ҳадди имкон ду то нул дорад, дода шудаанд. Суммаи ададҳоро аз маҷмӯи мазкур, ки дар байни ду нулҳои охирин ҷойгир шудаанд, хориҷ кунед (агар нулҳо ҳамсоя бошанд, пас 0(нул)-ро хориҷ кунед).`
	s.tj[25] = `Адади бутуни N ва маҷмӯъ аз N ададҳои бутун, ки ба ҳадди имкон ду то нул дорад, дода шудаанд. Суммаи ададҳоро аз маҷмӯи мазкур, ки дар байни нулҳои аввалин ва охирин ҷойгир шудаанд, хориҷ кунед (агар нулҳои аввалин ва охирин ҳамсоя бошанд, пас 0(нул)-ро хориҷ кунед).`
	s.tj[26] = `Ададҳои бутуни K, N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд: A<sub>1</sub>, A<sub>2</sub>, …, A<sub>N</sub>. Дараҷаи K-ӯми ададҳоро аз маҷмӯи мазкур хориҷ кунед: (A<sub>1</sub>)<sup>K</sup>, (A<sub>2</sub>)<sup>K</sup>, …, (A<sub>N</sub>)<sup>K</sup>.`
	s.tj[27] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд: A<sub>1</sub>, A<sub>2</sub>, …, A<sub>N</sub>. Ададҳои зеринро хориҷ кунед: A<sub>1</sub>, (A<sub>2</sub>)<sup>2</sup>, …, (A<sub>N-1</sub>)<sup>N-1</sup>, (A<sub>N</sub>)<sup>N</sup>.`
	s.tj[28] = `Адади бутуни N ва маҷмӯъ аз N ададҳои ҳақиқӣ дода шудаанд: A<sub>1</sub>, A<sub>2</sub>, …, A<sub>N</sub>. Ададҳои зеринро хориҷ кунед: (A<sub>1</sub>)<sup>N</sup>, (A<sub>2</sub>)<sup>N-1</sup>, …, (A<sub>N-1</sub>)<sup>2</sup>, A<sub>N</sub>.`
	s.tj[29] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Суммаи умумии ҳамаи элементҳои ба маҷмӯъҳои мазкур дохилгардандаро хориҷ кунед.`
	s.tj[30] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Барои ҳар як маҷмӯъ суммаи элементҳояшро хориҷ кунед.`
	s.tj[31] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Миқдори маҷмӯъҳоеро, ки адади 2-ро доростанд, ёбед. Агар чунин маҷмӯъ набошад, пас 0(нул)-ро хориҷ кунед.`
	s.tj[32] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Барои ҳар як маҷмӯъ рақами элементи аввалини онро, ки ба 2 баробар аст, хориҷ кунед, агар дар маҷмӯи мазкур адади ду набошад, адади 0(нул)-ро хориҷ кунед.`
	s.tj[33] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Барои ҳар як маҷмӯъ рақами элементи охирини онро, ки ба 2 баробар аст, хориҷ кунед, агар дар маҷмӯи мазкур адади ду набошад, адади 0(нул)-ро хориҷ кунед.`
	s.tj[34] = `Ададҳои бутуни K, N, ҳамчунин K то маҷмӯъҳои ададҳои бутун бо N тогӣ элементҳо дар ҳар як маҷмӯъ дода шудаанд. Барои ҳар як маҷмӯъ амалҳои зеринро иҷро кунед: агар дар маҷмӯъ адади 2 бошад, пас суммаи элементҳои онро хориҷ кунед; агар дар маҷмӯъ адади ду набошад, пас 0(нул)-ро хориҷ кунед.`
	s.tj[35] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Нишонаи баанҷомрасии ҳар як маҷмӯъ адади 0(нул) ба ҳисоб меравад. Барои ҳар як маҷмӯъ миқдори элементҳояшро хориҷ кунед. Ҳамчунин миқдори умумии элементҳоро дар ҳамаи маҷмӯъҳо хориҷ кунед.`
	s.tj[36] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Ҳар як маҷмӯъ на кам аз ду элемент дорад, ки нишонаи баанҷомрасии он адади 0(нул) ба ҳисоб меравад. Миқдори маҷмӯъҳоеро, ки элементҳои онҳо афзуншаванда ҳастанд, ёбед.`
	s.tj[37] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Ҳар як маҷмӯъ на кам аз ду элемент дорад, ки нишонаи баанҷомрасии он адади 0(нул) ба ҳисоб меравад. Миқдори маҷмӯъҳоеро, ки элементҳояшон афзуншаванда ва ё камшаванда ҳастанд, ёбед.`
	s.tj[38] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Ҳар як маҷмӯъ на кам аз ду элемент дорад, ки нишонаи баанҷомрасии он адади 0(нул) ба ҳисоб меравад. Барои ҳар як маҷмӯъ амали зеринро иҷро кунед: агар элементҳои маҷмӯъ афзуншаванда бошанд, пас 1-ро хориҷ кунед, агар элементҳои маҷмӯъ камшаванда бошанд, пас (-1)-ро хориҷ кунед; агар элементҳои маҷмӯъ на афзуншаванда ва на камшаванда бошанд, пас 0(нул)-ро хориҷ кунед.`
	s.tj[39] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Ҳар як маҷмӯъ на кам аз се элемент дорад, ки нишонаи баанҷомрасии он адади 0(нул) ба ҳисоб меравад. Миқдори маҷмӯъҳои аррашаклро (муайянкунии маҷмӯи аррашакл дар масъалаи Series23 оварда шудааст) ёбед.`
	s.tj[40] = `Адади бутуни K, ҳамчунин K то маҷмӯъҳои ададҳои бутуни ғайринулӣ дода шудаанд. Ҳар як маҷмӯъ на кам аз се элемент дорад, ки нишонаи баанҷомрасии он адади 0(нул) ба ҳисоб меравад. Барои ҳар як маҷмӯъ амали зеринро иҷро кунед: агар маҷмӯъ аррашакл бошад (ниг. ба масъалаи Series23), пас миқдори элементҳои онро хориҷ кунед; дар холати акс рақами элементи аввалини онро, ки дандона ба ҳисоб намеравад, хориҷ кунед.`
}

func (s *sseries) makeRu() {
	s.ru = make([]string, s.count)
	s.ru[1] = `Даны десять вещественных чисел. Найти их сумму.`
	s.ru[2] = `Даны десять вещественных чисел. Найти их произведение.`
	s.ru[3] = `Даны десять вещественных чисел. Найти их среднее арифметическое.`
	s.ru[4] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел. Вывести сумму и произведение чисел из данного набора.`
	s.ru[5] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;положительных вещественных чисел. Вывести в том же порядке целые части всех чисел из данного набора (как вещественные числа с нулевой дробной частью), а также сумму всех целых частей.`
	s.ru[6] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;положительных вещественных чисел. Вывести в том же порядке дробные части всех чисел из данного набора (как вещественные числа с нулевой целой частью), а также произведение всех дробных частей.`
	s.ru[7] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел. Вывести в том же порядке округленные значения  всех чисел из данного набора (как целые числа), а также сумму всех округленных значений.`
	s.ru[8] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Вывести в том же порядке все четные числа из данного набора и количество&nbsp;<i>K</i> таких чисел.`
	s.ru[9] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Вывести в том же порядке номера всех нечетных чисел из данного набора и количество&nbsp;<i>K</i> таких чисел.`
	s.ru[10] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Если в наборе имеются положительные числа, то вывести true; в противном случае вывести false.`
	s.ru[11] = `Даны целые числа&nbsp;<i>K</i>,&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел. Если в наборе имеются числа, меньшие&nbsp;<i>K</i>, то вывести true; в противном случае вывести false.`
	s.ru[12] = `Дан набор ненулевых целых чисел; признак его завершения &#8212; число&nbsp;0. Вывести количество чисел в наборе.`
	s.ru[13] = `Дан набор ненулевых целых чисел; признак его завершения &#8212; число&nbsp;0. Вывести сумму всех положительных четных чисел из данного набора. Если требуемые числа в наборе отсутствуют, то вывести&nbsp;0.`
	s.ru[14] = `Дано целое число&nbsp;<i>K</i> и набор ненулевых целых чисел; признак его завершения &#8212; число&nbsp;0. Вывести количество чисел в наборе, меньших&nbsp;<i>K</i>.`
	s.ru[15] = `Дано целое число&nbsp;<i>K</i> и набор ненулевых целых чисел; признак его завершения &#8212; число&nbsp;0. Вывести номер первого числа в наборе, большего&nbsp;<i>K</i>. Если таких чисел нет, то вывести&nbsp;0.`
	s.ru[16] = `Дано целое число&nbsp;<i>K</i> и набор ненулевых целых чисел; признак его завершения &#8212; число&nbsp;0. Вывести номер последнего числа в наборе, большего&nbsp;<i>K</i>. Если таких чисел нет, то вывести&nbsp;0.`
	s.ru[17] = `Дано вещественное число&nbsp;<i>B</i>, целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел, упорядоченных по возрастанию. Вывести элементы набора вместе с числом&nbsp;<i>B</i>, сохраняя упорядоченность выводимых чисел.`
	s.ru[18] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел, упорядоченный по возрастанию. Данный набор может содержать одинаковые элементы. Вывести в том же порядке все различные элементы данного набора.`
	s.ru[19] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;целых чисел. Вывести те элементы в наборе, которые меньше своего левого соседа, и количество&nbsp;<i>K</i> таких элементов.`
	s.ru[20] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;целых чисел. Вывести те элементы в наборе, которые меньше своего правого соседа, и количество&nbsp;<i>K</i> таких элементов.`
	s.ru[21] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;вещественных чисел. Проверить, образует ли данный набор возрастающую последовательность. Если образует, то вывести true, если нет &#8212; вывести false.`
	s.ru[22] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;1) и набор из <i>N</i>&nbsp;вещественных чисел. Если данный набор образует убывающую последовательность, то вывести&nbsp;0; в противном случае вывести номер первого числа, нарушающего закономерность.`
	s.ru[23] = `Дано целое число&nbsp;<i>N</i> (&gt;&nbsp;2) и набор из <i>N</i>&nbsp;вещественных чисел. Набор называется <i>пилообразным</i>, если каждый его внутренний элемент либо больше, либо меньше обоих своих соседей (т.&nbsp;е. является &#171;зубцом&#187;). Если данный набор является пилообразным, то вывести&nbsp;0; в противном случае вывести номер первого элемента, не являющегося зубцом.`
	s.ru[24] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел, содержащий по крайней мере два нуля. Вывести сумму чисел из данного набора, расположенных между последними двумя нулями (если последние нули идут подряд, то вывести&nbsp;0).`
	s.ru[25] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;целых чисел, содержащий по крайней мере два нуля. Вывести сумму чисел из данного набора, расположенных между первым и последним нулем (если первый и последний нули идут подряд, то вывести&nbsp;0).`
	s.ru[26] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел:&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub>. Вывести <i>K</i>-e степени чисел из данного набора: (<i>A</i><sub>1</sub>)<sup><i>K</i></sup>, (<i>A</i><sub>2</sub>)<sup><i>K</i></sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i></sub>)<sup><i>K</i></sup>`
	s.ru[27] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел:&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub>. Вывести следующие числа:<i>A</i><sub>1</sub>, (<i>A</i><sub>2</sub>)<sup>2</sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i>&#8722;1</sub>)<sup><i>N</i>&#8722;1</sup>, (<i>A</i><sub><i>N</i></sub>)<sup><i>N</i></sup>`
	s.ru[28] = `Дано целое число&nbsp;<i>N</i> и набор из <i>N</i>&nbsp;вещественных чисел:&nbsp;<i>A</i><sub>1</sub>, <i>A</i><sub>2</sub>,&nbsp;&#8230;, <i>A</i><sub><i>N</i></sub>. Вывести следующие числа:(<i>A</i><sub>1</sub>)<sup><i>N</i></sup>, (<i>A</i><sub>2</sub>)<sup><i>N</i>&#8722;1</sup>,&nbsp;&#8230;, (<i>A</i><sub><i>N</i>&#8722;1</sub>)<sup>2</sup>, <i>A</i><sub><i>N</i></sub>`
	s.ru[29] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Вывести общую сумму всех элементов, входящих в данные наборы.`
	s.ru[30] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Для каждого набора вывести сумму его элементов.`
	s.ru[31] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Найти количество наборов, содержащих число&nbsp;2. Если таких наборов нет, то вывести&nbsp;0.`
	s.ru[32] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Для каждого набора вывести номер его первого элемента, равного&nbsp;2, или число&nbsp;0, если в данном наборе нет двоек.`
	s.ru[33] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Для каждого набора вывести номер его последнего элемента, равного&nbsp;2, или число&nbsp;0, если в данном наборе нет двоек.`
	s.ru[34] = `Даны целые числа&nbsp;<i>K</i>, <i>N</i>, а также <i>K</i>&nbsp;наборов целых чисел по <i>N</i>&nbsp;элементов в каждом наборе. Для каждого набора выполнить следующее действие: если в наборе содержится число&nbsp;2, то вывести сумму его элементов; если в наборе нет двоек, то вывести&nbsp;0.`
	s.ru[35] = `Дано целое число&nbsp;<i>K</i>, а также <i>K</i>&nbsp;наборов ненулевых целых чисел. Признаком завершения каждого набора является число&nbsp;0. Для каждого набора вывести количество его элементов. Вывести также общее количество элементов во всех наборах.`
	s.ru[36] = `Дано целое число&nbsp;<i>K</i>, а также&nbsp;<i>K</i> наборов ненулевых целых чисел. Каждый набор содержит не менее двух элементов, признаком его завершения является число&nbsp;0. Найти количество наборов, элементы которых возрастают.`
	s.ru[37] = `Дано целое число&nbsp;<i>K</i>, а также&nbsp;<i>K</i> наборов ненулевых целых чисел. Каждый набор содержит не менее двух элементов, признаком его завершения является число&nbsp;0. Найти количество наборов, элементы которых возрастают или убывают.`
	s.ru[38] = `Дано целое число&nbsp;<i>K</i>, а также&nbsp;<i>K</i> наборов ненулевых целых чисел. Каждый набор содержит не менее двух элементов, признаком его завершения является число&nbsp;0. Для каждого набора выполнить следующее действие: если элементы набора возрастают, то вывести&nbsp;1; если элементы набора убывают, то вывести&nbsp;&#8722;1; если элементы набора не возрастают и не убывают, то вывести&nbsp;0.`
	s.ru[39] = `Дано целое число&nbsp;<i>K</i>, а также&nbsp;<i>K</i> наборов ненулевых целых чисел. Каждый набор содержит не менее трех элементов, признаком его завершения является число&nbsp;0. Найти количество пилообразных наборов (определение пилообразного набора дано в задании Series23).`
	s.ru[40] = `Дано целое число&nbsp;<i>K</i>, а также&nbsp;<i>K</i> наборов ненулевых целых чисел. Каждый набор содержит не менее трех элементов, признаком его завершения является число&nbsp;0. Для каждого набора выполнить следующее действие: если набор является пилообразным (см. задание Series23), то вывести количество его элементов; в противном случае вывести номер первого элемента, который не является зубцом.`
}
