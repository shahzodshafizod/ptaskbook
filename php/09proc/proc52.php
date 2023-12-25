<?php
function IsLeapYear($Y){
        if(($Y % 400) == 0)
                $mant = true;
        else if(($Y % 100) == 0)
                $mant = false;
        else if(($Y % 4) == 0)
                $mant = true;
        else
                $mant = false;
        return $mant;
}
function boolval($nat){
	return ($nat? 'true': 'false');
}
$hisob = 1;
while($hisob <= 5){
	$sol = mt_rand(1, 9999);
	$nat = IsLeapYear($sol);
	echo "Соли додашуда: ".$sol;
	echo "<br>Қабисавӣ будани соли додашуда: ".boolval($nat)."<br><br>";
	$hisob++;
}
?>