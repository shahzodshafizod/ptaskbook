package task

type bboolean struct {
	count int
	name  string
	en    []string
	tj    []string
	ru    []string
}

func NewBooleanTask() Task {
	b := &bboolean{
		count: 41,
		name:  "Boolean",
	}
	b.makeEn()
	b.makeTj()
	b.makeRu()
	return b
}

func (b *bboolean) Count() int   { return b.count }
func (b *bboolean) Name() string { return b.name }
func (b *bboolean) En() []string { return b.en }
func (b *bboolean) Tj() []string { return b.tj }
func (b *bboolean) Ru() []string { return b.ru }

func (b *bboolean) makeEn() {
	b.en = make([]string, b.count)
	b.en[1] = `Given integer&nbsp;<i>A</i>, verify the following proposition: &#34;The number&nbsp;<i>A</i> is positive&#34;.`
	b.en[2] = `Given integer&nbsp;<i>A</i>, verify the following proposition: &#34;The number&nbsp;<i>A</i> is odd&#34;.`
	b.en[3] = `Given integer&nbsp;<i>A</i>, verify the following proposition: &#34;The number&nbsp;<i>A</i> is even&#34;.`
	b.en[4] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;The inequalities <i>A</i>&nbsp;&gt;&nbsp;2 and&nbsp;<i>B</i>&nbsp;&#8804;&nbsp;3 both are fulfilled&#34;.`
	b.en[5] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;The inequality <i>A</i>&nbsp;&#8805;&nbsp;0 is fulfilled or the inequality <i>B</i>&nbsp;&lt;&nbsp;&#8722;2 is fulfilled&#34;.`
	b.en[6] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;The double inequality <i>A</i>&nbsp;&lt;&nbsp;<i>B</i>&nbsp;&lt;&nbsp;<i>C</i> is fulfilled&#34;.`
	b.en[7] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;The number&nbsp;<i>B</i> is between&nbsp;<i>A</i> and&nbsp;<i>C</i>&#34;.`
	b.en[8] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;Each of the numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> is odd&#34;.`
	b.en[9] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;At least one of the numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> is odd&#34;.`
	b.en[10] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;Exactly one of the numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> is odd&#34;.`
	b.en[11] = `Given two integers&nbsp;<i>A</i> and&nbsp;<i>B</i>, verify the following proposition: &#34;The numbers&nbsp;<i>A</i> and&nbsp;<i>B</i> have equal parity&#34;.`
	b.en[12] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;Each of the numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> is positive&#34;.`
	b.en[13] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;At least one of the numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> is positive&#34;.`
	b.en[14] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;Exactly one of the numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> is positive&#34;.`
	b.en[15] = `Given three integers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i>, verify the following proposition: &#34;Exactly two of the numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are positive&#34;.`
	b.en[16] = `Given a positive integer, verify the following proposition: &#34;The integer is a two-digit even number&#34;.`
	b.en[17] = `Given a positive integer, verify the following proposition: &#34;The integer is a three-digit odd number&#34;.`
	b.en[18] = `Verify the following proposition: &#34;Among three given integers there is at least one pair of equal ones&#34;.`
	b.en[19] = `Verify the following proposition: &#34;Among three given integers there is at least one pair of opposite ones&#34;.`
	b.en[20] = `Given a three-digit integer, verify the following proposition: &#34;All digits of the number are different&#34;.`
	b.en[21] = `Given a three-digit integer, verify the following proposition: &#34;All digits of the number are in ascending order&#34;.`
	b.en[22] = `Given a three-digit integer, verify the following proposition: &#34;All digits of the number are in ascending or descending order&#34;.`
	b.en[23] = `Given a four-digit integer, verify the following proposition: &#34;The number is read equally both from left to right and from right to left&#34;.`
	b.en[24] = `Three real numbers&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> are given (<i>A</i> is not equal to&nbsp;0). By means of a <i>discriminant</i> <i>D</i>&nbsp;=&nbsp;<i>B</i><sup>2</sup>&nbsp;&#8722;&nbsp;4&#183;<i>A</i>&#183;<i>C</i>, verify the following proposition: &#34;The quadratic equation <i>A</i>&#183;<i>x</i><sup>2</sup>&nbsp;+&nbsp;<i>B</i>&#183;<i>x</i>&nbsp;+&nbsp;<i>C</i>&nbsp;=&nbsp;0 has real roots&#34;.`
	b.en[25] = `Given two real numbers&nbsp;<i>x</i>, <i>y</i>, verify the following proposition: &#34;The point with coordinates (<i>x</i>,&nbsp;<i>y</i>) is in the second coordinate quarter&#34;.`
	b.en[26] = `Given two real numbers&nbsp;<i>x</i>, <i>y</i>, verify the following proposition: &#34;The point with coordinates (<i>x</i>,&nbsp;<i>y</i>) is in the fourth coordinate quarter&#34;.`
	b.en[27] = `Given two real numbers&nbsp;<i>x</i>, <i>y</i>, verify the following proposition: &#34;The point with coordinates (<i>x</i>,&nbsp;<i>y</i>) is in the second or third coordinate quarter&#34;.`
	b.en[28] = `Given two real numbers&nbsp;<i>x</i>, <i>y</i>, verify the following proposition: &#34;The point with coordinates (<i>x</i>,&nbsp;<i>y</i>) is in the first or third coordinate quarter&#34;.`
	b.en[29] = `Given real numbers&nbsp;<i>x</i>, <i>y</i>, <i>x</i><sub>1</sub>,&nbsp;<i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub>, verify the following proposition: &#34;The point (<i>x</i>,&nbsp;<i>y</i>) is inside of the rectangle whose left top vertex is (<i>x</i><sub>1</sub>,&nbsp;<i>y</i><sub>1</sub>), right bottom vertex is (<i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub>), and sides are parallel to coordinate axes&#34;.`
	b.en[30] = `Given three integers&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> that are the sides of a triangle, verify the following proposition: &#34;The triangle with sides&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> is equilateral&#34;.`
	b.en[31] = `Given three integers&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> that are the sides of a triangle, verify the following proposition: &#34;The triangle with sides&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> is isosceles&#34;.`
	b.en[32] = `Given three integers&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> that are the sides of a triangle, verify the following proposition: &#34;The triangle with sides&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> is a right triangle&#34;.`
	b.en[33] = `Given three integers&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i>, verify the following proposition: &#34;A triangle with the sides&nbsp;<i>a</i>, <i>b</i>,&nbsp;<i>c</i> exists&#34;.`
	b.en[34] = `Given coordinates&nbsp;<i>x</i>, <i>y</i> of a chessboard square (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;The chessboard square (<i>x</i>,&nbsp;<i>y</i>) is white&#34;. Note that the left bottom square (1,&nbsp;1) is black.`
	b.en[35] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;Both of the given chessboard squares have the same color&#34;.`
	b.en[36] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;A rook can move from one square to another during one turn&#34;.`
	b.en[37] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;A king can move from one square to another during one turn&#34;.`
	b.en[38] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;A bishop can move from one square to another during one turn&#34;.`
	b.en[39] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;A queen can move from one square to another during one turn&#34;.`
	b.en[40] = `Given coordinates&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> of two chessboard squares (as integers in the range&nbsp;1 to&nbsp;8), verify the following proposition: &#34;A knight can move from one square to another during one turn&#34;.`
}

func (b *bboolean) makeTj() {
	b.tj = make([]string, b.count)
	b.tj[1] = `Адади бутун A дода шудааст. Дурустии гуфтори: «Адади A мусбӣ аст»-ро санҷед.`
	b.tj[2] = `Адади бутун A дода шудааст. Дурустии гуфтори: «Адади A тоқ аст»-ро санҷед.`
	b.tj[3] = `Адади бутун A дода шудааст. Дурустии гуфтори: «Адади A ҷуфт аст»-ро санҷед.`
	b.tj[4] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Нобаробариҳои A>2 ва B<=3 дурустанд»-ро санҷед.`
	b.tj[5] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Нобаробариҳои A>=0 ё B<-2 дурустанд»-ро санҷед.`
	b.tj[6] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Нобаробарии дукаратаи A&lt;B&lt;C дуруст аст»-ро санҷед.`
	b.tj[7] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Адади B дар байни ададҳои A ва C ҷойгир аст»-ро санҷед.`
	b.tj[8] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Ҳар яке аз ададҳои A ва B тоқ аст»-ро санҷед.`
	b.tj[9] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Аққалан яке аз ададҳои A ва B тоқ аст»-ро санҷед.`
	b.tj[10] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Расо яке аз ададҳои A ва B тоқ аст»-ро санҷед.`
	b.tj[11] = `Ду ададҳои бутун дода шудаанд: A, B. Дурустии гуфтори: «Ададҳои A ва B ҷуфтии якхела доранд»-ро санҷед.`
	b.tj[12] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Ҳар яке аз ададҳои A, B, C мусбат аст»-ро санҷед.`
	b.tj[13] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Аққалан яке аз ададҳои A, B, C мусбат аст»-ро санҷед.`
	b.tj[14] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Расо яке аз ададҳои A, B, C мусбат аст»-ро санҷед.`
	b.tj[15] = `Се ададҳои бутун дода шудаанд: A, B, C. Дурустии гуфтори: «Расо дуто аз ададҳои A, B, C мусбатанд»-ро санҷед.`
	b.tj[16] = `Адади бутуни мусбӣ дода шудааст. Дурустии гуфтори: «Адади мазкур адади ҷуфти дурақама аст»-ро санҷед.`
	b.tj[17] = `Адади бутуни мусбӣ дода шудааст. Дурустии гуфтори: «Адади мазкур адади тоқи серақама аст»-ро санҷед.`
	b.tj[18] = `Дурустии гуфтори: «Дар байни се ададҳои додашудаи бутун аққалан як ҷуфти ададҳои мувофиқоянда ҳаст»-ро санҷед.`
	b.tj[19] = `Дурустии гуфтори: «Дар байни се ададҳои додашудаи бутун аққалан як ҷуфти ададҳои дутарафа муқобил ҳаст»-pо санҷед.`
	b.tj[20] = `Адади серақама дода шудааст. Дурустии гуфтори: «Ҳамаи рақамҳои адади мазкур гуногунанд»-ро санҷед.`
	b.tj[21] = `Адади серақама дода шудааст. Дурустии гуфтори: «Рақамҳои адади мазкур пайдарпайии афзуншавандаро ташкил медиҳанд»-ро санҷед.`
	b.tj[22] = `Адади серақама дода шудааст. Дурустии гуфтори: «Рақамҳои адади мазкур пайдарпайии камшаванда ё афзуншавандаро ташкил медиҳанд»-ро санҷед.`
	b.tj[23] = `Адади чоррақама дода шудааст. Дурустии гуфтори: «Адади мазкур аз чап ба рост ва аз рост ба чап якхела хонда мешавад»-ро санҷед.`
	b.tj[24] = `Ададҳои A, B, C (адади A НОбаробари 0 аст) дода шудаанд. Дискриминант d=b<sup>2</sup>–4*a*c-ро аз назар гузаронида, дурустии гуфтори: «Муодилаи квадратии a*x<sup>2</sup>+b*x+c=0 решаҳои ҳақиқӣ дорад»-ро санҷед.`
	b.tj[25] = `Ададҳои x, y дода шудаанд. Дурустии гуфтори: «Нуқта бо координатаҳои (x,y) дар чоряки координатавии дуюм мехобад»-ро санҷед.`
	b.tj[26] = `Ададҳои x, y дода шудаанд. Дурустии гуфтори: «Нуқта бо координатаҳои (x,y) дар чоряки координатавии чорӯм мехобад»-ро санҷед.`
	b.tj[27] = `Ададҳои x, y дода шудаанд. Дурустии гуфтори: «Нуқта бо координатаҳои (x,y) дар чоряки координатавии дуюм ё сеюм мехобад»-ро санҷед.`
	b.tj[28] = `Ададҳои x, y дода шудаанд. Дурустии гуфтори: «Нуқта бо координатаҳои (x,y) дар чоряки координатавии якӯм ё сеюм мехобад»-ро санҷед.`
	b.tj[29] = `Ададҳои x, y, x<sub>1</sub>, y<sub>1</sub>, x<sub>2</sub>, y<sub>2</sub> дода шудаанд. Дурустии гуфтори: «Нуқта бо координатаҳои (x,y) дар дохили росткунҷа мехобад, ки қуллаи чап-болоии он координатаҳои (x<sub>1</sub>,y<sub>1</sub>), рост-поёниаш — (x<sub>2</sub>,y<sub>2</sub>)-ро доро буда, тарафҳо ба тирҳои координатавӣ параллел мебошанд»-ро санҷед.`
	b.tj[30] = `Ададҳои бутун a, b, c, ки тарафҳои секунҷа ба ҳисоб мераванд, дода шудаанд. Дурустии гуфтори: «Секунҷа бо тарафҳои a, b, c секунҷаи баробартараф аст»-ро санҷед.`
	b.tj[31] = `Ададҳои бутун a, b, c, ки тарафҳои секунҷа ба ҳисоб мераванд, дода шудаанд. Дурустии гуфтори: «Секунҷа бо тарафҳои a, b, c секунҷаи баробарпаҳлӯ аст»-ро санҷед.`
	b.tj[32] = `Ададҳои бутун a, b, c, ки тарафҳои секунҷа ба ҳисоб мераванд, дода шудаанд. Дурустии гуфтори: «Секунҷа бо тарафҳои a, b, c секунҷаи росткунҷа аст»-ро санҷед.`
	b.tj[33] = `Ададҳои бутун a, b, c дода шудаанд. Дурустии гуфтори: «Секунҷа бо тарафҳои a, b, c вуҷуд дорад»-ро санҷед.`
	b.tj[34] = `Координатаҳои майдони тахтаи шоҳмот x, y (ададҳои бутун, дар фосилаи 1–8) дода шудаанд. Бо назардошти он, ки майдони чап-поёнии тахта (1,1) сиёҳ аст, дурустии гуфтори: «Майдони мазкур сафед аст»-ро санҷед.`
	b.tj[35] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Майдонҳои мазкур ранги якхела доранд»-ро санҷед.`
	b.tj[36] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Рух бо як ҳаракат аз як майдон ба майдони дигарӣ кӯчида метавонад»-ро санҷед.`
	b.tj[37] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Шоҳ бо як ҳаракат аз як майдон ба майдони дигарӣ кӯчида метавонад»-ро санҷед.`
	b.tj[38] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Фил бо як ҳаракат аз як майдон ба майдони дигарӣ кӯчида метавонад»-ро санҷед.`
	b.tj[39] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Фарзин бо як ҳаракат аз як майдон ба майдони дигарӣ кӯчида метавонад»-ро санҷед.`
	b.tj[40] = `Координатаҳои ду майдонҳои гуногуни тахтаи шоҳмот x<sub>1</sub>,y<sub>1</sub>,x<sub>2</sub>,y<sub>2</sub> (ададҳои бутуни дар фосилаи 1-8 хобанда) дода шудаанд. Дурустии гуфтори: «Асп бо як ҳаракат аз як майдон ба майдони дигарӣ кӯчида метавонад»-ро санҷед.`
}

func (b *bboolean) makeRu() {
	b.ru = make([]string, b.count)
	b.ru[1] = `Дано целое число&nbsp;<i>A</i>. Проверить истинность высказывания: &#171;Число&nbsp;<i>A</i> является положительным&#187;.`
	b.ru[2] = `Дано целое число&nbsp;<i>A</i>. Проверить истинность высказывания: &#171;Число&nbsp;<i>A</i> является нечетным&#187;.`
	b.ru[3] = `Дано целое число&nbsp;<i>A</i>. Проверить истинность высказывания: &#171;Число&nbsp;<i>A</i> является четным&#187;.`
	b.ru[4] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Справедливы неравенства <i>A</i>&nbsp;&gt;&nbsp;2 и&nbsp;<i>B</i>&nbsp;&#8804;&nbsp;3&#187;.`
	b.ru[5] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Справедливы неравенства <i>A</i>&nbsp;&#8805;&nbsp;0 или <i>B</i>&nbsp;&lt;&nbsp;&#8722;2&#187;.`
	b.ru[6] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Справедливо двойное неравенство <i>A</i>&nbsp;&lt;&nbsp;<i>B</i>&nbsp;&lt;&nbsp;<i>C</i>&#187;.`
	b.ru[7] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Число&nbsp;<i>B</i> находится между числами&nbsp;<i>A</i> и&nbsp;<i>C</i>&#187;.`
	b.ru[8] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Каждое из чисел&nbsp;<i>A</i> и&nbsp;<i>B</i> нечетное&#187;.`
	b.ru[9] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Хотя бы одно из чисел&nbsp;<i>A</i> и&nbsp;<i>B</i> нечетное&#187;.`
	b.ru[10] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Ровно одно из чисел&nbsp;<i>A</i> и&nbsp;<i>B</i> нечетное&#187;.`
	b.ru[11] = `Даны два целых числа: <i>A</i>,&nbsp;<i>B</i>. Проверить истинность высказывания: &#171;Числа&nbsp;<i>A</i> и&nbsp;<i>B</i> имеют одинаковую четность&#187;.`
	b.ru[12] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Каждое из чисел&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> положительное&#187;.`
	b.ru[13] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Хотя бы одно из чисел&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> положительное&#187;.`
	b.ru[14] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Ровно одно из чисел&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> положительное&#187;.`
	b.ru[15] = `Даны три целых числа: <i>A</i>, <i>B</i>,&nbsp;<i>C</i>. Проверить истинность высказывания: &#171;Ровно два из чисел&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> являются положительными&#187;.`
	b.ru[16] = `Дано целое положительное число. Проверить истинность высказывания: &#171;Данное число является четным двузначным&#187;.`
	b.ru[17] = `Дано целое положительное число. Проверить истинность высказывания: &#171;Данное число является нечетным трехзначным&#187;.`
	b.ru[18] = `Проверить истинность высказывания: &#171;Среди трех данных целых чисел есть хотя бы одна пара совпадающих&#187;.`
	b.ru[19] = `Проверить истинность высказывания: &#171;Среди трех данных целых чисел есть хотя бы одна пара взаимно противоположных&#187;.`
	b.ru[20] = `Дано трехзначное число. Проверить истинность высказывания: &#171;Все цифры данного числа различны&#187;.`
	b.ru[21] = `Дано трехзначное число. Проверить истинность высказывания: &#171;Цифры данного числа образуют возрастающую последовательность&#187;.`
	b.ru[22] = `Дано трехзначное число. Проверить истинность высказывания: &#171;Цифры данного числа образуют возрастающую или убывающую последовательность&#187;.`
	b.ru[23] = `Дано четырехзначное число. Проверить истинность высказывания: &#171;Данное число читается одинаково слева направо и справа налево&#187;.`
	b.ru[24] = `Даны числа&nbsp;<i>A</i>, <i>B</i>,&nbsp;<i>C</i> (число&nbsp;<i>A</i> не равно&nbsp;0).  Рассмотрев <i>дискриминант</i> <i>D</i>&nbsp;=&nbsp;<i>B</i><sup>2</sup>&nbsp;&#8722;&nbsp;4&#183;<i>A</i>&#183;<i>C</i>, проверить истинность высказывания: &#171;Квадратное уравнение <i>A</i>&#183;<i>x</i><sup>2</sup>&nbsp;+&nbsp;<i>B</i>&#183;<i>x</i>&nbsp;+&nbsp;<i>C</i>&nbsp;=&nbsp;0 имеет вещественные корни&#187;.`
	b.ru[25] = `Даны числа&nbsp;<i>x</i>, <i>y</i>. Проверить истинность высказывания: &#171;Точка с координатами (<i>x</i>,&nbsp;<i>y</i>) лежит во второй координатной четверти&#187;.`
	b.ru[26] = `Даны числа&nbsp;<i>x</i>, <i>y</i>. Проверить истинность высказывания: &#171;Точка с координатами (<i>x</i>,&nbsp;<i>y</i>) лежит в четвертой координатной четверти&#187;.`
	b.ru[27] = `Даны числа&nbsp;<i>x</i>, <i>y</i>. Проверить истинность высказывания: &#171;Точка с координатами (<i>x</i>,&nbsp;<i>y</i>) лежит во второй или третьей координатной четверти&#187;.`
	b.ru[28] = `Даны числа&nbsp;<i>x</i>, <i>y</i>. Проверить истинность высказывания: &#171;Точка с координатами (<i>x</i>,&nbsp;<i>y</i>) лежит в первой или третьей координатной четверти&#187;.`
	b.ru[29] = `Даны числа&nbsp;<i>x</i>, <i>y</i>, <i>x</i><sub>1</sub>,&nbsp;<i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub>. Проверить истинность высказывания: &#171;Точка с координатами (<i>x</i>,&nbsp;<i>y</i>) лежит внутри прямоугольника, левая верхняя вершина которого имеет координаты (<i>x</i><sub>1</sub>,&nbsp;<i>y</i><sub>1</sub>), правая нижняя &#8212; (<i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub>), а стороны параллельны координатным осям&#187;.`
	b.ru[30] = `Даны целые числа <i>a</i>, <i>b</i>, <i>c</i>, являющиеся сторонами некоторого треугольника. Проверить истинность высказывания: &#171;Треугольник со сторонами <i>a</i>, <i>b</i>, <i>c</i> является равносторонним&#187;.`
	b.ru[31] = `Даны целые числа <i>a</i>, <i>b</i>, <i>c</i>, являющиеся сторонами некоторого треугольника. Проверить истинность высказывания: &#171;Треугольник со сторонами <i>a</i>, <i>b</i>, <i>c</i> является равнобедренным&#187;.`
	b.ru[32] = `Даны целые числа <i>a</i>, <i>b</i>, <i>c</i>, являющиеся сторонами некоторого треугольника. Проверить истинность высказывания: &#171;Треугольник со сторонами <i>a</i>, <i>b</i>, <i>c</i> является прямоугольным&#187;.`
	b.ru[33] = `Даны целые числа <i>a</i>, <i>b</i>, <i>c</i>. Проверить истинность высказывания: &#171;Существует треугольник со сторонами <i>a</i>, <i>b</i>, <i>c</i>&#187;.`
	b.ru[34] = `Даны координаты поля шахматной доски&nbsp;<i>x</i>, <i>y</i> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Учитывая, что левое нижнее поле доски (1,&nbsp;1) является черным, проверить истинность высказывания: &#171;Данное поле является белым&#187;.`
	b.ru[35] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Данные поля имеют одинаковый цвет&#187;.`
	b.ru[36] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Ладья за один ход может перейти с одного поля на другое&#187;.`
	b.ru[37] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Король за один ход может перейти с одного поля на другое&#187;.`
	b.ru[38] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Слон за один ход может перейти с одного поля на другое&#187;.`
	b.ru[39] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Ферзь за один ход может перейти с одного поля на другое&#187;.`
	b.ru[40] = `Даны координаты двух различных полей шахматной доски&nbsp;<i>x</i><sub>1</sub>, <i>y</i><sub>1</sub>, <i>x</i><sub>2</sub>,&nbsp;<i>y</i><sub>2</sub> (целые числа, лежащие в диапазоне&nbsp;1&#8211;8). Проверить истинность высказывания: &#171;Конь за один ход может перейти с одного поля на другое&#187;.`
}
