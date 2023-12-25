<?php
$fact = 1;
$N = mt_rand(1, 50);
$X = mt_rand(1, 100) / 4;
$X = number_format($X, 2);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X."<br>";
$sum = $X;
for($i = 1; $i <= $N; $i++){
	$fact *= (2*$i);
	$fact *= (2*$i+1);
	$sum += pow(-1, $i) * pow($X, 2 * $N + 1) / $fact;
}
echo "Сумма = ".$sum;
?>