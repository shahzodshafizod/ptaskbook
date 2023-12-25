﻿<?php
function IsLeapYear($Y){
        $mant = false;
        if((($Y % 400) == 0) && (($Y % 4) == 0))
                $mant = true;
        return $mant;
}
function MonthDay($M, $Y){
        $ruz = 0;
        switch($M){
                case 1:  case 3:
                case 5:  case 7:
                case 8:  case 10:
                case 12:
                        $ruz = 31;
                        break;
                case 4:  case 6:
                case 9:  case 11:
                        $ruz = 30;
                        break;
                case 2:
                        if(IsLeapYear($Y))
                                $ruz = 29;
                        else
                                $ruz = 28;
        }
        return $ruz;
}
function NextDate(&$D, &$M, &$Y){
        if($D == MonthDay($M, $Y)){
                $D = 1;
                if($M == 12){
                        $M = 1;
                        $Y++;
                }
                else
                        $M++;
        }
        else
                $D++;
}
$hisob = 1;
while($hisob <= 3){
	$sol = mt_rand(1, 9999);
	$moh = mt_rand(1, 12);
	do{
		$ruz = mt_rand(1, 31);
	}
	while(($moh==2)&&((IsLeapYear($sol))&&($ruz>29)||(!IsLeapYear($sol)&&($ruz>28))) || 
	($moh==4)&&($moh==6)&&($moh==9)&&($moh==11)&&($ruz>30));
	echo "Санаи додашуда:";
	echo "<br>".$ruz." - ".$moh." - ".$sol;
	NextDate($ruz, $moh, $sol);
	echo "<br>Санаи як рӯз пac:";
	echo "<br>".$ruz." - ".$moh." - ".$sol."<br><br>";
	$hisob++;
}
?>