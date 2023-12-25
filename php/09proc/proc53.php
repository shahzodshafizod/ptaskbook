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
$hisob = 1;
$sol = mt_rand(1, 9999);
echo "Сол: ".$sol."<br>";
while($hisob <= 3){
	$moh = mt_rand(1, 12);
	$nat = MonthDay($moh, $sol);
	echo "<br>Моҳ: ".$moh;
	echo "<br>Рӯз: ".$nat."<br>";
	$hisob++;
}
?>