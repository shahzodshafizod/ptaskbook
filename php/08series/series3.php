<?php
$sum = 0;
$hisob = 1;
while($hisob <= 10){
	$N = mt_rand(1, 100);
	echo "N".$hisob." = ".$N."<br>";
	$sum += $N;
	$hisob++;
}
echo "Қимати миёнаи арифметикӣ:".number_format($sum / 10, 2);
?>