<?php
echo "Самтхо:<br>";
echo "N - Север<br>S - Юг<br>E - Восток<br>W - Запад<br>";
echo "Фармонхо:<br>";
echo "1 - Налево<br>-1 - Направо<br>0 - Продолжать<br>";
$mass = array('N', 'S', 'E', 'W');
$n = rand(0, 3);
$far = array(1, -1, 0);
$f = rand(0, 2);
switch($mass[$n]){
       case 'N':
             switch($far[$f]){
                    case 1:
                         $har = 'W';
                         break;
                    case -1:
                         $har = 'E';
                         break;
                    case 0:
                         $har = 'N';
             }
             break;
       case 'S':
             switch($far[$f]){
                    case 1:
                         $har = 'E';
                         break;
                    case -1:
                         $har = 'W';
                         break;
                    case 0:
                         $har = 'S';
             }
             break;
       case 'E':
             switch($far[$f]){
                    case 1:
                         $har = 'N';
                         break;
                    case -1:
                         $har = 'S';
                         break;
                    case 0:
                         $har = 'E';
             }
             break;
       case 'W':
             switch($far[$f]){
                    case 1:
                         $har = 'S';
                         break;
                    case -1:
                         $har = 'N';
                         break;
                    case 0:
                         $har = 'W';
             }
}
echo "Исходное направление робота - $mass[$n].<br>";
echo "Команда - $far[$f].<br>";
echo "Направление робота после выполнения команды $har.";
?>

/*
<?php
echo "Самтхо:<br>";
echo "N - Шимол<br>S - Чануб<br>E - Шарк<br>W - Гарб<br>";
echo "Фармонхо:<br>";
echo "1 - Ба чап<br>-1 - Ба рост<br>0 - Давом додан<br>";
$mass = array('N', 'S', 'E', 'W');
$n = rand(0, 3);
$far = array(1, -1, 0);
$f = rand(0, 2);
switch($mass[$n]){
       case 'N':
             switch($far[$f]){
                    case 1:
                         $har = 'W';
                         break;
                    case -1:
                         $har = 'E';
                         break;
                    case 0:
                         $har = 'N';
             }
             break;
       case 'S':
             switch($far[$f]){
                    case 1:
                         $har = 'E';
                         break;
                    case -1:
                         $har = 'W';
                         break;
                    case 0:
                         $har = 'S';
             }
             break;
       case 'E':
             switch($far[$f]){
                    case 1:
                         $har = 'N';
                         break;
                    case -1:
                         $har = 'S';
                         break;
                    case 0:
                         $har = 'E';
             }
             break;
       case 'W':
             switch($far[$f]){
                    case 1:
                         $har = 'S';
                         break;
                    case -1:
                         $har = 'N';
                         break;
                    case 0:
                         $har = 'W';
             }
}
echo "Самти робот $mass[$n] ast.<br>";
echo "Фармон $far[$f] ast.<br>";
echo "Самти робот пас аз ичрои фармон $har ast."
?>
*/