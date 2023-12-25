<?php
$fact = 1;
$sum = 1;
$N = mt_rand(1, 50);
$X = mt_rand(1, 100) / 4;
$X = number_format($X, 2);
echo "Ҳудуди давр: ".$N;
echo "<br>Адади ҳақиқӣ: ".$X."<br>";
for($i = 1; $i <= $N; $i++){
	$fact *= $i;
	$sum += pow($X, $i) / $fact;
}
echo "Сумма = ".number_format($sum, 2);
echo "<br>exp(".$X.") = ".number_format(exp($X), 2);
?>