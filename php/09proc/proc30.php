<?php
function DigitN($K, $N){
        $i = 1;
        while(($i <= $N) && ($K > 0)){
                $raq = $K % 10;
                $K = intval($K / 10);
                $i++;
        }
        if($i - 1 == $N)
                return $raq;
        else
                return -1;
}
$hisob = 1;
$N = mt_rand(1, 5);
echo "N = ".$N."<br>";
while($hisob <= 5){
	$K = mt_rand(1, 1000000);
	echo "K".$hisob." = ".$K."<br>";
	$raqam = DigitN($K, $N);
	echo "Миқдори рақамҳо: ".$raqam."<br><br>";
	$hisob++;
}
?>