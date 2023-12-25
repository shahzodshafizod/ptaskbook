<?php
function AddRightDigit($D, &$K){
        $K = $K * 10 + $D;
}
$hisob = 1;
$N = mt_rand(1, 10000);
echo "Адади додашуда: ".$N;
while($hisob <= 2){
	$d = mt_rand(0, 9);
	echo "<br>d".$hisob." = ".$d;
	AddRightDigit($d, $N);
	echo "<br>Натиҷаи ".$hisob.": ".$N."<br>";
	$hisob++;
}
?>