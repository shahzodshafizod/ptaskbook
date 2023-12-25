<?php
$sum = 0;
$hisob = 1;
while($hisob <= 10){
	$N = mt_rand(1, 100) / 4;
	$N = number_format($N, 2);
	echo "N".$hisob." = ".$N."<br>";
	$sum += $N;
	$hisob++;
}
echo "Ҷамъи ададҳо = ".$sum;
?>

/*
<?php
	$sum = 0;
	for ($i = 1; $i <= 10; ++$i) {
		$number = floatval(readline("number$i = "));
		$sum += $number;
	}
    echo sprintf('%.2f', $sum);
?>
*/