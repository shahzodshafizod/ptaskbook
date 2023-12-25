<?php
function AddLeftDigit($D, &$K){
        $adad = $K;
        while($adad > 0){
                $D *= 10;
                $adad = intval($adad / 10);
        }
        $K += $D;
}
$hisob = 1;
$N = mt_rand(1, 10000);
echo "Адади додашуда: ".$N;
while($hisob <= 2){
	$d = mt_rand(1, 9);
	echo "<br>d".$hisob." = ".$d;
	AddLeftDigit($d, $N);
	echo "<br>Натиҷаи ".$hisob.": ".$N."<br>";
	$hisob++;
}
?>