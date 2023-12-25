<?php
function Swap(&$X, &$Y){
        $Z = $X;
        $X = $Y;
        $Y = $Z;
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
$D = mt_rand(-100, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>D = ".$D;
Swap($A, $B);
Swap($C, $D);
Swap($B, $C);
echo "<br>Пас аз иҷрои амалиёт:";
echo "<br>A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>D = ".$D;
?>