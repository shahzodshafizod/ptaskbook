<?php
function Fib($N){
        $hisob = 3;
        $F1 = $F2 = 1;
        while($hisob <= $N){
                $F = $F1 + $F2;
                $F1 = $F2;
                $F2 = $F;
                $hisob++;
        }
        return $F;
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand(2, 100);
	$nat = Fib($N);
	echo "Элементи ".$N."-ӯми пайдарпайӣ ба ".$nat." баробар аст<br>";
	$hisob++;
}
?>