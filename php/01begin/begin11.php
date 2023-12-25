<?php
do{
	$a = mt_rand(-100, 100) / 4;
	$b = mt_rand(-100, 100) / 4;
}
while(($a == 0) || ($b == 0));
$a = number_format($a, 2);
$b = number_format($b, 2);
echo "Первое число: ".$a;
echo "<br>Второе число: ".$b;
$a = abs($a);
$b = abs($b);
$jamm = $a + $b;
$tarh = $a - $b;
$zarb = $a * $b;
$taqs = $a / $b;
$jamm = number_format($jamm, 2);
$tarh = number_format($tarh, 2);
$zarb = number_format($zarb, 2);
$taqs = number_format($taqs, 2);
echo "<br>Сумма модулей чисел: ".$jamm;
echo "<br>Разность модулей чисел: ".$tarh;
echo "<br>Произведение модулей чисел: ".$zarb;
echo "<br>Частное модулей чисел: ".$taqs;
?>

/*
<?php
do{
	$a = mt_rand(-100, 100) / 4;
	$b = mt_rand(-100, 100) / 4;
}
while(($a == 0) || ($b == 0));
$a = number_format($a, 2);
$b = number_format($b, 2);
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
$a = abs($a);
$b = abs($b);
$jamm = $a + $b;
$tarh = $a - $b;
$zarb = $a * $b;
$taqs = $a / $b;
$jamm = number_format($jamm, 2);
$tarh = number_format($tarh, 2);
$zarb = number_format($zarb, 2);
$taqs = number_format($taqs, 2);
echo "<br>Ҷамъи ададҳо: ".$jamm;
echo "<br>Тарҳи ададҳо: ".$tarh;
echo "<br>Зарби ададҳо: ".$zarb;
echo "<br>Тақсими ададҳо: ".$taqs;
?>
*/