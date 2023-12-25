<?php
function IsPowerN($K, $N){
		$dar = 1;
        while($dar < $K)
                $dar *= $N;
        $nat = ($dar == $K);
        return $nat;
}
$hisob = 1;
$miq = 0;
$N = mt_rand(1, 10);
echo "N = ".$N."<br>";
while($hisob <= 10){
	$K = mt_rand(1, 1000000);
	echo "K".$hisob." = ".$K."<br>";
	if(IsPowerN($K, $N))
		$miq++;
	$hisob++;
}
echo "Миқдори дараҷаҳои адади N: ".$miq;
?>