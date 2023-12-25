<?php
$fact = 1;
$sum = 1;
$N = mt_rand(1, 50);
$X = mt_rand(1, 100) / 4;
$X = number_format($X, 2);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X."<br>";
for($i = 1; $i <= $N; $i++){
	$fact *= (2 * $i - 1);
	$fact *= (2 * $i);
	$sum += pow(-1, $i) * pow($X, 2 * $i) / $fact;
}
echo "Сумма = ".number_format($sum, 2);
echo "<br>cos(".$X.") = ".number_format(cos($X), 2);
?>