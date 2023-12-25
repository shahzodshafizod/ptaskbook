<?php
$fact = 1;
$sum = 0;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= $N; $i++){
	$fact *= $i;
	$sum += 1 / $fact;
}
echo "Сумма = ".number_format($sum, 2);
echo "<br>e(1) = ".number_format(exp(1), 2);
?>