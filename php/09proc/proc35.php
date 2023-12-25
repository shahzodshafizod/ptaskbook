<?php
function Fact2($N){
        $fact2 = 1;
        while($N >= 1){
                $fact2 *= $N;
                $N -= 2;
        }
        return $fact2;
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand(1, 24);
	$nat = Fact2($N);
	echo $N."!! = ".$nat."<br>";
	$hisob++;
}
?>