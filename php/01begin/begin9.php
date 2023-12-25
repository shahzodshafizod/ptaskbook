<?php
$a = mt_rand(1, 100) / 4;
$b = mt_rand(1, 100) / 4;
$c = sqrt($a * $b);
$a = number_format($a, 2);
$b = number_format($b, 2);
$c = number_format($c, 2);
echo "Первое число: ".$a;
echo "<br>Второе число: ".$b;
echo "<br>Среднее геометрическое: ".$c;
?>

/*
<?php
$a = mt_rand(1, 100) / 4;
$b = mt_rand(1, 100) / 4;
$c = sqrt($a * $b);
$a = number_format($a, 2);
$b = number_format($b, 2);
$c = number_format($c, 2);
echo "Адади якӯм: ".$a;
echo "<br>Адади дуюм: ".$b;
echo "<br>Қимати миёнаи геометрӣ: ".$c;
?>
*/